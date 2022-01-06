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
	"github.com/hvxahv/hvxahv/internal/hvx"
	"github.com/hvxahv/hvxahv/pkg/microservices/consul"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run http server of hvxahv.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p := viper.GetString("microservices.hvx.port")

		tags := []string{"hvx", "http", "RESTFul"}
		nr := consul.NewRegister("hvx", p, tags, "localhost")
		err := nr.Register()
		if err != nil {
			fmt.Println(err)
			return
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		api := hvx.APIServer()
		go func() {
			if err := api.Run(fmt.Sprintf(":%s", p)); err != nil {
				fmt.Println(err)
				return
			}
		}()

		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			err := consul.Deregister(nr.ID)
			if err != nil {
				log.Println(err)
			}
			return
		default:
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
