/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/hvxahv/hvx/internal/channel"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

const serviceName = "channel"

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := x.NewService(serviceName).GetPort()

		tags := []string{serviceName, "gRPC"}
		nr := consul.NewRegister(serviceName, port, tags, "localhost")
		if err := nr.Register(); err != nil {
			fmt.Println(err)
			return
		}

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

		if err := channel.Run(); err != nil {
			fmt.Printf("failed to start %s gRPC service: %v", serviceName, err)
			return
		}
		fmt.Println("Starting gRPC server...")
		fmt.Printf("%s gRPC server listening on port:%v", serviceName, port)

		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			if err := consul.Deregister(nr.ID); err != nil {
				fmt.Println(err)
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
