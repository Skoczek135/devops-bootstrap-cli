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

// sshNode uses kubectl node-shell to ssh into a node
var sshNodeCmd = &cobra.Command{
	Use:   "ssh",
	Short: "DevOps Bootstrap is a CLI tool to bootstrap multiple DevOps environments of every day usage",
	Run: func(cmd *cobra.Command, args []string) {
		command := "kubectl get nodes | tail +2 | fzf"
		subProcess := exec.Command("bash", "-c", command)
		var out bytes.Buffer
		subProcess.Stderr = os.Stderr
		subProcess.Stdout = &out
		if err := subProcess.Run(); err != nil {
			panic(err)
		}

		nodeOutput := out.String()
		nodeName := strings.Split(nodeOutput, " ")[0]

		fmt.Printf("Connecting to node %s\n", nodeName)

		command = fmt.Sprintf("kubectl node-shell %s", nodeName)
		fmt.Println(command)
		subProcess = exec.Command("bash", "-c", command)
		subProcess.Stdin = os.Stdin
		subProcess.Stdout = os.Stdout
		subProcess.Stderr = os.Stderr
		err := subProcess.Run()
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {}
