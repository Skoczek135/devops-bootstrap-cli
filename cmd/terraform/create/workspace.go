/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	_ "embed"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	//go:embed configs/providers.tf
	providers string

	//go:embed configs/variables.tf
	variables string

	//go:embed configs/backend.tf
	backend string

	templatesMap = map[string]string{
		"providers.tf": providers,
		"variables.tf": variables,
		"backend.tf":   backend,
		"main.tf":      "",
		"outputs.tf":   "",
	}
)

// workspaceCmd represents the command for the creating workspace layout
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Initializes a terraform project",
	Run: func(cmd *cobra.Command, args []string) {
		for fileName, template := range templatesMap {
			file, err := createFile(fileName)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			err = fillFileWithTemplate(file, template)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("Terraform project initialized")
	},
}

func init() {}
