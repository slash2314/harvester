/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePreferences implements PreferenceInterface
type FakePreferences struct {
	Fake *FakeHarvesterhciV1beta1
	ns   string
}

var preferencesResource = schema.GroupVersionResource{Group: "harvesterhci.io", Version: "v1beta1", Resource: "preferences"}

var preferencesKind = schema.GroupVersionKind{Group: "harvesterhci.io", Version: "v1beta1", Kind: "Preference"}

// Get takes name of the preference, and returns the corresponding preference object, and an error if there is any.
func (c *FakePreferences) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Preference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(preferencesResource, c.ns, name), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

// List takes label and field selectors, and returns the list of Preferences that match those selectors.
func (c *FakePreferences) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PreferenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(preferencesResource, preferencesKind, c.ns, opts), &v1beta1.PreferenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PreferenceList{ListMeta: obj.(*v1beta1.PreferenceList).ListMeta}
	for _, item := range obj.(*v1beta1.PreferenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested preferences.
func (c *FakePreferences) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(preferencesResource, c.ns, opts))

}

// Create takes the representation of a preference and creates it.  Returns the server's representation of the preference, and an error, if there is any.
func (c *FakePreferences) Create(ctx context.Context, preference *v1beta1.Preference, opts v1.CreateOptions) (result *v1beta1.Preference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(preferencesResource, c.ns, preference), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

// Update takes the representation of a preference and updates it. Returns the server's representation of the preference, and an error, if there is any.
func (c *FakePreferences) Update(ctx context.Context, preference *v1beta1.Preference, opts v1.UpdateOptions) (result *v1beta1.Preference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(preferencesResource, c.ns, preference), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}

// Delete takes name of the preference and deletes it. Returns an error if one occurs.
func (c *FakePreferences) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(preferencesResource, c.ns, name, opts), &v1beta1.Preference{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePreferences) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(preferencesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PreferenceList{})
	return err
}

// Patch applies the patch and returns the patched preference.
func (c *FakePreferences) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Preference, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(preferencesResource, c.ns, name, pt, data, subresources...), &v1beta1.Preference{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Preference), err
}
