package accounts

import (
	"fmt"
	"github.com/spf13/viper"
	"hvxahv/pkg/db"
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	file, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i <= 1; i ++ {
		file = filepath.Dir(file)
	}
	cfgFile := fmt.Sprintf("%s/configs/config.yaml", file)
	viper.SetConfigFile(cfgFile)
	fmt.Println(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbName := viper.GetString("db.dbName")
	sslMode := viper.GetString("db.sslMode")

	nd := db.NewDb(host, port, user, password, dbName, sslMode)
	if err := nd.InitPostgre(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestNewAccounts(t *testing.T) {
	TestInitDB(t)

	na := NewAccounts(
		"hvturingga",
		"hvxahv123",
		"https://cdn.keyakizaka46.com/images/14/103/4f2a17f7f544a1635c244502dc8ea/400_320_102400.jpg",
		"I Love Programming",
		"HVTURINGGA",
		"x@disism.com",
		7064263007,
		"hvturingga",
		"",
		1,
		0,
		2,
		6,
		)

	if err := na.New(); err != nil {
		return 
	}
}

func TestAccounts_Update(t *testing.T) {
	TestInitDB(t)

	a := NewAccountByName("hvturingga")
	a.Bio  = "我很开心，现在我在录制视频, 欢迎关注我的频道!"
	if err := a.Update(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestAccounts_Query(t *testing.T) {
	TestInitDB(t)

	a := NewAccountByName("hvturinggas")
	r, err := a.Query()
	if err != nil {
		return
	}
	t.Log(r)
}