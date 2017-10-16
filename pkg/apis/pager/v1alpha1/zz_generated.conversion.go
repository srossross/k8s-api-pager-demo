// +build !ignore_autogenerated

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	pager "github.com/srossross/k8s-test-runner/pkg/apis/pager"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_TestRun_To_pager_TestRun,
		Convert_pager_TestRun_To_v1alpha1_TestRun,
		Convert_v1alpha1_TestRunList_To_pager_TestRunList,
		Convert_pager_TestRunList_To_v1alpha1_TestRunList,
		Convert_v1alpha1_TestRunSpec_To_pager_TestRunSpec,
		Convert_pager_TestRunSpec_To_v1alpha1_TestRunSpec,
		Convert_v1alpha1_TestRunStatus_To_pager_TestRunStatus,
		Convert_pager_TestRunStatus_To_v1alpha1_TestRunStatus,
	)
}

func autoConvert_v1alpha1_TestRun_To_pager_TestRun(in *TestRun, out *pager.TestRun, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_TestRunStatus_To_pager_TestRunStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_TestRun_To_pager_TestRun is an autogenerated conversion function.
func Convert_v1alpha1_TestRun_To_pager_TestRun(in *TestRun, out *pager.TestRun, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestRun_To_pager_TestRun(in, out, s)
}

func autoConvert_pager_TestRun_To_v1alpha1_TestRun(in *pager.TestRun, out *TestRun, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_pager_TestRunSpec_To_v1alpha1_TestRunSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_pager_TestRunStatus_To_v1alpha1_TestRunStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_pager_TestRun_To_v1alpha1_TestRun is an autogenerated conversion function.
func Convert_pager_TestRun_To_v1alpha1_TestRun(in *pager.TestRun, out *TestRun, s conversion.Scope) error {
	return autoConvert_pager_TestRun_To_v1alpha1_TestRun(in, out, s)
}

func autoConvert_v1alpha1_TestRunList_To_pager_TestRunList(in *TestRunList, out *pager.TestRunList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]pager.TestRun)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_TestRunList_To_pager_TestRunList is an autogenerated conversion function.
func Convert_v1alpha1_TestRunList_To_pager_TestRunList(in *TestRunList, out *pager.TestRunList, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestRunList_To_pager_TestRunList(in, out, s)
}

func autoConvert_pager_TestRunList_To_v1alpha1_TestRunList(in *pager.TestRunList, out *TestRunList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Items == nil {
		out.Items = make([]TestRun, 0)
	} else {
		out.Items = *(*[]TestRun)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_pager_TestRunList_To_v1alpha1_TestRunList is an autogenerated conversion function.
func Convert_pager_TestRunList_To_v1alpha1_TestRunList(in *pager.TestRunList, out *TestRunList, s conversion.Scope) error {
	return autoConvert_pager_TestRunList_To_v1alpha1_TestRunList(in, out, s)
}

func autoConvert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(in *TestRunSpec, out *pager.TestRunSpec, s conversion.Scope) error {
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxJobs = in.MaxJobs
	return nil
}

// Convert_v1alpha1_TestRunSpec_To_pager_TestRunSpec is an autogenerated conversion function.
func Convert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(in *TestRunSpec, out *pager.TestRunSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(in, out, s)
}

func autoConvert_pager_TestRunSpec_To_v1alpha1_TestRunSpec(in *pager.TestRunSpec, out *TestRunSpec, s conversion.Scope) error {
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxJobs = in.MaxJobs
	return nil
}

// Convert_pager_TestRunSpec_To_v1alpha1_TestRunSpec is an autogenerated conversion function.
func Convert_pager_TestRunSpec_To_v1alpha1_TestRunSpec(in *pager.TestRunSpec, out *TestRunSpec, s conversion.Scope) error {
	return autoConvert_pager_TestRunSpec_To_v1alpha1_TestRunSpec(in, out, s)
}

func autoConvert_v1alpha1_TestRunStatus_To_pager_TestRunStatus(in *TestRunStatus, out *pager.TestRunStatus, s conversion.Scope) error {
	out.Status = in.Status
	out.Message = in.Message
	out.Success = in.Success
	return nil
}

// Convert_v1alpha1_TestRunStatus_To_pager_TestRunStatus is an autogenerated conversion function.
func Convert_v1alpha1_TestRunStatus_To_pager_TestRunStatus(in *TestRunStatus, out *pager.TestRunStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestRunStatus_To_pager_TestRunStatus(in, out, s)
}

func autoConvert_pager_TestRunStatus_To_v1alpha1_TestRunStatus(in *pager.TestRunStatus, out *TestRunStatus, s conversion.Scope) error {
	out.Status = in.Status
	out.Message = in.Message
	out.Success = in.Success
	return nil
}

// Convert_pager_TestRunStatus_To_v1alpha1_TestRunStatus is an autogenerated conversion function.
func Convert_pager_TestRunStatus_To_v1alpha1_TestRunStatus(in *pager.TestRunStatus, out *TestRunStatus, s conversion.Scope) error {
	return autoConvert_pager_TestRunStatus_To_v1alpha1_TestRunStatus(in, out, s)
}
