package run

import (
  "log"
  "fmt"
  "time"
  v1 "k8s.io/api/core/v1"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  client "github.com/srossross/k8s-test-runner/pkg/client"
  v1alpha1 "github.com/srossross/k8s-test-runner/pkg/apis/pager/v1alpha1"
)

func CreateTestRunEvent(cl client.Interface, testRun *v1alpha1.TestRun, test *v1alpha1.Test, Reason string, Message string) (error) {

  objectReference := v1.ObjectReference{
    // FIXME: not sure why testRun.Kind is empty
    Kind: "TestRun",
    Namespace: testRun.Namespace,
    Name: testRun.Name,
    UID: testRun.UID,
     // FIXME: not sure why testRun.APIVersion is empty
    APIVersion: "pager.k8s.co/v1alpha1",
    ResourceVersion: testRun.ResourceVersion,
  }

  event := v1.Event{
    metav1.TypeMeta{},
    metav1.ObjectMeta{
      GenerateName: "test-run-event",
    },
    objectReference,
    Reason,
    Message,
    // FIXME: populate with real values
    v1.EventSource{
      "test-runner",
      "hostname",
    },
    metav1.Time{time.Now()},
    metav1.Time{time.Now()},
    1,
    "Normal",
  }

  _, err := cl.CoreV1().Events(test.Namespace).Create(&event)
  if err != nil {
    log.Printf("Error Creating event while starting test %v", err)
    return err
  }

  return nil

}

func CreateTestRunEventStart(cl client.Interface, testRun *v1alpha1.TestRun, test *v1alpha1.Test) (error) {
  return CreateTestRunEvent(
    cl, testRun, test,
    "TestStart",
    fmt.Sprintf("Starting test %s", test.Name),
  )
}

func CreateTestPod(cl client.Interface, testRun *v1alpha1.TestRun, test *v1alpha1.Test) (error){
	log.Printf("RunTest")
	log.Printf("Test '%v'", test)

  err := CreateTestRunEventStart(cl, testRun, test)
  if err != nil {
    return err
  }

  // TODO: use template labels and Annotations
  pod := &v1.Pod{
    metav1.TypeMeta{},
    metav1.ObjectMeta{
      GenerateName: test.Name,
      Namespace: test.Namespace,
      Annotations: map[string]string{
        "test-run": testRun.Name,
        "test": test.Name,
        "test-fullname": fmt.Sprintf("%s/%s", testRun.Name, test.Name),
      },
    },
    test.Spec.Template.Spec,
    v1.PodStatus{},
  }


  result, err := cl.CoreV1().Pods(test.Namespace).Create(pod)
  if err != nil {
    CreateTestRunEvent(
      cl, testRun, test, "PodCreationFailure",
      fmt.Sprintf("Could not create pod for test %s", test.Name),
    )
    log.Printf("Error Creating pod while starting test %v", err)

    return err
  }

  log.Printf("Pod '%v'", result)

  return nil
}
