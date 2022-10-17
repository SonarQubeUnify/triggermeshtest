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

// AzureServiceBusTopicSourceInformer provides access to a shared informer and lister for
// AzureServiceBusTopicSources.
type AzureServiceBusTopicSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AzureServiceBusTopicSourceLister
}

type azureServiceBusTopicSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAzureServiceBusTopicSourceInformer constructs a new informer for AzureServiceBusTopicSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureServiceBusTopicSourceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureServiceBusTopicSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAzureServiceBusTopicSourceInformer constructs a new informer for AzureServiceBusTopicSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureServiceBusTopicSourceInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().AzureServiceBusTopicSources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().AzureServiceBusTopicSources(namespace).Watch(context.TODO(), options)
			},
		},
		&sourcesv1alpha1.AzureServiceBusTopicSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureServiceBusTopicSourceInformer) defaultInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureServiceBusTopicSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureServiceBusTopicSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sourcesv1alpha1.AzureServiceBusTopicSource{}, f.defaultInformer)
}

func (f *azureServiceBusTopicSourceInformer) Lister() v1alpha1.AzureServiceBusTopicSourceLister {
	return v1alpha1.NewAzureServiceBusTopicSourceLister(f.Informer().GetIndexer())
}