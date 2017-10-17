
package run
import (
	"fmt"
	"log"

	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/util/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	client "github.com/srossross/k8s-test-controller/pkg/client"
	factory "github.com/srossross/k8s-test-controller/pkg/informers/externalversions"

	v1alpha1 "github.com/srossross/k8s-test-controller/pkg/apis/pager/v1alpha1"
)

var (
	StatusComplete = "Complete"
)

func TestRunner(sharedFactory factory.SharedInformerFactory, cl client.Interface, testRun *v1alpha1.TestRun) error {
	if testRun.Status.Status == StatusComplete {
		log.Printf("TestRun: %v is already Complete - Skipping", testRun.Name)
		return nil
	}

	log.Printf("Reconcile TestRun: %v", testRun.Name)

	selector, err := metav1.LabelSelectorAsSelector(testRun.Spec.Selector)

	if err != nil {
		return fmt.Errorf("error getting test selector: %s", err.Error())
	}

	tests, err := sharedFactory.Srossross().V1alpha1().Tests().Lister().Tests(testRun.Namespace).List(selector)

	if err != nil {
		return fmt.Errorf("error getting list of tests: %s", err.Error())
	}

	log.Printf("  Test Count: %v", len(tests))

	pods, err := GetPodLister(sharedFactory).Pods(testRun.Namespace).List(labels.Everything())
	if err != nil {
		return fmt.Errorf("Error getting list of pods: %s", err.Error())
	}

	pods = testRunFilter(pods, testRun.Name)

	log.Printf("  Pod Count: %v", len(pods))

	podMap := make(map[string]*corev1.Pod)
	for _, pod := range pods {
		log.Printf("Pod: %v", pod.Name)
		podMap[pod.Annotations["test"]] = pod
	}


	// FIXME: should be a default in the schema ...
	var JobsSlots int
	if testRun.Spec.MaxJobs > 0 {
		JobsSlots = testRun.Spec.MaxJobs
	} else {
		JobsSlots = 1
	}
	log.Printf("JobsSlots %v", JobsSlots)

	completedCount := 0
	failCount := 0
	for _, test := range tests {
		if JobsSlots <= 0 {
			log.Printf("  No more jobs allowed. moving on...", test.Name)
			return nil
		}

		log.Printf("  Test: %v", test.Name)

		if pod, ok := podMap[test.Name]; ok {
			log.Printf("pod '%v' exists - Status: %v", pod.Name, pod.Status.Phase)
			switch pod.Status.Phase {
			case "Succeeded":
				completedCount += 1
				continue
			case "Failed":
				completedCount += 1
				failCount += 1
				continue
			case "Unknown":
				completedCount += 1
				failCount += 1
				continue
			// These are running and taking up a job slot!
			case "Pending":
				JobsSlots -= 1
				continue
			case "Running":
				JobsSlots -= 1
				continue
			}
		} else {
			err = CreateTestPod(cl, testRun, test)

			if err != nil {
				return err
			}

			JobsSlots -= 1
		}
	}

	if completedCount == len(tests) {

		testRun.Status.Status = StatusComplete
		testRun.Status.Success = failCount == 0
		testRun.Status.Message = fmt.Sprintf("Ran %v tests, %v failures", completedCount, failCount)

		log.Printf("Saving '%v/%v'", testRun.Namespace, testRun.Name)
		if _, err := cl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRun); err != nil {
			return err
		}
		log.Printf("We are done here %v tests completed", completedCount)
	} else {
		log.Printf("Completed %v of %v tests", completedCount, len(tests))
	}

	return nil
}

func Reconcile(sharedFactory factory.SharedInformerFactory, cl client.Interface) {

	lister := sharedFactory.Srossross().V1alpha1().TestRuns().Lister()
  runs, err := lister.TestRuns(metav1.NamespaceAll).List(labels.Everything())

  if err != nil {
    runtime.HandleError(fmt.Errorf("error getting list of testruns: %s", err.Error()))
    return
  }


	// FIXME: make our informers more efficient
	// informer should queue the testrun we care about instead of looping...
  for _, testRun := range runs {

		// FIXME: why is this not set?
		if len(testRun.Namespace) == 0 {
			testRun.Namespace = "default"
		}

		err := TestRunner(sharedFactory, cl, testRun)

		if err != nil {
			testRun.Status.Status = StatusComplete
			testRun.Status.Success = false
			testRun.Status.Message = fmt.Sprintf("Critical error during test run (%v)", err.Error())

			log.Printf("Saving Error state for '%v/%v'", testRun.Namespace, testRun.Name)
			if _, err := cl.SrossrossV1alpha1().TestRuns(testRun.Namespace).Update(testRun); err != nil {
				runtime.HandleError(fmt.Errorf("Error saving update to testrun (This could cause an infinite Reconcile loop): %s", err.Error()))

			}

		}
	}

}
