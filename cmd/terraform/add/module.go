/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	//go:embed resources/terraform_init.gotmpl
	initTemplate string
)

type tfvar struct {
	Name         string
	Description  string
	DefaultValue any
}

type Module struct {
	Url       string `json:"url,omitempty"`
	Version   string `json:"version,omitempty"`
	Variables []tfvar
}

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Subcommand for adding module to the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		m := &Module{}
		m.Url = args[0]
		if len(args) > 1 {
			m.Version = args[1]
		} else {
			m.Version = ""
		}

		// Add init module to the workspace
		file, err := os.OpenFile(fmt.Sprintf("%s%s", "module", ".tf"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		t := template.Must(template.New("module").Parse(initTemplate))
		err = t.Execute(file, m)
		if err != nil {
			fmt.Println(err)
		}

		// Run get command
		command := "terraform get"
		subProcess := exec.Command("bash", "-c", command)
		subProcess.Stdin = os.Stdin
		subProcess.Stdout = os.Stdout
		subProcess.Stderr = os.Stderr
		_ = subProcess.Run()

		// Iterate through the modules and find all variables
		dirEntry, err := os.ReadDir(".terraform/modules/template/")
		if err != nil {
			fmt.Println(err)
			return
		}

		variables := []tfvar{}

		for _, entry := range dirEntry {
			if entry.IsDir() {
				continue
			}

			file, err := os.Open(fmt.Sprintf(".terraform/modules/template/%s", entry.Name()))
			if err != nil {
				fmt.Println(err)
				return
			}
			var isVariable bool
			var variableName string
			var variableDescription string
			var variableDefault any

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "variable \"") {
					isVariable = true
					variableName = strings.Split(line, "\"")[1]
					continue
				}

				if isVariable && strings.Contains(line, "description") {
					variableDescription = strings.Split(line, "\"")[1]
					continue
				}

				if isVariable && strings.Contains(line, "default") {
					variableDefault = strings.Split(line, "=")[1]
				}

				if isVariable && strings.Contains(line, "type") {
					if variableDefault == "" {
						switch strings.Split(line, "=")[1] {
						case "string":
							variableDefault = "\"\""
						case "number":
							variableDefault = -1
						case "bool":
							variableDefault = false
						}
					}
				}

				if isVariable && strings.Contains(line, "}") {
					isVariable = false
					variables = append(variables, tfvar{Name: variableName, Description: variableDescription, DefaultValue: variableDefault})
					variableName = ""
					variableDescription = ""
					variableDefault = ""
				}
			}
		}

		m.Variables = variables

		file, err = os.OpenFile(fmt.Sprintf("%s%s", "module", ".tf"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		t = template.Must(template.New("module").Parse(initTemplate))
		err = t.Execute(file, m)
		if err != nil {
			fmt.Println(err)
		}
	},
}
