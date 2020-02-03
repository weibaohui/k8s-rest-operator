package handler

import (
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/handler/deploy"
	"cmit.com/paas/k8s/rest-operator/pkg/k8s/handler/pod"
	"cmit.com/paas/k8s/rest-operator/pkg/watcher"

	"k8s.io/client-go/informers"
)

func Register(f informers.SharedInformerFactory) error {
	r := &watcher.Register{
		Handlers: []watcher.HandlersRegister{
			deploy.Register,
			pod.Register,
		},
	}
	return r.Register(f)
}
