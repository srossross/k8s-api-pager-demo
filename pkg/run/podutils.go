package run

import (
	// "flag"
	// "fmt"
	// "log"
	"time"

	pkgRuntime "k8s.io/apimachinery/pkg/runtime"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
  listerv1 "k8s.io/client-go/listers/core/v1"
	watch "k8s.io/apimachinery/pkg/watch"

	"k8s.io/client-go/tools/cache"

	"github.com/srossross/k8s-test-controller/pkg/client"
	factory "github.com/srossross/k8s-test-controller/pkg/informers/externalversions"
)


func newPodInformer(client client.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (pkgRuntime.Object, error) {
				return client.CoreV1().Pods(metav1.NamespaceAll).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return client.CoreV1().Pods(metav1.NamespaceAll).Watch(options)
			},
		},
		&corev1.Pod{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	return sharedIndexInformer
}

func GetPodInformer(sharedFactory factory.SharedInformerFactory) cache.SharedIndexInformer {
  return sharedFactory.InformerFor(&corev1.Pod{}, newPodInformer)
}

func GetPodLister(sharedFactory factory.SharedInformerFactory) listerv1.PodLister {
	return listerv1.NewPodLister(GetPodInformer(sharedFactory).GetIndexer())
}
