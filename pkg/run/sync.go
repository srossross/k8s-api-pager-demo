package run

import (
  // "log"
  // "fmt"
  // v1 "k8s.io/api/core/v1"
  // metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  client "github.com/munnerz/k8s-api-pager-demo/pkg/client"
  v1alpha1 "github.com/munnerz/k8s-api-pager-demo/pkg/apis/pager/v1alpha1"
)

// sync will attempt to 'Sync' an alert resource. It checks to see if the alert
// has already been sent, and if not will send it and update the resource
// accordingly. This method is called whenever this controller starts, and
// whenever the resource changes, and also periodically every resyncPeriod.
func Sync(cl *client.Interface, tr *v1alpha1.TestRun) error {

  return nil
	// If this message has already been sent, we exit with no error
	// if tr.Status.Sent {
	// 	log.Printf("Skipping already Sent alert '%s/%s'", tr.Namespace, tr.Name)
	// 	return nil
	// }

	// // create our note instance
	// // note := requests.NewNote()
	// log.Printf(fmt.Sprintf("Kubernetes alert for %s/%s", tr.Namespace, tr.Name))
	// log.Printf(tr.Spec.Message)
  //
	// // send the note. If an error occurs here, we return an error which will
	// // cause the calling function to re-queue the item to be tried again later.
	// // if _, err := pb.PostPushesNote(note); err != nil {
	// // 	return fmt.Errorf("error sending pushbullet message: %s", err.Error())
	// // }
	// log.Printf("Sent pushbullet note!")
  //
	// // as we've sent the note, we will update the resource accordingly.
	// // if this request fails, this item will be requeued and a second alert
	// // will be sent. It's therefore worth noting that this control loop will
	// // send you *at least one* alert, and not *at most one*.
	// tr.Status.Sent = true
	// if _, err := (*cl).PagerV1alpha1().TestRuns(tr.Namespace).Update(tr); err != nil {
	// 	return fmt.Errorf("error saving update to pager TestRun resource: %s", err.Error())
	// }
	// log.Printf("Finished saving update to pager TestRun resource '%s/%s'", tr.Namespace, tr.Name)
  //
	// // we didn't encounter any errors, so we return nil to allow the callee
	// // to 'forget' this item from the queue altogether.
	// return nil
}
