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

	v1alpha1 "github.com/triggermesh/triggermesh/pkg/apis/targets/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlibabaOSSTargets implements AlibabaOSSTargetInterface
type FakeAlibabaOSSTargets struct {
	Fake *FakeTargetsV1alpha1
	ns   string
}

var alibabaosstargetsResource = schema.GroupVersionResource{Group: "targets.triggermesh.io", Version: "v1alpha1", Resource: "alibabaosstargets"}

var alibabaosstargetsKind = schema.GroupVersionKind{Group: "targets.triggermesh.io", Version: "v1alpha1", Kind: "AlibabaOSSTarget"}

// Get takes name of the alibabaOSSTarget, and returns the corresponding alibabaOSSTarget object, and an error if there is any.
func (c *FakeAlibabaOSSTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AlibabaOSSTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alibabaosstargetsResource, c.ns, name), &v1alpha1.AlibabaOSSTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlibabaOSSTarget), err
}

// List takes label and field selectors, and returns the list of AlibabaOSSTargets that match those selectors.
func (c *FakeAlibabaOSSTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlibabaOSSTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alibabaosstargetsResource, alibabaosstargetsKind, c.ns, opts), &v1alpha1.AlibabaOSSTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlibabaOSSTargetList{ListMeta: obj.(*v1alpha1.AlibabaOSSTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlibabaOSSTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alibabaOSSTargets.
func (c *FakeAlibabaOSSTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alibabaosstargetsResource, c.ns, opts))

}

// Create takes the representation of a alibabaOSSTarget and creates it.  Returns the server's representation of the alibabaOSSTarget, and an error, if there is any.
func (c *FakeAlibabaOSSTargets) Create(ctx context.Context, alibabaOSSTarget *v1alpha1.AlibabaOSSTarget, opts v1.CreateOptions) (result *v1alpha1.AlibabaOSSTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alibabaosstargetsResource, c.ns, alibabaOSSTarget), &v1alpha1.AlibabaOSSTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlibabaOSSTarget), err
}

// Update takes the representation of a alibabaOSSTarget and updates it. Returns the server's representation of the alibabaOSSTarget, and an error, if there is any.
func (c *FakeAlibabaOSSTargets) Update(ctx context.Context, alibabaOSSTarget *v1alpha1.AlibabaOSSTarget, opts v1.UpdateOptions) (result *v1alpha1.AlibabaOSSTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alibabaosstargetsResource, c.ns, alibabaOSSTarget), &v1alpha1.AlibabaOSSTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlibabaOSSTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlibabaOSSTargets) UpdateStatus(ctx context.Context, alibabaOSSTarget *v1alpha1.AlibabaOSSTarget, opts v1.UpdateOptions) (*v1alpha1.AlibabaOSSTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alibabaosstargetsResource, "status", c.ns, alibabaOSSTarget), &v1alpha1.AlibabaOSSTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlibabaOSSTarget), err
}

// Delete takes name of the alibabaOSSTarget and deletes it. Returns an error if one occurs.
func (c *FakeAlibabaOSSTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(alibabaosstargetsResource, c.ns, name, opts), &v1alpha1.AlibabaOSSTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlibabaOSSTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alibabaosstargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlibabaOSSTargetList{})
	return err
}

// Patch applies the patch and returns the patched alibabaOSSTarget.
func (c *FakeAlibabaOSSTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AlibabaOSSTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alibabaosstargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AlibabaOSSTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AlibabaOSSTarget), err
}
