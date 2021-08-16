package channel

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

func TestNewChannels(t *testing.T) {
	nc := NewChannels("ALICE HOUSE", "", "avatar", "bio", "alice", true)
	fmt.Println(nc.Id)


	nc1 := NewChannels("Hvxahv Chan", "", "avatar", "bio", "hvturingga",false)
	fmt.Println(nc1)
}

func TestNewChannels2(t *testing.T) {
	TestInitDB(t)
	nc := NewChannels("ALICE HOUSE", "", "avatar", "bio", "alice", true)
	i, s, _, err := nc.New()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("code: %v; message: %s", i, s)

	nc2 := NewChannels("JSUT 4 FUN", "", "avatar", "bio", "hvturingga", false)
	i2, s2, _, err := nc2.New()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("code: %v; message: %s", i2, s2)

}


func TestNewFindChannelByID(t *testing.T) {
	TestInitDB(t)
	nfc := NewChannelsByID("f6574uSSqGQ7CJX")
	ch := nfc.Find()
	fmt.Println(ch.Name)
}


func TestMChannels_AddAdmin(t *testing.T) {
	TestInitDB(t)
	nmc := NewChanAdmins("Ja5QZv-fgxhg182", "bob")
	admin, s, err := nmc.AddAdmin()
	if err != nil {
		return
	}
	fmt.Printf("code: %v; message: %s", admin, s)

	nmc2 := NewChanAdmins("f6574uSSqGQ7CJX", "hvturingga")
	admin2, s2, err := nmc2.AddAdmin()
	if err != nil {
		return
	}
	fmt.Printf("code: %v; message: %s", admin2, s2)

}

func TestChanAdmins_GetChannelListByName(t *testing.T) {
	TestInitDB(t)
	name := NewChanAdminsByName("hvturingga")
	_, cls, _ := name.GetChannelListByName()
	fmt.Println(cls)

	for _, i := range cls {
		fmt.Println(i.Name)
	}

}

func TestChanAdmins_GetChannelListByID(t *testing.T) {
	TestInitDB(t)
	adm := NewChanAdminsByID("a-t0FuZY9ySeBlR")
	adm.GetChanAdmLisByID()

}

func TestNewChanSub(t *testing.T) {
	TestInitDB(t)
	ncs := NewChanSub("a-t0FuZY9ySeBlR", "bob")
	subscriber, s, err := ncs.NewSubscriber()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(subscriber, s)
}

func TestChanSubs_GetSubscriberByID(t *testing.T) {
	TestInitDB(t)
	ncs := NewChanSubByID("a-t0FuZY9ySeBlR")
	ncs.GetSubscriberByID()
}