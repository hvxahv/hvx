package minio

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
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

func TestInitMinIO(t *testing.T) {
	minio, err := InitMinIO()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(minio)
}

func TestMinIO_CreateMinIOBucket(t *testing.T) {
	nb := NewBucket("avatar", "ap-northeast-3")
	if nb == nil {
		fmt.Printf("failed to connect to minio")
	}
	if err := nb.Create(); err != nil {
		log.Println(err)
		return
	}
}
