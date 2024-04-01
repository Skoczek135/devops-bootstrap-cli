/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package add

// import (
// 	"github.com/spf13/cobra"
// )
//
// var moduleCmd = &cobra.Command{
// 	Use:   "module",
// 	Short: "Subcommand for adding module to the workspace",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		url := args[0]
// 		version := args[1] != "" ? args[1] : "latest"
//
// 		// pull the module
// 		module, err := pullModule(url, version)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Printf("/+v\n", module)
//
// 		// add the module to the workspace
// 		err = addModule(module)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 	},
// }
// func pullModule(url string, version string) (*module.Module, error) {
// 	// pull the module
// 	module, err := action.NewPullModule(url, version)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return module, nil
// }
//
// func init() {}
