/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package kubernetes

import (
	"github.com/spf13/cobra"
)

// KubernetesCmd represents the base command when called without any subcommands
var KubernetesCmd = &cobra.Command{
	Use:     "kubernetes",
	Short:   "Kubernetes related commands",
	Aliases: []string{"k8s"},
}

func init() {
	KubernetesCmd.AddCommand(debugCmd, sshNodeCmd)
}
