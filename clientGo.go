package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type ClientSet struct {
	clientset *kubernetes.Clientset
}

func (cs *ClientSet) getClientSet() {
	// Create a new Kubernetes client
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	//return clientset
	cs.clientset = clientset
}

type AgentMetaData struct {
	ShipID    string
	HarbourID string
	Ns        string
}

func (a *AgentMetaData) getPods(cs *ClientSet) ([]string, error) {
	ctx := context.TODO()
	pods, err := cs.clientset.CoreV1().Pods(a.Ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	var taurusPods []string
	for _, pod := range pods.Items {
		// Only select pods that do NOT have the annotation and match the labels
		if pod.Annotations["cluster-autoscaler.kubernetes.io/safe-to-evict"] == "" &&
			a.HarbourID == pod.Labels["BZM_HARBOR_ID"] &&
			a.ShipID == pod.Labels["BZM_SHIP_ID"] {
			fmt.Printf("Found taurus-cloud Pod: %s\n", pod.Name)
			taurusPods = append(taurusPods, pod.Name)
		}
	}
	return taurusPods, nil
}

func (a AgentMetaData) addAnnotations(cs *ClientSet, podNames []string) error {
	ctx := context.TODO()
	for _, podName := range podNames {
		patch := []byte(`{"metadata":{"annotations":{"cluster-autoscaler.kubernetes.io/safe-to-evict": "false"}}}`)
		_, err := cs.clientset.CoreV1().Pods(a.Ns).Patch(ctx, podName, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}

type annotationUpdate interface {
	getPods(cs *ClientSet) ([]string, error)
	addAnnotations(cs *ClientSet, podNames []string) error
}

func podUpdateAnnotaion(a annotationUpdate) {
	cs := &ClientSet{}
	cs.getClientSet()

	podList, err := a.getPods(cs)
	if err != nil {
		fmt.Printf("Error getting pods: %v\n", err)
		panic(err.Error())
	}
	// Add annotations to the pods
	err = a.addAnnotations(cs, podList)
	if err != nil {
		fmt.Printf("Error adding annotations: %v\n", err)
		panic(err.Error())
	}

}
