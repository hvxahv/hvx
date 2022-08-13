/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"fmt"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"
	"testing"

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
}

func TestArticles_Get(t *testing.T) {
	g, err := NewArticlesId(787516945347018753).Get(785518573776797697)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok", g)

	g2, err := NewArticlesId(787516945347018753).Get(785747557033967617)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("ok", g2)

	g3, err := NewArticlesId(787516945347018753).Get(787507052643319809)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("err", g3)
}
