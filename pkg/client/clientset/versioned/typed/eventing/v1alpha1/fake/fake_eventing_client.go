// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/w-h-a/crd/pkg/client/clientset/versioned/typed/eventing/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEventingV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEventingV1alpha1) EventSources(namespace string) v1alpha1.EventSourceInterface {
	return &FakeEventSources{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEventingV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
