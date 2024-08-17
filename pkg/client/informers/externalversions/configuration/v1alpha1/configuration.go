// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	configurationv1alpha1 "github.com/w-h-a/crd/pkg/apis/configuration/v1alpha1"
	versioned "github.com/w-h-a/crd/pkg/client/clientset/versioned"
	internalinterfaces "github.com/w-h-a/crd/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/w-h-a/crd/pkg/client/listers/configuration/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConfigurationInformer provides access to a shared informer and lister for
// Configurations.
type ConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ConfigurationLister
}

type configurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewConfigurationInformer constructs a new informer for Configuration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredConfigurationInformer constructs a new informer for Configuration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigurationV1alpha1().Configurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigurationV1alpha1().Configurations(namespace).Watch(context.TODO(), options)
			},
		},
		&configurationv1alpha1.Configuration{},
		resyncPeriod,
		indexers,
	)
}

func (f *configurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *configurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configurationv1alpha1.Configuration{}, f.defaultInformer)
}

func (f *configurationInformer) Lister() v1alpha1.ConfigurationLister {
	return v1alpha1.NewConfigurationLister(f.Informer().GetIndexer())
}
