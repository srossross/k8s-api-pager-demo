package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/srossross/k8s-test-runner/pkg/client"
	"github.com/srossross/k8s-test-runner/pkg/run"
	factory "github.com/srossross/k8s-test-runner/pkg/informers/externalversions"
)

var (
	VERSION string
	BUILD_TIME string

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

	// TODO: add proper linker flags
	log.Printf("Test controller version: %s", VERSION)
	log.Printf("               Built on: %s", BUILD_TIME)

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

	testRunInformer := run.NewTestRunInformer(sharedFactory, queue)

	testInformer := run.NewTestInformer(sharedFactory, queue)

	podInformer := run.NewPodInformer(sharedFactory, queue)


	// start the informer. This will cause it to begin receiving updates from
	// the configured API server and firing event handlers in response.
	sharedFactory.Start(stopCh)
	log.Printf("Started informer factory.")

	// wait for the informe rcache to finish performing it's initial sync of
	// resources
	if !cache.WaitForCacheSync(stopCh, testRunInformer.HasSynced) {
		log.Fatalf("error waiting for testRunInformer cache to sync: %s", err.Error())
	}

	if !cache.WaitForCacheSync(stopCh, testInformer.HasSynced) {
		log.Fatalf("error waiting for testInformer cache to sync: %s", err.Error())
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
	queue.Add("Reconcile")
}

func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
