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

// FakeAzureActivityLogsSources implements AzureActivityLogsSourceInterface
type FakeAzureActivityLogsSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var azureactivitylogssourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "azureactivitylogssources"}

var azureactivitylogssourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "AzureActivityLogsSource"}

// Get takes name of the azureActivityLogsSource, and returns the corresponding azureActivityLogsSource object, and an error if there is any.
func (c *FakeAzureActivityLogsSources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureActivityLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(azureactivitylogssourcesResource, c.ns, name), &v1alpha1.AzureActivityLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureActivityLogsSource), err
}

// List takes label and field selectors, and returns the list of AzureActivityLogsSources that match those selectors.
func (c *FakeAzureActivityLogsSources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureActivityLogsSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(azureactivitylogssourcesResource, azureactivitylogssourcesKind, c.ns, opts), &v1alpha1.AzureActivityLogsSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AzureActivityLogsSourceList{ListMeta: obj.(*v1alpha1.AzureActivityLogsSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AzureActivityLogsSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested azureActivityLogsSources.
func (c *FakeAzureActivityLogsSources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(azureactivitylogssourcesResource, c.ns, opts))

}

// Create takes the representation of a azureActivityLogsSource and creates it.  Returns the server's representation of the azureActivityLogsSource, and an error, if there is any.
func (c *FakeAzureActivityLogsSources) Create(ctx context.Context, azureActivityLogsSource *v1alpha1.AzureActivityLogsSource, opts v1.CreateOptions) (result *v1alpha1.AzureActivityLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(azureactivitylogssourcesResource, c.ns, azureActivityLogsSource), &v1alpha1.AzureActivityLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureActivityLogsSource), err
}

// Update takes the representation of a azureActivityLogsSource and updates it. Returns the server's representation of the azureActivityLogsSource, and an error, if there is any.
func (c *FakeAzureActivityLogsSources) Update(ctx context.Context, azureActivityLogsSource *v1alpha1.AzureActivityLogsSource, opts v1.UpdateOptions) (result *v1alpha1.AzureActivityLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(azureactivitylogssourcesResource, c.ns, azureActivityLogsSource), &v1alpha1.AzureActivityLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureActivityLogsSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAzureActivityLogsSources) UpdateStatus(ctx context.Context, azureActivityLogsSource *v1alpha1.AzureActivityLogsSource, opts v1.UpdateOptions) (*v1alpha1.AzureActivityLogsSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(azureactivitylogssourcesResource, "status", c.ns, azureActivityLogsSource), &v1alpha1.AzureActivityLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureActivityLogsSource), err
}

// Delete takes name of the azureActivityLogsSource and deletes it. Returns an error if one occurs.
func (c *FakeAzureActivityLogsSources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(azureactivitylogssourcesResource, c.ns, name, opts), &v1alpha1.AzureActivityLogsSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAzureActivityLogsSources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(azureactivitylogssourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AzureActivityLogsSourceList{})
	return err
}

// Patch applies the patch and returns the patched azureActivityLogsSource.
func (c *FakeAzureActivityLogsSources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureActivityLogsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(azureactivitylogssourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AzureActivityLogsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AzureActivityLogsSource), err
}