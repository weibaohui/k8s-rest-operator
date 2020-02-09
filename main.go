package main

import (
	"cmit.com/paas/k8s/rest-operator/pkg/k8s"
	"cmit.com/paas/k8s/rest-operator/pkg/webservice"
)

func main() {
	go k8s.Start()
	go webservice.Start()
	select {}
}
