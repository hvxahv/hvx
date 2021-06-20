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
	cfgFile := fmt.Sprintf("%s/configs/configs.yaml", file)
	viper.SetConfigFile(cfgFile)
	fmt.Println(file)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error configs file: %s \n", err))
	}


	nd := db.NewDb()
	if err := nd.InitDB(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestNewAccounts(t *testing.T) {
	TestInitDB(t)

	na := NewAccounts(
		"hvturingga",
		"hvxahv",
		"https://cdn.keyakizaka46.com/images/14/103/4f2a17f7f544a1635c244502dc8ea/400_320_102400.jpg",
		"HVTURINGGA",
		"x@disism.com",
		0,
		)

	if err := na.New(); err != nil {
		return 
	}
}

func TestAccounts_Update(t *testing.T) {
	TestInitDB(t)

	a := NewAccountByName("hvturingga")

	a.Bio = "我很开心，现在我在录制视频, 欢迎关注我的频道!"
	if err := a.Update(); err != nil {
		t.Errorf("%v",err)
	}
}

func TestAccounts_Query(t *testing.T) {
	TestInitDB(t)

	a := NewAccountByName("hvturingga")
	r, err := a.Query()
	if err != nil {
		return
	}
	t.Log(r)
}

func TestAccounts_Delete(t *testing.T) {

}

func TestAccounts_Login(t *testing.T) {
	TestInitDB(t)
	//a := NewAccountLogin("hvturingga", "hvxahv")
	a := NewAccountLogin("hvturingga", "hvxahv")
	r, err := a.Login()
	if err != nil {
		t.Error(err)
	}
	t.Log(r)


}