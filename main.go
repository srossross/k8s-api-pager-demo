package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/munnerz/k8s-api-pager-demo/pkg/client"
	"github.com/munnerz/k8s-api-pager-demo/pkg/run"
	factory "github.com/munnerz/k8s-api-pager-demo/pkg/informers/externalversions"
)

var (
	// apiserverURL is the URL of the API server to connect to
	kubeconfig = flag.String("kubeconfig", "", "Path to a kubeconfig file")
	// pushbulletToken is the pushbullet API token to use
	// pushbulletToken = flag.String("pushbullet-token", "", "the api token to use to send pushbullet messages")

	// queue is a queue of resources to be processed. It performs exponential
	// backoff rate limiting, with a minimum retry period of 5 seconds and a
	// maximum of 1 minute.
	rateLimiter = workqueue.NewItemExponentialFailureRateLimiter(time.Second*5, time.Minute)
	queue = workqueue.NewRateLimitingQueue(rateLimiter)

	config *rest.Config
	// stopCh can be used to stop all the informer, as well as control loops
	// within the application.
	stopCh = make(chan struct{})

	// sharedFactory is a shared informer factory that is used a a cache for
	// items in the API server. It saves each informer listing and watching the
	// same resources independently of each other, thus providing more up to
	// date results with less 'effort'
	sharedFactory factory.SharedInformerFactory

	// cl is a Kubernetes API client for our custom resource definition type
	cl client.Interface

	// pb is the pushbullet client to use to send alerts
	// pb *pushbullet.Pushbullet
)

func main() {
	flag.Parse()

	var err error

	config, err = GetClientConfig(*kubeconfig)

	if err != nil {
		log.Fatalf("error creating config: %s", err.Error())
	}

	// create an instance of our own API client
	cl, err = client.NewForConfig(config)

	if err != nil {
		log.Fatalf("error creating api client: %s", err.Error())
	}

	log.Printf("Created Kubernetes client.")

	// we use a shared informer from the informer factory, to save calls to the
	// API as we grow our application and so state is consistent between our
	// control loops. We set a resync period of 30 seconds, in case any
	// create/replace/update/delete operations are missed when watching
	sharedFactory = factory.NewSharedInformerFactory(cl, time.Second*30)

	informer := sharedFactory.Pager().V1alpha1().TestRuns().Informer()
	// we add a new event handler, watching for changes to API resources.
	informer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: enqueue,
			UpdateFunc: func(old, cur interface{}) {
				if !reflect.DeepEqual(old, cur) {
					enqueue(cur)
				}
			},
			DeleteFunc: enqueue,
		},
	)

	testInformer := sharedFactory.Pager().V1alpha1().Tests().Informer()
	// we add a new event handler, watching for changes to API resources.
	testInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: enqueue,
			UpdateFunc: func(old, cur interface{}) {
				if !reflect.DeepEqual(old, cur) {
					enqueue(cur)
				}
			},
			DeleteFunc: enqueue,
		},
	)

	podInformer := run.GetPodInformer(sharedFactory)

	podInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: enqueue,
			UpdateFunc: func(old, cur interface{}) {
				if !reflect.DeepEqual(old, cur) {
					enqueue(cur)
				}
			},
			DeleteFunc: enqueue,
		},
	)

	// TODO: remove comments
	// run1, err := cl.PagerV1alpha1().TestRuns("default").Get("test-run1", metav1.GetOptions{})
	// // obj, err := sharedFactory.Pager().V1alpha1().TestRuns().Lister().TestRuns("default").Get("cncf-test-runner")
	// if err != nil {
	// 	panic(err)
	// 	// return
	// }
	//
	// log.Printf("Namespace: '%v'", run1.Namespace)
	//
	// selector, err := metav1.LabelSelectorAsSelector(run1.Spec.Selector)
	//
	// if err != nil {
	// 	panic(err)
	// 	// return
	// }
	//
	// log.Printf("Selector: '%v'", selector)
	// log.Printf("--")
	//
	//
	// matchingTests, err := cl.PagerV1alpha1().Tests(run1.Namespace).List(metav1.ListOptions{
	// 	LabelSelector: selector.String(),
	// })
	// // matchingTests, err := cl.PagerV1alpha1().Tests(run1.Namespace).List(metav1.ListOptions{})
	//
	// if err != nil {
	// 	log.Fatalf("Could not get tests: %s", err.Error())
	// 	panic("can not get tests", )
	// 	// return
	// }
	//
	// for _, matchingTest := range matchingTests.Items {
	// 	run.RunTest(cl, run1, matchingTest)
	// 	// log.Printf("matchingTest %v", matchingTest.Spec.Template)
	// }
	//
	//
	// panic("!all good!")

	// start the informer. This will cause it to begin receiving updates from
	// the configured API server and firing event handlers in response.
	sharedFactory.Start(stopCh)
	log.Printf("Started informer factory.")

	// wait for the informe rcache to finish performing it's initial sync of
	// resources
	if !cache.WaitForCacheSync(stopCh, informer.HasSynced) {
		log.Fatalf("error waiting for informer cache to sync: %s", err.Error())
	}

	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced) {
		log.Fatalf("error waiting for podInformer cache to sync: %s", err.Error())
	}

	log.Printf("Finished populating shared informer cache.")
	// here we start just one worker reading objects off the queue. If you
	// wanted to parallelize this, you could start many instances of the worker
	// function, then ensure your application handles concurrency correctly.
	work()
}


func work() {
	for {
		// we read a message off the queue
		key, shutdown := queue.Get()

		// if the queue has been shut down, we should exit the work queue here
		if shutdown {
			stopCh <- struct{}{}
			return
		}

		// convert the queue item into a string. If it's not a string, we'll
		// simply discard it as invalid data and log a message.
		var strKey string
		var ok bool
		if strKey, ok = key.(string); !ok {
			runtime.HandleError(fmt.Errorf("key in queue should be of type string but got %T. discarding", key))
			return
		}

		// we define a function here to process a queue item, so that we can
		// use 'defer' to make sure the message is marked as Done on the queue
		func(key string) {
			defer queue.Done(key)
			log.Printf("Read item '%s' off workqueue.", key)
			run.Reconcile(sharedFactory, cl)
			queue.Forget(key)
		}(strKey)
	}
}

// enqueue will add an object 'obj' into the workqueue. The object being added
// must be of type metav1.Object, metav1.ObjectAccessor or cache.ExplicitKey.
func enqueue(obj interface{}) {
	// DeletionHandlingMetaNamespaceKeyFunc will convert an object into a
	// 'namespace/name' string. We do this because our item may be processed
	// much later than now, and so we want to ensure it gets a fresh copy of
	// the resource when it starts. Also, this allows us to keep adding the
	// same item into the work queue without duplicates building up.
	// key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	// if err != nil {
	// 	runtime.HandleError(fmt.Errorf("error obtaining key for object being enqueue: %s", err.Error()))
	// 	return
	// }
	// add the item to the queue
	queue.Add("Reconcile")
}

func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
