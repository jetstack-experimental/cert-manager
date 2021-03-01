/*
Copyright The cert-manager Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	acmev1beta1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1beta1"
	versioned "github.com/cert-manager/cert-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cert-manager/cert-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/cert-manager/cert-manager/pkg/client/listers/acme/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ChallengeInformer provides access to a shared informer and lister for
// Challenges.
type ChallengeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ChallengeLister
}

type challengeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChallengeInformer constructs a new informer for Challenge type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChallengeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChallengeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredChallengeInformer constructs a new informer for Challenge type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredChallengeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AcmeV1beta1().Challenges(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AcmeV1beta1().Challenges(namespace).Watch(context.TODO(), options)
			},
		},
		&acmev1beta1.Challenge{},
		resyncPeriod,
		indexers,
	)
}

func (f *challengeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredChallengeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *challengeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&acmev1beta1.Challenge{}, f.defaultInformer)
}

func (f *challengeInformer) Lister() v1beta1.ChallengeLister {
	return v1beta1.NewChallengeLister(f.Informer().GetIndexer())
}
