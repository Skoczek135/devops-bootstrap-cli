/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"

	"devops-bootstrap/cmd/dockerfile"
)

// Lint represents the linting command
var Lint = &cobra.Command{
	Use:   "lint",
	Short: "Subcommand for linting",
}

func init() {
	Lint.AddCommand(dockerfile.DockerfileLint)
}
