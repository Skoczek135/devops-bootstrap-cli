/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	_ "embed"

	"github.com/spf13/cobra"
)

// CreateCmd is a subcomannd which does nothing but works as a placeholder
var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Responsible for creating modules and workspaces",
	Run: func(cmd *cobra.Command, args []string) {
		//
	},
}

func init() {
	CreateCmd.AddCommand(moduleCmd, workspaceCmd)
}
