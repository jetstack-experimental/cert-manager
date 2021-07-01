/*
Copyright 2021 The cert-manager Authors.

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

package cmapichecker

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
)

// Interface is used to check that the cert-manager CRDs have been installed and are usable.
type Interface interface {
	Check(context.Context) error
}

type cmapiChecker struct {
	dryRunClient client.Client
	namespace    string
}

// New returns a cert-manager API checker
func New(restcfg *rest.Config, namespace string) (Interface, error) {
	scheme := runtime.NewScheme()
	if err := cmapi.AddToScheme(scheme); err != nil {
		return nil, err
	}
	cl, err := client.New(restcfg, client.Options{
		Scheme: scheme,
	})
	if err != nil {
		return nil, err
	}
	return &cmapiChecker{
		dryRunClient: client.NewDryRunClient(cl),
		namespace:    namespace,
	}, nil
}

// Check attempts to perform a dry-run create of a cert-manager Certificate
// resource in order to verify that CRDs are installed and all the required
// webhooks are reachable by the K8S API server.
func (o *cmapiChecker) Check(ctx context.Context) error {
	cert := &cmapi.Certificate{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "cmapichecker-",
			Namespace:    o.namespace,
		},
		Spec: cmapi.CertificateSpec{
			DNSNames:   []string{"cmapichecker.example"},
			SecretName: "cmapichecker",
			IssuerRef: cmmeta.ObjectReference{
				Name: "cmapichecker",
			},
		},
	}
	if err := o.dryRunClient.Create(ctx, cert); err != nil {
		return fmt.Errorf("error creating Certificate: %v", err)
	}
	return nil
}
