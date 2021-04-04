package maria

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var (
	MariaDB *gorm.DB
)

// InitDB ... Initialize the mongo
func InitMariaDB() error {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	u := viper.Get("maria.username")
	pwd := viper.Get("maria.password")
	h := viper.Get("maria.hostname")
	port := viper.Get("maria.prot")
	dbNme := viper.Get("maria.dbname")

	l := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		u, pwd, h, port, dbNme,
	)

	db, err := gorm.Open("mysql", l)
	if err != nil {
		fmt.Println("连接数据库失败！")
	}

	db.SingularTable(true)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(20)
	MariaDB = db

	return err
}

func GetMaria() *gorm.DB {
	return MariaDB
}