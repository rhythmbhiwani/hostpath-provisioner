package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"sigs.k8s.io/sig-storage-lib-external-provisioner/v7/controller"
)

const (
	provisionerName = "rhythmbhiwani.in/hostpath"
)

type hostPathProvisioner struct {
	client kubernetes.Interface
}

func NewHostPathProvisioner(client kubernetes.Interface) controller.Provisioner {
	return &hostPathProvisioner{client: client}
}

func (p *hostPathProvisioner) Provision(ctx context.Context, options controller.ProvisionOptions) (*corev1.PersistentVolume, controller.ProvisioningState, error) {
	pvPath := filepath.Join("/mnt/data", options.PVName)
	err := os.MkdirAll(pvPath, 0750)
	if err != nil {
		return nil, controller.ProvisioningFinished, fmt.Errorf("failed to create hostpath directory: %s", err)
	}

	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name: options.PVName,
		},
		Spec: corev1.PersistentVolumeSpec{
			PersistentVolumeReclaimPolicy: corev1.PersistentVolumeReclaimPolicy(options.StorageClass.Parameters["reclaimPolicy"]),
			AccessModes:                   options.PVC.Spec.AccessModes,
			Capacity:                      options.PVC.Spec.Resources.Requests,
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: pvPath,
				},
			},
		},
	}

	return pv, controller.ProvisioningFinished, nil
}

func (p *hostPathProvisioner) Delete(ctx context.Context, volume *corev1.PersistentVolume) error {
	pvPath := volume.Spec.HostPath.Path
	return os.RemoveAll(pvPath)
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	provisioner := NewHostPathProvisioner(clientset)
	pc := controller.NewProvisionController(clientset, provisionerName, provisioner, controller.LeaderElection(false))
	pc.Run(context.Background())
}