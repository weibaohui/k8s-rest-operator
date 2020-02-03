package k8s

import (
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/handler"
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/starter"
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/starter/deployment"
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/starter/pod"
	"cmit.com/paas/k8s/rest-operator/pkg/utils"
	"cmit.com/paas/k8s/rest-operator/pkg/watcher"

	"k8s.io/client-go/informers"
	"time"
)

var w watcher.Watcher

func GetWatcher() *watcher.Watcher {
	return &w
}
func Start() {
	stop := make(chan struct{})

	var KubeClient = utils.NewK8sClient().GetKubeClient()
	factory := informers.NewSharedInformerFactory(KubeClient, time.Hour*2)

	w = watcher.Watcher{
		Factory:      factory,
		DaemonSets:   factory.Apps().V1().DaemonSets(),
		Deployments:  factory.Apps().V1().Deployments(),
		ReplicaSets:  factory.Apps().V1().ReplicaSets(),
		StatefulSets: factory.Apps().V1().StatefulSets(),
		ConfigMaps:   factory.Core().V1().ConfigMaps(),
		Endpoints:    factory.Core().V1().Endpoints(),
		Namespaces:   factory.Core().V1().Namespaces(),
		Pods:         factory.Core().V1().Pods(),
		Services:     factory.Core().V1().Services(),
		Ingresses:    factory.Extensions().V1beta1().Ingresses(),
		Nodes:        factory.Core().V1().Nodes(),
	}

	//
	w.Factory.Start(stop)

	w.Factory.WaitForCacheSync(stop)

	starter.StartThenSync(factory, stop, &deployment.Starter{}, &pod.Starter{})

	handler.Register(factory)

	select {}
}
