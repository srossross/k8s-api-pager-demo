package pager

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRun struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   TestRunSpec
	Status TestRunStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRunList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []TestRun
}

type TestRunSpec struct {
	Message string
}

type TestRunStatus struct {
	Sent bool
}
