/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"context"
	"fmt"
	"os"
	"os/exec"

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
						Image:   "ubuntu",
						Command: []string{"sleep", "3600"},
						SecurityContext: &corev1.SecurityContext{
							Privileged: func(b bool) *bool { return &b }(true),
						},
					},
				},
			},
		}

		err = c.Create(context.Background(), pod)
		if err != nil && !errors.IsAlreadyExists(err) {
			return
		}

		wait := make(chan bool)
		go func() {
			pod := &corev1.Pod{}
			for {
				c.Get(context.Background(), client.ObjectKey{
					Name:      "debug",
					Namespace: ns,
				}, pod)
				if pod.Status.Phase == corev1.PodRunning {
					wait <- true
				}
			}
		}()

		<-wait

		command := fmt.Sprintf("kubectl -n %s exec -it %s bash", ns, "debug")
		subProcess := exec.Command("bash", "-c", command)
		subProcess.Stdin = os.Stdin
		subProcess.Stdout = os.Stdout
		subProcess.Stderr = os.Stderr
		_ = subProcess.Run()

		defer func() {
			err = c.Delete(context.Background(), pod)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Debug pod deleted successfully")
		}()
	},
}

func init() {
	ns = *debugCmd.Flags().String("namespace", "default", "The namespace to deploy the debug pod to")
}
