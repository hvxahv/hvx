/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/hvxahv/hvxahv/app/hvx/pkg"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:        "new",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "A new method of hvx.",
	Long:       `You can use this method to create databases and microservices.`,
	Example:    "hvx new database <name> or hvx new svc <name>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("You must enter 2 parameters. For example, hvx new <parameter> <name> .")
			return
		}

		name := args[1]

		switch args[0] {
		case "database":
			err := pkg.NewDB(name)
			if err != nil {
				fmt.Println(err)
			}
		case "svc":
			err := pkg.NewServices(name)
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("execute command error, please use help to execute the command correctly.")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
