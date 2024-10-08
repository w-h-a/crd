// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/w-h-a/crd/pkg/apis/eventing/v1alpha1"
	eventingv1alpha1 "github.com/w-h-a/crd/pkg/client/applyconfiguration/eventing/v1alpha1"
	scheme "github.com/w-h-a/crd/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EventSourcesGetter has a method to return a EventSourceInterface.
// A group's client should implement this interface.
type EventSourcesGetter interface {
	EventSources(namespace string) EventSourceInterface
}

// EventSourceInterface has methods to work with EventSource resources.
type EventSourceInterface interface {
	Create(ctx context.Context, eventSource *v1alpha1.EventSource, opts v1.CreateOptions) (*v1alpha1.EventSource, error)
	Update(ctx context.Context, eventSource *v1alpha1.EventSource, opts v1.UpdateOptions) (*v1alpha1.EventSource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EventSource, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EventSourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventSource, err error)
	Apply(ctx context.Context, eventSource *eventingv1alpha1.EventSourceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.EventSource, err error)
	EventSourceExpansion
}

// eventSources implements EventSourceInterface
type eventSources struct {
	client rest.Interface
	ns     string
}

// newEventSources returns a EventSources
func newEventSources(c *EventingV1alpha1Client, namespace string) *eventSources {
	return &eventSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eventSource, and returns the corresponding eventSource object, and an error if there is any.
func (c *eventSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventSource, err error) {
	result = &v1alpha1.EventSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EventSources that match those selectors.
func (c *eventSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EventSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eventsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eventSources.
func (c *eventSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eventsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eventSource and creates it.  Returns the server's representation of the eventSource, and an error, if there is any.
func (c *eventSources) Create(ctx context.Context, eventSource *v1alpha1.EventSource, opts v1.CreateOptions) (result *v1alpha1.EventSource, err error) {
	result = &v1alpha1.EventSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eventsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventSource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eventSource and updates it. Returns the server's representation of the eventSource, and an error, if there is any.
func (c *eventSources) Update(ctx context.Context, eventSource *v1alpha1.EventSource, opts v1.UpdateOptions) (result *v1alpha1.EventSource, err error) {
	result = &v1alpha1.EventSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eventsources").
		Name(eventSource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eventSource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eventSource and deletes it. Returns an error if one occurs.
func (c *eventSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventsources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eventSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eventsources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eventSource.
func (c *eventSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventSource, err error) {
	result = &v1alpha1.EventSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eventsources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied eventSource.
func (c *eventSources) Apply(ctx context.Context, eventSource *eventingv1alpha1.EventSourceApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.EventSource, err error) {
	if eventSource == nil {
		return nil, fmt.Errorf("eventSource provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(eventSource)
	if err != nil {
		return nil, err
	}
	name := eventSource.Name
	if name == nil {
		return nil, fmt.Errorf("eventSource.Name must be provided to Apply")
	}
	result = &v1alpha1.EventSource{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("eventsources").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
