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
	"github.com/spf13/viper"
	"hvxahv/pkg/db"

	"github.com/spf13/cobra"
)

// cdbCmd represents the cdb command
var cdbCmd = &cobra.Command{
	Use:   "cdb",
	Short: "Fast Create PostgreSQL or Mysql DB.",
	Long: `hvx cdb postgres hvxahv`,
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("db.host")
		port := viper.GetString("db.port")
		user := viper.GetString("db.user")
		password := viper.GetString("db.password")
		dbName := viper.GetString("db.dbName")
		sslMode := viper.GetString("db.sslMode")


		nd :=  db.NewDb(host, port, user, password, dbName, sslMode)
		if err := nd.Create(args[0]); err != nil {
			fmt.Printf("Failed to initialize PostgreSQL : %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cdbCmd)
	rootCmd.Flags().String("database name", "", "Enter your database name.")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
