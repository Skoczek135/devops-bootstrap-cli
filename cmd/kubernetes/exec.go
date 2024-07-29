/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// execCmd exec -it into a pod
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execs into a pod based on the fizzed output of kubectl get pods",
	Run: func(cmd *cobra.Command, args []string) {
		command := "kubectl get pods -A | tail +2 | fzf"
		subProcess := exec.Command("bash", "-c", command)
		var out bytes.Buffer
		subProcess.Stderr = os.Stderr
		subProcess.Stdout = &out
		if err := subProcess.Run(); err != nil {
			panic(err)
		}

		podOutput := out.String()
		noSpacedOutput := strings.Fields(podOutput)
		namespace := noSpacedOutput[0]
		podName := noSpacedOutput[1]

		command = fmt.Sprintf("kubectl -n %s exec -it %s bash", namespace, podName)
		subProcess = exec.Command("bash", "-c", command)
		subProcess.Stdin = os.Stdin
		subProcess.Stdout = os.Stdout
		subProcess.Stderr = os.Stderr
		_ = subProcess.Run()
	},
}

func init() {}
