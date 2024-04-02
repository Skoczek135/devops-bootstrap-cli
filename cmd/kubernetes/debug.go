/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

var ns string

// debugCmd deploys a debug pod to the Kubernetes cluster
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Deploys a lightweight debug pod to the Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := client.New(config.GetConfigOrDie(), client.Options{})
		if err != nil {
			panic(err)
		}

		pod := &corev1.Pod{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "debug",
				Namespace: ns,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:    "debug",
						Image:   "busybox",
						Command: []string{"sleep", "3600"},
						SecurityContext: &corev1.SecurityContext{
							Privileged: func(b bool) *bool { return &b }(true),
						},
					},
				},
			},
		}

		successMessage := fmt.Sprintf("You can now exec to the debug pod by running:\nkubectl -n %s exec -it debug bash", ns)

		err = c.Create(context.Background(), pod)
		if errors.IsAlreadyExists(err) {
			fmt.Println(successMessage)
			return
		}
		if err != nil {
			panic(err)
		}

	},
}

func init() {
	ns = *debugCmd.Flags().String("namespace", "default", "The namespace to deploy the debug pod to")
}
