package webservice

import (
	"cmit.com/paas/k8s/rest-operator/pkg/k8s"
	"cmit.com/paas/k8s/rest-operator/pkg/utils"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func Start() {
	container := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Route(ws.GET("/ports").
		To(ports).
		Produces(restful.MIME_JSON))
	container.Add(ws)
	log.Fatal(http.ListenAndServe(":9999", container))
}

// GET /ports
func ports(request *restful.Request, response *restful.Response) {
	deployment, err := k8s.GetWatcher().Deployments.Lister().Deployments("docker").Get("compose")
	if err != nil {
		response.WriteAsJson(err)
		return
	}
	labels := deployment.Labels
	labels["x"] = "nn"
	deployment.Labels = labels

	update, err := utils.NewK8sClient().GetKubeClient().AppsV1().Deployments("docker").Update(deployment)
	if err != nil {
		response.WriteAsJson(err)
		return
	}
	response.WriteEntity(update)
}
