/*
Copyright 2020 The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"time"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	gatewayapi "sigs.k8s.io/gateway-api/apis/v1alpha1"
	gatewayclient "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	gatewayinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
	gapilisters "sigs.k8s.io/gateway-api/pkg/client/listers/apis/v1alpha1"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	controllerpkg "github.com/jetstack/cert-manager/pkg/controller"
	shimhelper "github.com/jetstack/cert-manager/pkg/controller/certificate-shim"
	logf "github.com/jetstack/cert-manager/pkg/logs"
)

const (
	ControllerName = "gateway-shim"

	// resyncPeriod is set to 10 hours across cert-manager. These 10 hours come
	// from a discussion on the controller-runtime project that boils down to:
	// never change this without an explicit reason.
	// https://github.com/kubernetes-sigs/controller-runtime/pull/88#issuecomment-408500629
	resyncPeriod = 10 * time.Hour
)

type defaults struct {
	autoCertificateAnnotations          []string
	issuerName, issuerKind, issuerGroup string
}

type controller struct {
	gatewayList gapilisters.GatewayLister
	sync        func(context.Context, metav1.Object) error
}

func (c *controller) Register(ctx *controllerpkg.Context) (workqueue.RateLimitingInterface, []cache.InformerSynced, error) {
	// The user may have enabled the gateway-shim controller but forgot to
	// install the Gateway API CRDs. This will cause cert-manager to go into
	// CrashLoopBackoff which is nice and obvious.
	d, err := discovery.NewDiscoveryClientForConfig(ctx.RESTConfig)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: couldn't construct discovery client: %w", ControllerName, err)
	}
	resources, err := d.ServerResourcesForGroupVersion(gatewayapi.GroupVersion.String())
	if err != nil {
		return nil, nil, fmt.Errorf("%s: couldn't discover gateway API resources (are the Gateway API CRDS installed?): %w", ControllerName, err)
	}
	if len(resources.APIResources) == 0 {
		return nil, nil, fmt.Errorf("%s: no gateway API resources were discovered (are the Gateway API CRDS installed?)", ControllerName)
	}

	// The Gateway API is an external CRD, which means its shared informers are
	// not available in controllerpkg.Context.
	gapiShared := gatewayinformers.NewSharedInformerFactory(gatewayclient.NewForConfigOrDie(ctx.RESTConfig), resyncPeriod)
	cmShared := ctx.SharedInformerFactory

	c.gatewayList = gapiShared.Networking().V1alpha1().Gateways().Lister()
	log := logf.FromContext(ctx.RootContext, ControllerName)
	c.sync = shimhelper.SyncFn(ctx.Recorder, log, ctx.CMClient, cmShared.Certmanager().V1().Certificates().Lister(), ctx.IngressShimOptions)

	queue := workqueue.NewNamedRateLimitingQueue(controllerpkg.DefaultItemBasedRateLimiter(), ControllerName)

	gapiShared.Networking().V1alpha1().Gateways().Informer().AddEventHandler(&controllerpkg.QueuingEventHandler{Queue: queue})
	cmShared.Certmanager().V1().Certificates().Informer().AddEventHandler(&controllerpkg.BlockingEventHandler{WorkFunc: certificateDeleted(queue)})

	mustSync := []cache.InformerSynced{
		gapiShared.Networking().V1alpha1().Gateways().Informer().HasSynced,
		cmShared.Certmanager().V1().Certificates().Informer().HasSynced,
	}

	return queue, mustSync, nil
}

func (c *controller) ProcessItem(ctx context.Context, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	crt, err := c.gatewayList.Gateways(namespace).Get(name)

	if err != nil {
		if k8sErrors.IsNotFound(err) {
			runtime.HandleError(fmt.Errorf("ingress '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	return c.sync(ctx, crt)
}

func certificateDeleted(queue workqueue.RateLimitingInterface) func(obj interface{}) {
	return func(obj interface{}) {
		crt, ok := obj.(*cmapi.Certificate)
		if !ok {
			runtime.HandleError(fmt.Errorf("not a Certificate object: %#v", obj))
			return
		}

		ref := metav1.GetControllerOf(crt)
		if ref == nil {
			// No controller should care about orphans being deleted or
			// updated.
			return
		}

		if ref.Kind != "Ingress" {
			return
		}
		queue.Add(crt.Namespace + "/" + ref.Name)
	}
}

func init() {
	controllerpkg.Register(ControllerName, func(ctx *controllerpkg.Context) (controllerpkg.Interface, error) {
		return controllerpkg.NewBuilder(ctx, ControllerName).
			For(&controller{}).
			Complete()
	})
}
