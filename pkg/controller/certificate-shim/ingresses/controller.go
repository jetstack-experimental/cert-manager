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

package ingresses

import (
	"context"
	"fmt"

	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	networkinglisters "k8s.io/client-go/listers/networking/v1beta1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	controllerpkg "github.com/jetstack/cert-manager/pkg/controller"
	shimhelper "github.com/jetstack/cert-manager/pkg/controller/certificate-shim"
	logf "github.com/jetstack/cert-manager/pkg/logs"
)

const (
	ControllerName = "ingress-shim"
)

type controller struct {
	ingressLister networkinglisters.IngressLister
	sync          shimhelper.SyncFn
}

func (c *controller) Register(ctx *controllerpkg.Context) (workqueue.RateLimitingInterface, []cache.InformerSynced, error) {
	kShared := ctx.KubeSharedInformerFactory
	cmShared := ctx.SharedInformerFactory

	c.ingressLister = kShared.Networking().V1beta1().Ingresses().Lister()
	log := logf.FromContext(ctx.RootContext, ControllerName)
	c.sync = shimhelper.SyncFnFor(ctx.Recorder, log, ctx.CMClient, cmShared.Certmanager().V1().Certificates().Lister(), ctx.IngressShimOptions)

	queue := workqueue.NewNamedRateLimitingQueue(controllerpkg.DefaultItemBasedRateLimiter(), ControllerName)

	mustSync := []cache.InformerSynced{
		kShared.Networking().V1beta1().Ingresses().Informer().HasSynced,
		cmShared.Certmanager().V1().Certificates().Informer().HasSynced,
	}

	kShared.Networking().V1beta1().Ingresses().Informer().AddEventHandler(&controllerpkg.QueuingEventHandler{Queue: queue})
	cmShared.Certmanager().V1().Certificates().Informer().AddEventHandler(&controllerpkg.BlockingEventHandler{WorkFunc: certificateDeleted(queue)})

	return queue, mustSync, nil
}

func (c *controller) ProcessItem(ctx context.Context, key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	crt, err := c.ingressLister.Ingresses(namespace).Get(name)

	if err != nil {
		if k8sErrors.IsNotFound(err) {
			runtime.HandleError(fmt.Errorf("ingress '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	return c.sync(ctx, crt)
}

// Whenever a Certificate gets deleted, we want to reconcile its parent Ingress.
// This parent Ingress is called "controller object". For example, the following
// Certificate is controlled by the Ingress "example":
//
//     kind: Certificate
//     metadata:
//       namespace: cert-that-was-deleted
//       ownerReferences:
//       - controller: true                                       ← this
//         apiVersion: networking.k8s.io/v1beta1
//         kind: Ingress
//         name: example
//         blockOwnerDeletion: true
//         uid: 7d3897c2-ce27-4144-883a-e1b5f89bd65a
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

		// We don't check the apiVersion e.g. "networking.k8s.io/v1beta1"
		// because there is no chance that another object called "Ingress" be
		// the controller of a Certificate.
		if ref.Kind != "Ingress" {
			return
		}

		// Queue items are simple strings of the form "namespace-1/ingress-1".
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
