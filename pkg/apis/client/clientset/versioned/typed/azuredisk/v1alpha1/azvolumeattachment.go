/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/azuredisk/v1alpha1"
	scheme "sigs.k8s.io/azuredisk-csi-driver/pkg/apis/client/clientset/versioned/scheme"
)

// AzVolumeAttachmentsGetter has a method to return a AzVolumeAttachmentInterface.
// A group's client should implement this interface.
type AzVolumeAttachmentsGetter interface {
	AzVolumeAttachments(namespace string) AzVolumeAttachmentInterface
}

// AzVolumeAttachmentInterface has methods to work with AzVolumeAttachment resources.
type AzVolumeAttachmentInterface interface {
	Create(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.CreateOptions) (*v1alpha1.AzVolumeAttachment, error)
	Update(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (*v1alpha1.AzVolumeAttachment, error)
	UpdateStatus(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (*v1alpha1.AzVolumeAttachment, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzVolumeAttachment, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzVolumeAttachmentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzVolumeAttachment, err error)
	AzVolumeAttachmentExpansion
}

// azVolumeAttachments implements AzVolumeAttachmentInterface
type azVolumeAttachments struct {
	client rest.Interface
	ns     string
}

// newAzVolumeAttachments returns a AzVolumeAttachments
func newAzVolumeAttachments(c *DiskV1alpha1Client, namespace string) *azVolumeAttachments {
	return &azVolumeAttachments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azVolumeAttachment, and returns the corresponding azVolumeAttachment object, and an error if there is any.
func (c *azVolumeAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	result = &v1alpha1.AzVolumeAttachment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzVolumeAttachments that match those selectors.
func (c *azVolumeAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzVolumeAttachmentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzVolumeAttachmentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azVolumeAttachments.
func (c *azVolumeAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azVolumeAttachment and creates it.  Returns the server's representation of the azVolumeAttachment, and an error, if there is any.
func (c *azVolumeAttachments) Create(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.CreateOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	result = &v1alpha1.AzVolumeAttachment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azVolumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azVolumeAttachment and updates it. Returns the server's representation of the azVolumeAttachment, and an error, if there is any.
func (c *azVolumeAttachments) Update(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	result = &v1alpha1.AzVolumeAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		Name(azVolumeAttachment.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azVolumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *azVolumeAttachments) UpdateStatus(ctx context.Context, azVolumeAttachment *v1alpha1.AzVolumeAttachment, opts v1.UpdateOptions) (result *v1alpha1.AzVolumeAttachment, err error) {
	result = &v1alpha1.AzVolumeAttachment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		Name(azVolumeAttachment.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azVolumeAttachment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azVolumeAttachment and deletes it. Returns an error if one occurs.
func (c *azVolumeAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azVolumeAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azvolumeattachments").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azVolumeAttachment.
func (c *azVolumeAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzVolumeAttachment, err error) {
	result = &v1alpha1.AzVolumeAttachment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azvolumeattachments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
