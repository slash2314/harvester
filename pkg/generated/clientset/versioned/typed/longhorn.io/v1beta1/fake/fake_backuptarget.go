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

// FakeBackupTargets implements BackupTargetInterface
type FakeBackupTargets struct {
	Fake *FakeLonghornV1beta1
	ns   string
}

var backuptargetsResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta1", Resource: "backuptargets"}

var backuptargetsKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta1", Kind: "BackupTarget"}

// Get takes name of the backupTarget, and returns the corresponding backupTarget object, and an error if there is any.
func (c *FakeBackupTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BackupTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backuptargetsResource, c.ns, name), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

// List takes label and field selectors, and returns the list of BackupTargets that match those selectors.
func (c *FakeBackupTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BackupTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backuptargetsResource, backuptargetsKind, c.ns, opts), &v1beta1.BackupTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BackupTargetList{ListMeta: obj.(*v1beta1.BackupTargetList).ListMeta}
	for _, item := range obj.(*v1beta1.BackupTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backupTargets.
func (c *FakeBackupTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backuptargetsResource, c.ns, opts))

}

// Create takes the representation of a backupTarget and creates it.  Returns the server's representation of the backupTarget, and an error, if there is any.
func (c *FakeBackupTargets) Create(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.CreateOptions) (result *v1beta1.BackupTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backuptargetsResource, c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

// Update takes the representation of a backupTarget and updates it. Returns the server's representation of the backupTarget, and an error, if there is any.
func (c *FakeBackupTargets) Update(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (result *v1beta1.BackupTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backuptargetsResource, c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackupTargets) UpdateStatus(ctx context.Context, backupTarget *v1beta1.BackupTarget, opts v1.UpdateOptions) (*v1beta1.BackupTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backuptargetsResource, "status", c.ns, backupTarget), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}

// Delete takes name of the backupTarget and deletes it. Returns an error if one occurs.
func (c *FakeBackupTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(backuptargetsResource, c.ns, name, opts), &v1beta1.BackupTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackupTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backuptargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BackupTargetList{})
	return err
}

// Patch applies the patch and returns the patched backupTarget.
func (c *FakeBackupTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BackupTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backuptargetsResource, c.ns, name, pt, data, subresources...), &v1beta1.BackupTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BackupTarget), err
}
