/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package cmd

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/fs"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "db":
			fmt.Println("create database...")

			if err := cockroach.NewRoach().Create(args[1]); err != nil {
				fmt.Println(err)
				return
			}
		case "fs":
			buckets := []string{"avatar"}
			for _, i := range buckets {
				fmt.Println(fmt.Printf("create %s bucket...", i))
				if err := fs.NewBucket(args[1], i).Create(); err != nil {
					fmt.Println(err)
					return
				}
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
