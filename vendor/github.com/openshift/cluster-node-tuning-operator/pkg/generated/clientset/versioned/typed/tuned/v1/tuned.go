// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v1"
	scheme "github.com/openshift/cluster-node-tuning-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TunedsGetter has a method to return a TunedInterface.
// A group's client should implement this interface.
type TunedsGetter interface {
	Tuneds(namespace string) TunedInterface
}

// TunedInterface has methods to work with Tuned resources.
type TunedInterface interface {
	Create(*v1.Tuned) (*v1.Tuned, error)
	Update(*v1.Tuned) (*v1.Tuned, error)
	UpdateStatus(*v1.Tuned) (*v1.Tuned, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Tuned, error)
	List(opts metav1.ListOptions) (*v1.TunedList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Tuned, err error)
	TunedExpansion
}

// tuneds implements TunedInterface
type tuneds struct {
	client rest.Interface
	ns     string
}

// newTuneds returns a Tuneds
func newTuneds(c *TunedV1Client, namespace string) *tuneds {
	return &tuneds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tuned, and returns the corresponding tuned object, and an error if there is any.
func (c *tuneds) Get(name string, options metav1.GetOptions) (result *v1.Tuned, err error) {
	result = &v1.Tuned{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tuneds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Tuneds that match those selectors.
func (c *tuneds) List(opts metav1.ListOptions) (result *v1.TunedList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TunedList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tuneds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tuneds.
func (c *tuneds) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tuneds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a tuned and creates it.  Returns the server's representation of the tuned, and an error, if there is any.
func (c *tuneds) Create(tuned *v1.Tuned) (result *v1.Tuned, err error) {
	result = &v1.Tuned{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tuneds").
		Body(tuned).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tuned and updates it. Returns the server's representation of the tuned, and an error, if there is any.
func (c *tuneds) Update(tuned *v1.Tuned) (result *v1.Tuned, err error) {
	result = &v1.Tuned{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tuneds").
		Name(tuned.Name).
		Body(tuned).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *tuneds) UpdateStatus(tuned *v1.Tuned) (result *v1.Tuned, err error) {
	result = &v1.Tuned{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tuneds").
		Name(tuned.Name).
		SubResource("status").
		Body(tuned).
		Do().
		Into(result)
	return
}

// Delete takes name of the tuned and deletes it. Returns an error if one occurs.
func (c *tuneds) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tuneds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tuneds) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tuneds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tuned.
func (c *tuneds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Tuned, err error) {
	result = &v1.Tuned{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tuneds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
