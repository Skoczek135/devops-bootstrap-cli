/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package shell

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// ShellLint represents the shelllint command
var ShellLint = &cobra.Command{
	Use:     "shell",
	Short:   "Subcommand for shell linting",
	Aliases: []string{"sh"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 0 {
			log.Fatal("Argument has not been passed!")
		}
		shellScript := args[0]

		if _, err := os.Stat(shellScript); os.IsNotExist(err) {
			log.Printf("File %s does not exist", shellScript)
			os.Exit(1)
		}

		c := exec.Command("bash", "-c", fmt.Sprintf("shellcheck %s", shellScript))
		c.Stdout = os.Stdout

		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
