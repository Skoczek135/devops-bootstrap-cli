/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"devops-bootstrap/cmd/aws"
	"devops-bootstrap/cmd/kubernetes"
	"devops-bootstrap/cmd/terraform"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "devops-bootstrap",
	Short: "DevOps Bootstrap is a CLI tool to bootstrap multiple DevOps environments of every day usage",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(terraform.TerraformCmd, kubernetes.KubernetesCmd, Lint, aws.CloudTrail)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
