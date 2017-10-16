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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	internalinterfaces "github.com//pkg/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Tests returns a TestInformer.
	Tests() TestInformer
	// TestRuns returns a TestRunInformer.
	TestRuns() TestRunInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// Tests returns a TestInformer.
func (v *version) Tests() TestInformer {
	return &testInformer{factory: v.SharedInformerFactory}
}

// TestRuns returns a TestRunInformer.
func (v *version) TestRuns() TestRunInformer {
	return &testRunInformer{factory: v.SharedInformerFactory}
}
