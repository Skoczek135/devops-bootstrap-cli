/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package provider

import (
	"github.com/spf13/cobra"
)

// TerraformCmd represents the terraform command
var ProviderCmd = &cobra.Command{
	Use:   "provider",
	Short: "Subcommand for terraform adding provider",
}

func init() {
	ProviderCmd.AddCommand(kubernetesCmd, helmCmd, awsCmd)
}
