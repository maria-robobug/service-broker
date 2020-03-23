/* Copyright (C) Couchbase, Inc 2020 - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/couchbase/service-broker/generated/clientset/versioned"
	internalinterfaces "github.com/couchbase/service-broker/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/couchbase/service-broker/generated/listers/broker.couchbase.com/v1alpha1"
	brokercouchbasecomv1alpha1 "github.com/couchbase/service-broker/pkg/apis/broker.couchbase.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CouchbaseServiceBrokerConfigInformer provides access to a shared informer and lister for
// CouchbaseServiceBrokerConfigs.
type CouchbaseServiceBrokerConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CouchbaseServiceBrokerConfigLister
}

type couchbaseServiceBrokerConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCouchbaseServiceBrokerConfigInformer constructs a new informer for CouchbaseServiceBrokerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCouchbaseServiceBrokerConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCouchbaseServiceBrokerConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCouchbaseServiceBrokerConfigInformer constructs a new informer for CouchbaseServiceBrokerConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCouchbaseServiceBrokerConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BrokerV1alpha1().CouchbaseServiceBrokerConfigs(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BrokerV1alpha1().CouchbaseServiceBrokerConfigs(namespace).Watch(options)
			},
		},
		&brokercouchbasecomv1alpha1.CouchbaseServiceBrokerConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *couchbaseServiceBrokerConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCouchbaseServiceBrokerConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *couchbaseServiceBrokerConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&brokercouchbasecomv1alpha1.CouchbaseServiceBrokerConfig{}, f.defaultInformer)
}

func (f *couchbaseServiceBrokerConfigInformer) Lister() v1alpha1.CouchbaseServiceBrokerConfigLister {
	return v1alpha1.NewCouchbaseServiceBrokerConfigLister(f.Informer().GetIndexer())
}