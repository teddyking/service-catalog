/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	scheme "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExtensionRequestsGetter has a method to return a ExtensionRequestInterface.
// A group's client should implement this interface.
type ExtensionRequestsGetter interface {
	ExtensionRequests(namespace string) ExtensionRequestInterface
}

// ExtensionRequestInterface has methods to work with ExtensionRequest resources.
type ExtensionRequestInterface interface {
	Create(*v1beta1.ExtensionRequest) (*v1beta1.ExtensionRequest, error)
	Update(*v1beta1.ExtensionRequest) (*v1beta1.ExtensionRequest, error)
	UpdateStatus(*v1beta1.ExtensionRequest) (*v1beta1.ExtensionRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.ExtensionRequest, error)
	List(opts v1.ListOptions) (*v1beta1.ExtensionRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ExtensionRequest, err error)
	ExtensionRequestExpansion
}

// extensionRequests implements ExtensionRequestInterface
type extensionRequests struct {
	client rest.Interface
	ns     string
}

// newExtensionRequests returns a ExtensionRequests
func newExtensionRequests(c *ServicecatalogV1beta1Client, namespace string) *extensionRequests {
	return &extensionRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the extensionRequest, and returns the corresponding extensionRequest object, and an error if there is any.
func (c *extensionRequests) Get(name string, options v1.GetOptions) (result *v1beta1.ExtensionRequest, err error) {
	result = &v1beta1.ExtensionRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extensionrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExtensionRequests that match those selectors.
func (c *extensionRequests) List(opts v1.ListOptions) (result *v1beta1.ExtensionRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ExtensionRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extensionrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested extensionRequests.
func (c *extensionRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("extensionrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a extensionRequest and creates it.  Returns the server's representation of the extensionRequest, and an error, if there is any.
func (c *extensionRequests) Create(extensionRequest *v1beta1.ExtensionRequest) (result *v1beta1.ExtensionRequest, err error) {
	result = &v1beta1.ExtensionRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("extensionrequests").
		Body(extensionRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a extensionRequest and updates it. Returns the server's representation of the extensionRequest, and an error, if there is any.
func (c *extensionRequests) Update(extensionRequest *v1beta1.ExtensionRequest) (result *v1beta1.ExtensionRequest, err error) {
	result = &v1beta1.ExtensionRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extensionrequests").
		Name(extensionRequest.Name).
		Body(extensionRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *extensionRequests) UpdateStatus(extensionRequest *v1beta1.ExtensionRequest) (result *v1beta1.ExtensionRequest, err error) {
	result = &v1beta1.ExtensionRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extensionrequests").
		Name(extensionRequest.Name).
		SubResource("status").
		Body(extensionRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the extensionRequest and deletes it. Returns an error if one occurs.
func (c *extensionRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extensionrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *extensionRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extensionrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched extensionRequest.
func (c *extensionRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ExtensionRequest, err error) {
	result = &v1beta1.ExtensionRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("extensionrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
