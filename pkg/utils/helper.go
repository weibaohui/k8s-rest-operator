package utils

import (
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	typeV1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"sync"
)

var cli kubernetes.Interface
var once = sync.Once{}

type helper struct {
	cli kubernetes.Interface
}

func NewK8sClient() *helper {
	once.Do(func() {
		cli = getClient()
	})

	return &helper{cli: cli}
}

func (h *helper) GetKubeClient() kubernetes.Interface {
	return h.cli
}
func (h *helper) RESTClient() rest.Interface {
	return h.cli.CoreV1().RESTClient()
}
func (h *helper) Pods(ns string) typeV1.PodInterface {
	return h.cli.CoreV1().Pods(ns)
}
func (h *helper) Services(ns string) typeV1.ServiceInterface {
	return h.cli.CoreV1().Services(ns)
}

func (h *helper) GetPod(ns, podName string) (*coreV1.Pod, error) {
	return h.Pods(ns).Get(podName, metaV1.GetOptions{})
}

func (h *helper) GetService(ns, svcName string) (*coreV1.Service, error) {
	return h.Services(ns).Get(svcName, metaV1.GetOptions{})
}
func (h *helper) IsServiceExists(ns, svcName string) bool {
	_, e := h.Services(ns).Get(svcName, metaV1.GetOptions{})
	if e == nil {
		return true
	}
	return false
}

func getClient() kubernetes.Interface {
	var config *rest.Config
	var err error
	if home := homeDir(); home != "" {
		kubeConfig := filepath.Join(home, ".kube", "config")
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
	} else {
		config, err = rest.InClusterConfig()
	}

	if err != nil {
		panic(err.Error())
	}
	cli, e := kubernetes.NewForConfig(config)
	if e != nil {
		panic(e.Error())
	}
	return cli

}

func homeDir() string {
	if s := os.Getenv("HOME"); s != "" {
		return s
	}
	return os.Getenv("USERPROFILE")
}
