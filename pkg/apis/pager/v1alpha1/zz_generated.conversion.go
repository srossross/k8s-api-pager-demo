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
	pager "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager"
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
		Convert_v1alpha1_Alert_To_pager_Alert,
		Convert_pager_Alert_To_v1alpha1_Alert,
		Convert_v1alpha1_AlertList_To_pager_AlertList,
		Convert_pager_AlertList_To_v1alpha1_AlertList,
		Convert_v1alpha1_AlertSpec_To_pager_AlertSpec,
		Convert_pager_AlertSpec_To_v1alpha1_AlertSpec,
		Convert_v1alpha1_AlertStatus_To_pager_AlertStatus,
		Convert_pager_AlertStatus_To_v1alpha1_AlertStatus,
	)
}

func autoConvert_v1alpha1_Alert_To_pager_Alert(in *Alert, out *pager.Alert, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_AlertSpec_To_pager_AlertSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_AlertStatus_To_pager_AlertStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_Alert_To_pager_Alert is an autogenerated conversion function.
func Convert_v1alpha1_Alert_To_pager_Alert(in *Alert, out *pager.Alert, s conversion.Scope) error {
	return autoConvert_v1alpha1_Alert_To_pager_Alert(in, out, s)
}

func autoConvert_pager_Alert_To_v1alpha1_Alert(in *pager.Alert, out *Alert, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_pager_AlertSpec_To_v1alpha1_AlertSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_pager_AlertStatus_To_v1alpha1_AlertStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_pager_Alert_To_v1alpha1_Alert is an autogenerated conversion function.
func Convert_pager_Alert_To_v1alpha1_Alert(in *pager.Alert, out *Alert, s conversion.Scope) error {
	return autoConvert_pager_Alert_To_v1alpha1_Alert(in, out, s)
}

func autoConvert_v1alpha1_AlertList_To_pager_AlertList(in *AlertList, out *pager.AlertList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]pager.Alert)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_AlertList_To_pager_AlertList is an autogenerated conversion function.
func Convert_v1alpha1_AlertList_To_pager_AlertList(in *AlertList, out *pager.AlertList, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertList_To_pager_AlertList(in, out, s)
}

func autoConvert_pager_AlertList_To_v1alpha1_AlertList(in *pager.AlertList, out *AlertList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if in.Items == nil {
		out.Items = make([]Alert, 0)
	} else {
		out.Items = *(*[]Alert)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_pager_AlertList_To_v1alpha1_AlertList is an autogenerated conversion function.
func Convert_pager_AlertList_To_v1alpha1_AlertList(in *pager.AlertList, out *AlertList, s conversion.Scope) error {
	return autoConvert_pager_AlertList_To_v1alpha1_AlertList(in, out, s)
}

func autoConvert_v1alpha1_AlertSpec_To_pager_AlertSpec(in *AlertSpec, out *pager.AlertSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_v1alpha1_AlertSpec_To_pager_AlertSpec is an autogenerated conversion function.
func Convert_v1alpha1_AlertSpec_To_pager_AlertSpec(in *AlertSpec, out *pager.AlertSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertSpec_To_pager_AlertSpec(in, out, s)
}

func autoConvert_pager_AlertSpec_To_v1alpha1_AlertSpec(in *pager.AlertSpec, out *AlertSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_pager_AlertSpec_To_v1alpha1_AlertSpec is an autogenerated conversion function.
func Convert_pager_AlertSpec_To_v1alpha1_AlertSpec(in *pager.AlertSpec, out *AlertSpec, s conversion.Scope) error {
	return autoConvert_pager_AlertSpec_To_v1alpha1_AlertSpec(in, out, s)
}

func autoConvert_v1alpha1_AlertStatus_To_pager_AlertStatus(in *AlertStatus, out *pager.AlertStatus, s conversion.Scope) error {
	out.Sent = in.Sent
	return nil
}

// Convert_v1alpha1_AlertStatus_To_pager_AlertStatus is an autogenerated conversion function.
func Convert_v1alpha1_AlertStatus_To_pager_AlertStatus(in *AlertStatus, out *pager.AlertStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_AlertStatus_To_pager_AlertStatus(in, out, s)
}

func autoConvert_pager_AlertStatus_To_v1alpha1_AlertStatus(in *pager.AlertStatus, out *AlertStatus, s conversion.Scope) error {
	out.Sent = in.Sent
	return nil
}

// Convert_pager_AlertStatus_To_v1alpha1_AlertStatus is an autogenerated conversion function.
func Convert_pager_AlertStatus_To_v1alpha1_AlertStatus(in *pager.AlertStatus, out *AlertStatus, s conversion.Scope) error {
	return autoConvert_pager_AlertStatus_To_v1alpha1_AlertStatus(in, out, s)
}
