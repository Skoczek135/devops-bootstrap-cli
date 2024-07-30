/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"devops-bootstrap/cmd/dockerfile"
	"devops-bootstrap/cmd/kubernetes"
	"devops-bootstrap/cmd/shell"
)

// Lint represents the linting command
var Lint = &cobra.Command{
	Use:   "lint",
	Short: "Subcommand for linting",
}

func init() {
	Lint.AddCommand(dockerfile.DockerfileLint, shell.ShellLint, kubernetes.KubernetesLint)
}
