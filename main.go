package main

import (
	"flag"

	"app.io/pkg/k8s"
	k8sDeployment "app.io/pkg/k8s/deployment"
	apiv1 "k8s.io/api/core/v1"
)

var kubeconfig *string

func init() {
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()
}

func main() {

	client, err := k8s.GetKubernetesClient(*kubeconfig)
	if err != nil {
		panic(err)
	}

	deploymentName := "webserver"
	nameSpace := "default"
	k8sDeployment.List(*client, apiv1.NamespaceAll)
	k8sDeployment.PerformDeployment(*client, nameSpace, deploymentName, "nginx:alpine", 1)
	k8sDeployment.Scale(*client, nameSpace, deploymentName, 5)
	k8sDeployment.ChangeImage(*client, nameSpace, deploymentName, "nginx:alpine")
	k8sDeployment.Delete(*client, nameSpace, deploymentName)

}
