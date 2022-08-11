package internal

import (
	"fmt"
	"os"
	"testing"

	"github.com/hvxahv/hvx/rsa"

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
}

func TestAccount_IsExist(t *testing.T) {
	ok := NewUsername("hvxahv").IsExist()
	fmt.Println(ok)

	ok2 := NewUsername("hvturingga").IsExist()
	fmt.Println(ok2)
}

func TestAccount_Create(t *testing.T) {
	generate, err := rsa.NewRsa(2048).Generate()
	if err != nil {
		t.Error(err)
		return
	}
	if err := NewAccountsCreate(
		"hvxahv",
		"hvxahv@disism.com",
		"hvxahv123",
	).Create(generate.PublicKey); err != nil {
		t.Error(err)
		return
	}
}

func TestAccounts_Delete(t *testing.T) {
	if err := NewAccountsDelete("hvxahv", "hvxahv123").Delete(); err != nil {
		t.Error(err)
		return
	}
}

func TestAccounts_EditUsername(t *testing.T) {
	if err := NewAccountsID(1234567890).EditUsername("hvx"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccounts_EditPassword(t *testing.T) {
	if err := NewAccountsID(1234567890).EditPassword("hvx"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccounts_EditMail(t *testing.T) {
	if err := NewAccountsID(1234567890).EditEmail("hvx@disism.com"); err != nil {
		t.Error(err)
		return
	}
}

func TestAccounts_Verify(t *testing.T) {
	verify, err := NewVerify("hvturingga").Verify("hvxahv123")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(verify)
}
