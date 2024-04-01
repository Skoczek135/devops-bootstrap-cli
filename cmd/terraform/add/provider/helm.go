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

//go:embed configs/helm.gotmpl
var helmTmpl string

// TerraformCmd represents the terraform command
var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Subcommand for terraform adding helm provider",
	Run: func(cmd *cobra.Command, args []string) {
		t := template.Must(template.New("helm").Parse(helmTmpl))
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
