/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTestRunners implements TestRunnerInterface
type FakeTestRunners struct {
	Fake *FakePagerV1alpha1
	ns   string
}

var testrunnersResource = schema.GroupVersionResource{Group: "pager.k8s.co", Version: "v1alpha1", Resource: "testrunners"}

var testrunnersKind = schema.GroupVersionKind{Group: "pager.k8s.co", Version: "v1alpha1", Kind: "TestRunner"}

func (c *FakeTestRunners) Create(testRunner *v1alpha1.TestRunner) (result *v1alpha1.TestRunner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(testrunnersResource, c.ns, testRunner), &v1alpha1.TestRunner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestRunner), err
}

func (c *FakeTestRunners) Update(testRunner *v1alpha1.TestRunner) (result *v1alpha1.TestRunner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(testrunnersResource, c.ns, testRunner), &v1alpha1.TestRunner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestRunner), err
}

func (c *FakeTestRunners) UpdateStatus(testRunner *v1alpha1.TestRunner) (*v1alpha1.TestRunner, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(testrunnersResource, "status", c.ns, testRunner), &v1alpha1.TestRunner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestRunner), err
}

func (c *FakeTestRunners) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(testrunnersResource, c.ns, name), &v1alpha1.TestRunner{})

	return err
}

func (c *FakeTestRunners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(testrunnersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TestRunnerList{})
	return err
}

func (c *FakeTestRunners) Get(name string, options v1.GetOptions) (result *v1alpha1.TestRunner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(testrunnersResource, c.ns, name), &v1alpha1.TestRunner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestRunner), err
}

func (c *FakeTestRunners) List(opts v1.ListOptions) (result *v1alpha1.TestRunnerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(testrunnersResource, testrunnersKind, c.ns, opts), &v1alpha1.TestRunnerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TestRunnerList{}
	for _, item := range obj.(*v1alpha1.TestRunnerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested testRunners.
func (c *FakeTestRunners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(testrunnersResource, c.ns, opts))

}

// Patch applies the patch and returns the patched testRunner.
func (c *FakeTestRunners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestRunner, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(testrunnersResource, c.ns, name, data, subresources...), &v1alpha1.TestRunner{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TestRunner), err
}
