package saved

import (
	"context"
	"fmt"
	"github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
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
	if err2 := viper.ReadInConfig(); err2 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}
	// Initialize the database.
	if err := cockroach.NewDBAddr().InitDB(); err != nil {
		fmt.Println(err)
		return
	}
}

func TestSaved_Create(t *testing.T) {
	d := &v1alpha1.NewSavedCreate{
		AccountId:   "733124680636596225",
		Name:        "The Cathedral and the Bazaar",
		Description: "The Cathedral and the Bazaar: Musings on Linux and Open Source by an Accidental Revolutionary is an essay, and later a book, by Eric S. Raymond on software engineering methods, based on his observations of the Linux kernel development process and his experiences managing an open source project, fetchmail.",
		Hash:        "1-2-3-4-5",
		Types:       "application/pdf",
	}
	s := saved{}
	create, err := s.Create(context.Background(), d)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(create)
}

func TestSaved_GetSaves(t *testing.T) {
	d := &v1alpha1.NewSavedAccountID{
		AccountId: "733124680636596225",
	}
	s := saved{}
	saves, err := s.GetSaves(context.Background(), d)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(saves)
}

func TestSaved_GatSaved(t *testing.T) {
	d := &v1alpha1.NewSavedID{
		Id: "733164705208172545",
	}
	s := saved{}
	saved, err := s.GetSaved(context.Background(), d)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Println(saved)
}
