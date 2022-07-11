package internal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

	// Initialize the database.
	if err := cockroach.NewDBAddr().InitDB(); err != nil {
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

func TestSaved_Create(t *testing.T) {
	d := v1alpha1.CreateSavedRequest{
		AccountId:   "733124680636596225",
		Name:        "YUI",
		Description: "description",
		Cid:         "QmVgBz2p2P3PnfiicJUHpyPVaiXDCNAKBnhimF9rP8c2zD",
		Types:       "jpeg/png",
		IsPrivate:   false,
	}

	s := saved{}
	create, err := s.CreateSaved(context.Background(), &d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(create)
}

func TestSaved_GetSaves(t *testing.T) {
	s := saved{}
	saves, err := s.GetSaves(context.Background(), &v1alpha1.GetSavesRequest{
		AccountId: "733124680636596225",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(saves)
}

func TestSaved_GetSaved(t *testing.T) {
	s := saved{}
	save, err := s.GetSaved(context.Background(), &v1alpha1.GetSavedRequest{
		AccountId: "733124680636596225",
		Id:        "738165894748110849",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(save)
}

func TestSaved_EditSaved(t *testing.T) {
	d := v1alpha1.EditSavedRequest{
		Id:          "738165894748110849",
		Name:        "",
		Description: "",
	}

	s := saved{}
	edit, err := s.EditSaved(context.Background(), &d)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(edit)
}
