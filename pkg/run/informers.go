package run

import (
  "reflect"
  "log"

  workqueue "k8s.io/client-go/util/workqueue"
  cache "k8s.io/client-go/tools/cache"
  factory "github.com/srossross/k8s-test-controller/pkg/informers/externalversions"
)

func NewTestRunInformer(sharedFactory factory.SharedInformerFactory, queue workqueue.RateLimitingInterface) cache.SharedIndexInformer {

  informer := sharedFactory.Srossross().V1alpha1().TestRuns().Informer()
	// we add a new event handler, watching for changes to API resources.
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(cur interface{}) { queue.Add("Reconcile") } ,
			UpdateFunc: func(old, cur interface{}) {
				if !reflect.DeepEqual(old, cur) {
					queue.Add("Reconcile")
				}
			},
			DeleteFunc: func(cur interface{}) { queue.Add("Reconcile")} ,
		},
	)

  return informer
}

func NewTestInformer(sharedFactory factory.SharedInformerFactory, queue workqueue.RateLimitingInterface) cache.SharedIndexInformer  {
  testInformer := sharedFactory.Srossross().V1alpha1().Tests().Informer()
	// we add a new event handler, watching for changes to API resources.
	testInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(cur interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cur)
				if err != nil {
					log.Fatalf("Error in DeletionHandlingMetaNamespaceKeyFunc %v", err.Error())
				}
				log.Printf("Test %v Added (not triggering reconsile loop)", key)
			},
			UpdateFunc: func(old, cur interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cur)
				if err != nil {
					log.Fatalf("Error in DeletionHandlingMetaNamespaceKeyFunc %v", err.Error())
				}
				log.Printf("Test %v Updated (not triggering reconsile loop)", key)
			},
			DeleteFunc: func(cur interface{}) {
				key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(cur)
				if err != nil {
					log.Fatalf("Error in DeletionHandlingMetaNamespaceKeyFunc %v", err.Error())
				}
				log.Printf("Test %v Deleted (not triggering reconsile loop)", key)
			},
		},
	)
  return testInformer
}

func NewPodInformer(sharedFactory factory.SharedInformerFactory, queue workqueue.RateLimitingInterface) cache.SharedIndexInformer {

  podInformer := GetPodInformer(sharedFactory)

  podInformer.AddEventHandler(
    cache.ResourceEventHandlerFuncs{
      AddFunc: func(cur interface{}) { queue.Add("Reconcile") } ,
      UpdateFunc: func(old, cur interface{}) {
        if !reflect.DeepEqual(old, cur) {
          queue.Add("Reconcile")
        }
      },
      DeleteFunc: func(cur interface{}) { queue.Add("Reconcile") } ,
    },
  )
  return podInformer
}
