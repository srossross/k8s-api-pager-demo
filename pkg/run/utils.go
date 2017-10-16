package run
import (
  "k8s.io/api/core/v1"
)

func testRunFilter(pods []*v1.Pod, testRunName string) []*v1.Pod {
  result := []*v1.Pod{}
  for _, pod := range pods {
    if pod.Annotations["test-run"] == testRunName {
      result = append(result, pod)
    }
      // if(!strings.HasPrefix(a[i], "foo_") && len(a[i]) <= 7) {
      //     nofoos = append(nofoos, a[i])
  }

  return result
}
