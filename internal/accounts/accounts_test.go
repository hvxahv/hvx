package accounts

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {

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
	n := cockroach.NewDBAddr()
	if err2 := n.InitDB(); err2 != nil {
		return
	}

	// If a configs file is found, read it in.
	if err3 := viper.ReadInConfig(); err3 == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}

	cache.InitRedis(1)

}

func TestNewAccounts(t *testing.T) {
	TestInitDB(t)

	a := NewAccounts("hvturingga", "x@disism.com", "Hvxahv123")
	err := a.Create()
	if err != nil {
		log.Println(err)
	}
}

func TestAccounts_FindAccountByName(t *testing.T) {
	TestInitDB(t)

	a := NewAccountsName("hvturingga")
	accounts, err := a.FindAccountByUsername()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(accounts)
}

func TestAccounts_Update(t *testing.T) {
	TestInitDB(t)

	a := NewAccountsName("hvturingga")
	a.Password = "Hvxahv123"
	a.Mail = "x@disism.com"


	err := a.Update()
	if err != nil {
		log.Println(err)
		return
	}

}

func TestAccounts_ChangeUsername(t *testing.T) {
	TestInitDB(t)

	a := NewAcctNameANDActorID("hvturinggas", 696077920006668289)
	err := a.UpdateUsername("hvturingga")
	if err != nil {
		log.Println(err)
		return
	}
}

func TestAccounts_Delete(t *testing.T) {
	TestInitDB(t)

	a := NewAcctNameANDActorID("hvturingga", 696077920006668289)
	err := a.Delete()
	if err != nil {
		log.Println(err)
		return
	}
}
