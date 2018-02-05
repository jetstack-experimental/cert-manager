/*
Copyright 2018 Jetstack Ltd.

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

// This file was automatically generated by informer-gen

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=certmanager.k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("acmevalidations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha1().ACMEValidations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("certificates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha1().Certificates().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterissuers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha1().ClusterIssuers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("issuers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Certmanager().V1alpha1().Issuers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
