package accounts

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cache"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	nd :=  db.NewDb()
	if err := nd.InitDB(); err != nil {
		return
	}

	// If a configs file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using configs file:", viper.ConfigFileUsed())
	}


	cache.InitRedis(1)

}

func TestNewAccounts(t *testing.T) {
	TestInitDB(t)
	na, _ := NewAccounts(
		"hvturingga",
		"hvxahv",
		"x@disism.com",
		)

	code, message := na.New()

	t.Log(code, message)
}


func TestAccounts_Update(t *testing.T) {
	TestInitDB(t)

	a := Accounts{
		Username:   "hvturingga",
		Password:   "123",
		Avatar:     "http://stage48.net/wiki/images/5/5b/KobayashiYui8th.jpg",
		Bio:        "我很开心，现在我在录制视频, 欢迎关注我的 YouTube 频道! AHHHHh.....",
		Name:       "HVTURINGGA",
		Mail:       "",
		Phone:      "",
		IsPrivate:    false,
		PrivateKey: "",
		PublicKey:  "",
	}

	if err := a.Update(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestAccounts_Query(t *testing.T) {
	TestInitDB(t)

	a := NewAcctByName("hvturingga")
	r, err := a.Find()
	if err != nil {
		return
	}
	t.Log(r)
}

func TestAccounts_Delete(t *testing.T) {
	TestInitDB(t)

	a := NewAcctByName("hvturingga")
	if err := a.Delete(); err != nil {
		t.Log(err)
		return
	}
	t.Log("Delete account successfully.")


}

func TestAccounts_Login(t *testing.T) {
	TestInitDB(t)

	a := NewAccountLogin("x@disism.com", "hvxahv")

	r, s, err := a.Login()
 	if err != nil {
		t.Error(err)
	}
	t.Log(r, s)

}