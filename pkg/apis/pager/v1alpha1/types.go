package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient=true
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=alerts

type Alert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertSpec   `json:"spec,omitempty"`
	Status AlertStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AlertList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Alert `json:"items"`
}

type AlertSpec struct {
	Message string `json:"message"`
}

type AlertStatus struct {
	Sent bool `json:"sent"`
}

// +genclient=true
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=alerts

type TestRunner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestRunnerSpec   `json:"spec,omitempty"`
	Status TestRunnerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestRunnerList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []TestRunner `json:"items"`
}

type TestRunnerSpec struct {
	Message string `json:"message"`
}

type TestRunnerStatus struct {
	Sent bool `json:"sent"`
}
