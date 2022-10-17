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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPPollerSources implements HTTPPollerSourceInterface
type FakeHTTPPollerSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var httppollersourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "httppollersources"}

var httppollersourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "HTTPPollerSource"}

// Get takes name of the hTTPPollerSource, and returns the corresponding hTTPPollerSource object, and an error if there is any.
func (c *FakeHTTPPollerSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HTTPPollerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httppollersourcesResource, c.ns, name), &v1alpha1.HTTPPollerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPPollerSource), err
}

// List takes label and field selectors, and returns the list of HTTPPollerSources that match those selectors.
func (c *FakeHTTPPollerSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HTTPPollerSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httppollersourcesResource, httppollersourcesKind, c.ns, opts), &v1alpha1.HTTPPollerSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HTTPPollerSourceList{ListMeta: obj.(*v1alpha1.HTTPPollerSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.HTTPPollerSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPPollerSources.
func (c *FakeHTTPPollerSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httppollersourcesResource, c.ns, opts))

}

// Create takes the representation of a hTTPPollerSource and creates it.  Returns the server's representation of the hTTPPollerSource, and an error, if there is any.
func (c *FakeHTTPPollerSources) Create(ctx context.Context, hTTPPollerSource *v1alpha1.HTTPPollerSource, opts v1.CreateOptions) (result *v1alpha1.HTTPPollerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httppollersourcesResource, c.ns, hTTPPollerSource), &v1alpha1.HTTPPollerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPPollerSource), err
}

// Update takes the representation of a hTTPPollerSource and updates it. Returns the server's representation of the hTTPPollerSource, and an error, if there is any.
func (c *FakeHTTPPollerSources) Update(ctx context.Context, hTTPPollerSource *v1alpha1.HTTPPollerSource, opts v1.UpdateOptions) (result *v1alpha1.HTTPPollerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httppollersourcesResource, c.ns, hTTPPollerSource), &v1alpha1.HTTPPollerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPPollerSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHTTPPollerSources) UpdateStatus(ctx context.Context, hTTPPollerSource *v1alpha1.HTTPPollerSource, opts v1.UpdateOptions) (*v1alpha1.HTTPPollerSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(httppollersourcesResource, "status", c.ns, hTTPPollerSource), &v1alpha1.HTTPPollerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPPollerSource), err
}

// Delete takes name of the hTTPPollerSource and deletes it. Returns an error if one occurs.
func (c *FakeHTTPPollerSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(httppollersourcesResource, c.ns, name, opts), &v1alpha1.HTTPPollerSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPPollerSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httppollersourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HTTPPollerSourceList{})
	return err
}

// Patch applies the patch and returns the patched hTTPPollerSource.
func (c *FakeHTTPPollerSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HTTPPollerSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httppollersourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HTTPPollerSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPPollerSource), err
}