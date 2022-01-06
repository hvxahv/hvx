package account

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hvxahv/hvxahv/pkg/cache"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
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
	n := cockroach.NewDBAddr()
	if err := n.InitDB(); err != nil {
		fmt.Println(err)
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println(err)
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
		return
	}

	cache.InitRedis(1)
}

func TestAccounts_Create(t *testing.T) {
	actor, err := NewActors("", "", "").Create()
	if err != nil {
		log.Println(err)
		return
	}
	if err := NewAccounts("hvturingga", "x@disism.com", "Hvxahv123", actor.ID).Create(); err != nil {
		log.Println(err)
		return
	}
}

func TestAccounts_FindAccountByName(t *testing.T) {
	a := NewAccountsUsername("hvturingga")
	accounts, err := a.GetAccountByUsername()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(accounts)
}

func TestAccounts_Update(t *testing.T) {
	a := NewAccountsUsername("hvturingga")
	a.Password = "Hvxahv123"
	a.Mail = "x@disism.com"

	err := a.Update()
	if err != nil {
		log.Println(err)
		return
	}

}

func TestAccounts_ChangeUsername(t *testing.T) {
	a := NewAcctNameANDActorID("hvturinggas", 696077920006668289)
	err := a.UpdateUsername("hvturingga")
	if err != nil {
		log.Println(err)
		return
	}
}

func TestAccounts_Delete(t *testing.T) {
	a := NewAcctNameANDActorID("hvturingga", 696077920006668289)
	err := a.Delete()
	if err != nil {
		log.Println(err)
		return
	}
}
