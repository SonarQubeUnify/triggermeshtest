/*
Copyright 2022 TriggerMesh Inc.

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

package v1alpha1

import (
	"context"
	time "time"

	sourcesv1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	internalclientset "github.com/triggermesh/triggermesh/pkg/client/generated/clientset/internalclientset"
	internalinterfaces "github.com/triggermesh/triggermesh/pkg/client/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/triggermesh/triggermesh/pkg/client/generated/listers/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureQueueStorageSourceInformer provides access to a shared informer and lister for
// AzureQueueStorageSources.
type AzureQueueStorageSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureQueueStorageSourceLister
}

type azureQueueStorageSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAzureQueueStorageSourceInformer constructs a new informer for AzureQueueStorageSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureQueueStorageSourceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureQueueStorageSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAzureQueueStorageSourceInformer constructs a new informer for AzureQueueStorageSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureQueueStorageSourceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().AzureQueueStorageSources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().AzureQueueStorageSources(namespace).Watch(context.TODO(), options)
			},
		},
		&sourcesv1alpha1.AzureQueueStorageSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureQueueStorageSourceInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureQueueStorageSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureQueueStorageSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sourcesv1alpha1.AzureQueueStorageSource{}, f.defaultInformer)
}

func (f *azureQueueStorageSourceInformer) Lister() v1alpha1.AzureQueueStorageSourceLister {
	return v1alpha1.NewAzureQueueStorageSourceLister(f.Informer().GetIndexer())
}
