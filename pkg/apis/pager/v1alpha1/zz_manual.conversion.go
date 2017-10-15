package v1alpha1

import (
	pager "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager"
	conversion "k8s.io/apimachinery/pkg/conversion"
	// runtime "k8s.io/apimachinery/pkg/runtime"
	// unsafe "unsafe"
)


func Convert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(in *TestRunSpec, out *pager.TestRunSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_TestRunSpec_To_pager_TestRunSpec(in, out, s)
}
