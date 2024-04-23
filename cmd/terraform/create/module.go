/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	//go:embed configs/requires.tf
	requires string

	modulesTemplatesMap = map[string]string{
		"variables.tf": "",
		"outputs.tf":   "",
		"requires.tf":  requires,
	}
	// moduleCmd represents the module to create modules components
	moduleCmd = &cobra.Command{
		Use:   "module",
		Short: "Initializes a terraform module",
		Long: "Creates a terraform module with the following structure:\n" +
			"variables.tf\n" +
			"outputs.tf\n" +
			"requires.tf\n",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			err := createDir(name)
			if err != nil {
				fmt.Println(err)
			}
			for fileName, template := range templatesMap {
				file, err := createFile(fileName)
				if err != nil {
					fmt.Println(err)
				}
				defer file.Close()

				if template != "" {
					err = fillFileWithTemplate(file, template)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
			fmt.Println("Terraform project initialized")
		},
	}
)

func init() {}

func createDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return err
	}
	return nil
}
