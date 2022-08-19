package fs

import (
	"fmt"
	"os"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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
}

func TestGetMinio(t *testing.T) {
	if err := NewMinio().Dial(); err != nil {
		t.Error(err)
		return
	}

	client := GetMinio()

	ctx := context.Background()

	exists, err := client.BucketExists(ctx, "avatar")
	if err != nil {
		t.Error(err)
		return
	}
	if exists {
		t.Log("Bucket found")
	}
}
