package internal

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvx/cockroach"
	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

func TestAccount_IsExist(t *testing.T) {
	ok := NewUsername("hvturingga").IsExist()
	fmt.Println(ok)

	ok2 := NewUsername("hvx1").IsExist()
	fmt.Println(ok2)
}

func TestAccount_Create(t *testing.T) {
	if err := NewAccountsCreate("hvx1", "hvx1@disism.com", "hvxahv123").Create("publicKey"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccount_GetAccountByUsername(t *testing.T) {
	a, err := NewUsername("hvx1").GetAccountByUsername()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestAccounts_GetActorByUsername(t *testing.T) {
	actor, err := NewUsername("hvturingga").GetActorByUsername()
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(actor)
}

func TestAccount_Delete(t *testing.T) {
	if err := NewAccountsDelete("hvx1", "hvxahv123").Delete(); err != nil {
		t.Error(err)
		return
	}
}

func TestAccount_EditUsername(t *testing.T) {
	if err := NewAccountsID(12345).EditUsername("hvx2"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccount_EditEmail(t *testing.T) {
	if err := NewAccountsID(12345).EditEmail("hvx2@disism.com"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccount_EditPassword(t *testing.T) {
	if err := NewEditPassword("hvx2", "hvxahv123").EditPassword("hvx123456"); err != nil {
		t.Error(err)
		return
	}
}
