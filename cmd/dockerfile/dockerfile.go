/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package dockerfile

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// DockerfileLint represents the dockerfilelint command
var DockerfileLint = &cobra.Command{
	Use:     "dockerfile",
	Short:   "Subcommand for dockerfile linting",
	Aliases: []string{"docker", "df"},
	Run: func(cmd *cobra.Command, args []string) {
		dockerFileName := "Dockerfile"
		if len(args) > 0 {
			dockerFileName = args[0]
		}

		if _, err := os.Stat(dockerFileName); os.IsNotExist(err) {
			log.Printf("File %s does not exist", dockerFileName)
			os.Exit(1)
		}

		c := exec.Command("bash", "-c", fmt.Sprintf("docker run --rm -i hadolint/hadolint < %s", dockerFileName))
		c.Stdout = os.Stdout

		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	},
}
