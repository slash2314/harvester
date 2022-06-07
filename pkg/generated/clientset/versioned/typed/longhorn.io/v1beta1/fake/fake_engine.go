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

	v1beta1 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEngines implements EngineInterface
type FakeEngines struct {
	Fake *FakeLonghornV1beta1
	ns   string
}

var enginesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "engines"}

var enginesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "Engine"}

// Get takes name of the engine, and returns the corresponding engine object, and an error if there is any.
func (c *FakeEngines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(enginesResource, c.ns, name), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

// List takes label and field selectors, and returns the list of Engines that match those selectors.
func (c *FakeEngines) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.EngineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(enginesResource, enginesKind, c.ns, opts), &v1beta1.EngineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.EngineList{ListMeta: obj.(*v1beta1.EngineList).ListMeta}
	for _, item := range obj.(*v1beta1.EngineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested engines.
func (c *FakeEngines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(enginesResource, c.ns, opts))

}

// Create takes the representation of a engine and creates it.  Returns the server's representation of the engine, and an error, if there is any.
func (c *FakeEngines) Create(ctx context.Context, engine *v1beta1.Engine, opts v1.CreateOptions) (result *v1beta1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(enginesResource, c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

// Update takes the representation of a engine and updates it. Returns the server's representation of the engine, and an error, if there is any.
func (c *FakeEngines) Update(ctx context.Context, engine *v1beta1.Engine, opts v1.UpdateOptions) (result *v1beta1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(enginesResource, c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEngines) UpdateStatus(ctx context.Context, engine *v1beta1.Engine, opts v1.UpdateOptions) (*v1beta1.Engine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(enginesResource, "status", c.ns, engine), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}

// Delete takes name of the engine and deletes it. Returns an error if one occurs.
func (c *FakeEngines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(enginesResource, c.ns, name, opts), &v1beta1.Engine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEngines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(enginesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.EngineList{})
	return err
}

// Patch applies the patch and returns the patched engine.
func (c *FakeEngines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(enginesResource, c.ns, name, pt, data, subresources...), &v1beta1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Engine), err
}
