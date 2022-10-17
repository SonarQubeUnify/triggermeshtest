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

// FakeAWSS3Sources implements AWSS3SourceInterface
type FakeAWSS3Sources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var awss3sourcesResource = schema.GroupVersionResource{Group: "sources.triggermesh.io", Version: "v1alpha1", Resource: "awss3sources"}

var awss3sourcesKind = schema.GroupVersionKind{Group: "sources.triggermesh.io", Version: "v1alpha1", Kind: "AWSS3Source"}

// Get takes name of the aWSS3Source, and returns the corresponding aWSS3Source object, and an error if there is any.
func (c *FakeAWSS3Sources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSS3Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awss3sourcesResource, c.ns, name), &v1alpha1.AWSS3Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSS3Source), err
}

// List takes label and field selectors, and returns the list of AWSS3Sources that match those selectors.
func (c *FakeAWSS3Sources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSS3SourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awss3sourcesResource, awss3sourcesKind, c.ns, opts), &v1alpha1.AWSS3SourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSS3SourceList{ListMeta: obj.(*v1alpha1.AWSS3SourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSS3SourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSS3Sources.
func (c *FakeAWSS3Sources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awss3sourcesResource, c.ns, opts))

}

// Create takes the representation of a aWSS3Source and creates it.  Returns the server's representation of the aWSS3Source, and an error, if there is any.
func (c *FakeAWSS3Sources) Create(ctx context.Context, aWSS3Source *v1alpha1.AWSS3Source, opts v1.CreateOptions) (result *v1alpha1.AWSS3Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awss3sourcesResource, c.ns, aWSS3Source), &v1alpha1.AWSS3Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSS3Source), err
}

// Update takes the representation of a aWSS3Source and updates it. Returns the server's representation of the aWSS3Source, and an error, if there is any.
func (c *FakeAWSS3Sources) Update(ctx context.Context, aWSS3Source *v1alpha1.AWSS3Source, opts v1.UpdateOptions) (result *v1alpha1.AWSS3Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awss3sourcesResource, c.ns, aWSS3Source), &v1alpha1.AWSS3Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSS3Source), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSS3Sources) UpdateStatus(ctx context.Context, aWSS3Source *v1alpha1.AWSS3Source, opts v1.UpdateOptions) (*v1alpha1.AWSS3Source, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awss3sourcesResource, "status", c.ns, aWSS3Source), &v1alpha1.AWSS3Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSS3Source), err
}

// Delete takes name of the aWSS3Source and deletes it. Returns an error if one occurs.
func (c *FakeAWSS3Sources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(awss3sourcesResource, c.ns, name, opts), &v1alpha1.AWSS3Source{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSS3Sources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awss3sourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSS3SourceList{})
	return err
}

// Patch applies the patch and returns the patched aWSS3Source.
func (c *FakeAWSS3Sources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSS3Source, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awss3sourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSS3Source{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSS3Source), err
}
