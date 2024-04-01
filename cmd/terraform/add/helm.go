/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/registry"
)

type ChartConfig struct {
	ChartName    string
	ChartVersion string
	ChartUrl     string
}

var (
	//go:embed resources/helm_release.gotmpl
	helmReleaseTemplate string

	tmpDir, _ = os.MkdirTemp(".", "helm-release")

	helmCmd = &cobra.Command{
		Use:   "helm",
		Short: "Subcommand for adding helm releases",
		Run: func(cmd *cobra.Command, args []string) {
			cc := ChartConfig{
				ChartUrl:     args[0],
				ChartName:    args[1],
				ChartVersion: "",
			}
			err := cc.pullChart("", "")
			if err != nil {
				fmt.Println(err)
			}

			err = copyChartValues(fmt.Sprintf("%s/%s", tmpDir, cc.ChartName), ".")
			if err != nil {
				fmt.Println(err)
			}

			// render the helm release template
			file, err := os.OpenFile(fmt.Sprintf("%s%s", cc.ChartName, ".tf"), os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
			}
			t := template.Must(template.New(cc.ChartName).Parse(helmReleaseTemplate))
			err = t.Execute(file, cc)
			if err != nil {
				fmt.Println(err)
			}

			cleanup()
		},
	}
)

func init() {}

func logDebug(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

func (c *ChartConfig) pullChart(username string, password string) error {
	// registry client
	registryClient, err := registry.NewClient(
		registry.ClientOptDebug(false),
		registry.ClientOptEnableCache(false),
	)
	if err != nil {
		return err
	}

	// init helm action config
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(nil, "", "secret", logDebug); err != nil {
		return err
	}

	actionConfig.RegistryClient = registryClient

	// pull the chart
	pull := action.NewPullWithOpts(action.WithConfig(actionConfig))
	pull.Settings = cli.New() // didn't want to do this but otherwise it goes nil pointer
	pull.ChartPathOptions.Version = c.ChartVersion
	pull.RepoURL = c.ChartUrl
	pull.Untar = true
	pull.UntarDir = tmpDir

	_, err = pull.Run(c.ChartName)
	if err != nil {
		return err
	}
	return nil
}

// copyChartValues copies the values from the source chart to the destination chart
func copyChartValues(sourceDir string, destDir string) error {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/values.yaml", sourceDir))
	if err != nil {
		return err
	}
	// Write data to dst
	err = ioutil.WriteFile(fmt.Sprintf("%s/values.yaml", destDir), data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// cleanUp removes the temporary directory
func cleanup() {
	err := os.RemoveAll(tmpDir)
	if err != nil {
		fmt.Println(err)
	}
}
