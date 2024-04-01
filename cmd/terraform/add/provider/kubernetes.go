/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package provider

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed configs/kubernetes.gotmpl
var kubernetesTmpl string

// TerraformCmd represents the terraform command
var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Subcommand for terraform adding kubernetes provider",
	Run: func(cmd *cobra.Command, args []string) {
		t := template.Must(template.New("kubernetes").Parse(kubernetesTmpl))
		file, err := os.OpenFile("providers.tf", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("First run `devops-cli terraform init`")
		}
		err = t.Execute(file, nil)
		if err != nil {
			fmt.Printf("%+v", file)
			fmt.Println(err)
		}

	},
}

func init() {}
