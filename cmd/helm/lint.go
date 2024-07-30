/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package helm

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// helmLint is used for linting helm yaml
var HelmLint = &cobra.Command{
	Use:   "helm",
	Short: "Subcommand for helm package linting",
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
