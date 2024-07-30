/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// KubernetesLint is used for linting kubernetes yaml
var KubernetesLint = &cobra.Command{
	Use:     "kubernetes",
	Aliases: []string{"k", "k8s"},
	Short:   "Subcommand for kubernetes yaml linting",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("No yaml file given")
		}
		yamlFile := args[0]

		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			log.Fatal("File doesn't exist")
		}

		command := fmt.Sprintf("kube-linter lint %s", yamlFile)

		c := exec.Command("bash", "-c", command)
		c.Stderr = os.Stderr
		c.Stdout = os.Stdout

		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
}
