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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/srossross/k8s-test-runner/pkg/apis/pager/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TestRunLister helps list TestRuns.
type TestRunLister interface {
	// List lists all TestRuns in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TestRun, err error)
	// TestRuns returns an object that can list and get TestRuns.
	TestRuns(namespace string) TestRunNamespaceLister
	TestRunListerExpansion
}

// testRunLister implements the TestRunLister interface.
type testRunLister struct {
	indexer cache.Indexer
}

// NewTestRunLister returns a new TestRunLister.
func NewTestRunLister(indexer cache.Indexer) TestRunLister {
	return &testRunLister{indexer: indexer}
}

// List lists all TestRuns in the indexer.
func (s *testRunLister) List(selector labels.Selector) (ret []*v1alpha1.TestRun, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TestRun))
	})
	return ret, err
}

// TestRuns returns an object that can list and get TestRuns.
func (s *testRunLister) TestRuns(namespace string) TestRunNamespaceLister {
	return testRunNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TestRunNamespaceLister helps list and get TestRuns.
type TestRunNamespaceLister interface {
	// List lists all TestRuns in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TestRun, err error)
	// Get retrieves the TestRun from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TestRun, error)
	TestRunNamespaceListerExpansion
}

// testRunNamespaceLister implements the TestRunNamespaceLister
// interface.
type testRunNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TestRuns in the indexer for a given namespace.
func (s testRunNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TestRun, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TestRun))
	})
	return ret, err
}

// Get retrieves the TestRun from the indexer for a given namespace and name.
func (s testRunNamespaceLister) Get(name string) (*v1alpha1.TestRun, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("testrun"), name)
	}
	return obj.(*v1alpha1.TestRun), nil
}
