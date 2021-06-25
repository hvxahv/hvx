package accounts

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"hvxahv/pkg/db"
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

}

func TestNewAccounts(t *testing.T) {
	TestInitDB(t)

	na := NewAccounts(
		"hvturingga",
		"hvxahv",
		"https://cdn.keyakizaka46.com/images/14/103/4f2a17f7f544a1635c244502dc8ea/400_320_102400.jpg",
		"HVTURINGGA" ,
		"x@disism.com",
		0,
		)

	if err := na.New(); err != nil {
		t.Error(err)
	}
}

func TestAccounts_Update(t *testing.T) {
	TestInitDB(t)

	a := NewUpdateAcct()
	a.Username = "hvturingga"
	a.Bio = "我很开心，现在我在录制视频, 欢迎关注我的频道!"
	if err := a.Update(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestAccounts_Query(t *testing.T) {
	TestInitDB(t)

	a := NewQueryAcctByName("hvturingga")
	r, err := a.Query()
	if err != nil {
		return
	}
	t.Log(r)
}

func TestAccounts_Delete(t *testing.T) {
	TestInitDB(t)

	a := NewDelAcctByName("hvturingga")
	if err := a.Delete(); err != nil {
		t.Log(err)
		return
	}
	t.Log("Delete account successfully.")


}

func TestAccounts_Login(t *testing.T) {
	TestInitDB(t)

	a := NewAccountLogin("hvturingga", "hvxahv")

	r, err := a.Login()
 	if err != nil {
		t.Error(err)
	}
	t.Log(r)

}