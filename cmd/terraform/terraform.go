/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package terraform

import (
	"github.com/spf13/cobra"

	"devops-bootstrap/cmd/terraform/add"
	"devops-bootstrap/cmd/terraform/create"
)

// TerraformCmd represents the terraform command
var TerraformCmd = &cobra.Command{
	Use:     "terraform",
	Short:   "Subcommand for terraform scope",
	Aliases: []string{"tf"},
}

func init() {
	TerraformCmd.AddCommand(add.AddCmd, create.CreateCmd)
}
