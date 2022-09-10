/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package activitypub

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func TestGetActorName(t *testing.T) {
	resource := "acct:hvturingga@halfmemories.com"
	actor := GetActorName(resource)
	t.Log(actor)

}

func TestGetActorByWebfinger(t *testing.T) {
	handler, err := GetWebFingerHandler("hvturingga@mastodon.disism.com")
	if err != nil {
		t.Error(err)
		return
	}
	webfinger, err := GetActorByWebfinger(handler)
	if err != nil {
		t.Error()
		return
	}
	t.Log(webfinger)
}

//
//func TestGetActorHost(t *testing.T) {
//	resource := "hvturingga@halfmemories.com"
//	host := GetActorHost(resource)
//	fmt.Println(host)
//}

//func TestIsRemote(t *testing.T) {
//	resource := "acct:hvturingga@halfmemories.com"
//	resource2 := "acct:hvturingga@disism.com"
//	i := IsRemote(resource)
//	fmt.Println(i)
//	i2 := IsRemote(resource2)
//	fmt.Println(i2)
//}
