/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"github.com/spf13/cobra"

	"devops-bootstrap/cmd/terraform/add/provider"
)

var commands = []*cobra.Command{provider.ProviderCmd, helmCmd}

// TerraformCmd represents the terraform command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Subcommand for terraform adding providers / helm releases",
}

func init() {
	AddCmd.AddCommand(commands...)
}
