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

package v1alpha1

import (
	v1alpha1 "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager/v1alpha1"
	scheme "github.com/munnerz/k8s-api-pager-demo/pkg/client/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TestRunnersGetter has a method to return a TestRunnerInterface.
// A group's client should implement this interface.
type TestRunnersGetter interface {
	TestRunners(namespace string) TestRunnerInterface
}

// TestRunnerInterface has methods to work with TestRunner resources.
type TestRunnerInterface interface {
	Create(*v1alpha1.TestRunner) (*v1alpha1.TestRunner, error)
	Update(*v1alpha1.TestRunner) (*v1alpha1.TestRunner, error)
	UpdateStatus(*v1alpha1.TestRunner) (*v1alpha1.TestRunner, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TestRunner, error)
	List(opts v1.ListOptions) (*v1alpha1.TestRunnerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestRunner, err error)
	TestRunnerExpansion
}

// testRunners implements TestRunnerInterface
type testRunners struct {
	client rest.Interface
	ns     string
}

// newTestRunners returns a TestRunners
func newTestRunners(c *PagerV1alpha1Client, namespace string) *testRunners {
	return &testRunners{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a testRunner and creates it.  Returns the server's representation of the testRunner, and an error, if there is any.
func (c *testRunners) Create(testRunner *v1alpha1.TestRunner) (result *v1alpha1.TestRunner, err error) {
	result = &v1alpha1.TestRunner{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testrunners").
		Body(testRunner).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testRunner and updates it. Returns the server's representation of the testRunner, and an error, if there is any.
func (c *testRunners) Update(testRunner *v1alpha1.TestRunner) (result *v1alpha1.TestRunner, err error) {
	result = &v1alpha1.TestRunner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testrunners").
		Name(testRunner.Name).
		Body(testRunner).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *testRunners) UpdateStatus(testRunner *v1alpha1.TestRunner) (result *v1alpha1.TestRunner, err error) {
	result = &v1alpha1.TestRunner{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testrunners").
		Name(testRunner.Name).
		SubResource("status").
		Body(testRunner).
		Do().
		Into(result)
	return
}

// Delete takes name of the testRunner and deletes it. Returns an error if one occurs.
func (c *testRunners) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testrunners").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testRunners) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testrunners").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the testRunner, and returns the corresponding testRunner object, and an error if there is any.
func (c *testRunners) Get(name string, options v1.GetOptions) (result *v1alpha1.TestRunner, err error) {
	result = &v1alpha1.TestRunner{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testrunners").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestRunners that match those selectors.
func (c *testRunners) List(opts v1.ListOptions) (result *v1alpha1.TestRunnerList, err error) {
	result = &v1alpha1.TestRunnerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testrunners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testRunners.
func (c *testRunners) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testrunners").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched testRunner.
func (c *testRunners) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TestRunner, err error) {
	result = &v1alpha1.TestRunner{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testrunners").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
