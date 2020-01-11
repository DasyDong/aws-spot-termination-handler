package pkg

import (
	"log"
	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/drain"
)

var drainHelper *drain.Helper
var node *corev1.Node

//InitDrainer will ensure the kube has all the resources necessary to complete a drain
func InitDrainer(nodeName string) {
	if drainHelper == nil {
		drainHelper = getDrainHelper()
	}
	node = &corev1.Node{}
	var err error
	node, err = drainHelper.Client.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Couldn't get node %q: %s\n", nodeName, err.Error())
	}
	log.Printf("Successufully retrieved node: %s", node.Name)
}

//Drain will cordon the node and evict pods based on the config
func Drain(nodeName string) {
	if drainHelper == nil || node == nil {
		InitDrainer(nodeName)
	}
	var err error
	err = drain.RunCordonOrUncordon(drainHelper, node, true)
	if err != nil {
		log.Fatalf("Couldn't cordon node %q: %s\n", node, err.Error())
	}

	// Delete all pods on the node
	err = drain.RunNodeDrain(drainHelper, nodeName)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getDrainHelper() *drain.Helper {
	var clientset = &kubernetes.Clientset{}
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalln("Failed to create in-cluster config: ", err.Error())
	}

	// creates the clientset
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln("Failed to create kubernetes clientset: ", err.Error())
	}

	return &drain.Helper{
		Client:              clientset,
		Force:               true,
		GracePeriodSeconds:  1,
		IgnoreAllDaemonSets: true,
		DeleteLocalData:     true,
		Timeout:             time.Duration(30) * time.Second,
		Out:                 os.Stdout,
		ErrOut:              os.Stderr,
	}
}
