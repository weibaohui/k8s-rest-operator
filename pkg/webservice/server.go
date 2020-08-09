package webservice

import (
	"context"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"cmit.com/paas/k8s/rest-operator/pkg/k8s"
	"cmit.com/paas/k8s/rest-operator/pkg/utils"
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

	update, err := utils.NewK8sClient().GetKubeClient().AppsV1().Deployments("docker").Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	if err != nil {
		response.WriteAsJson(err)
		return
	}
	response.WriteEntity(update)

}
