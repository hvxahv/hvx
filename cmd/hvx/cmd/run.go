/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package cmd

import (
	"fmt"
	"github.com/hvxahv/hvxahv/api/proto/hvx"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"github.com/hvxahv/hvxahv/pkg/microservices/consul"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p := microservices.NewService("hvx").GetPort()

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
