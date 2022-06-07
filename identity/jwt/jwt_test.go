package jwt

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Search configs in home directory with name ".hvxahv" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".hvxahv")

	viper.AutomaticEnv()

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	// Initialize the database.
	if err := cockroach.NewRoach().Dial(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}
}

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyZGF0YSI6eyJtYWlsIjoieEBkaXNpc20uY29tIiwiaWQiOiIxMjMxMjMiLCJ1c2VybmFtZSI6Imh2dHVyaW5nZ2EiLCJkZXZpY2VfaWQiOiIxLTItMy00LTUifSwiand0Xy5fc3RhbmRhcmRfY2xhaW1zIjp7ImlzcyI6IjEtMi0zLTQtNS1odngtaXNzdWVzIiwic3ViIjoiaHZ0dXJpbmdnYTp0b2tlbiIsImV4cCI6MTY2MDIxNzU4MSwibmJmIjoxNjU1MDMzNTgxLCJpYXQiOjE2NTUwMzM1ODEsImp0aSI6IjRhOTViYTFmLTJiMGEtNDVkZC1iN2FmLWFmNGFhMTkzNjE2NiJ9fQ.KC3FVGoA2MTm0C2ZNSpswij3GB_JAcsoHCkY_DxvxWE"
)
