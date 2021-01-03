package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

// InitDB ... Initialize the database
func NewDB() (*gorm.DB, error) {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	u := viper.Get("mariadb.username")
	pwd := viper.Get("mariadb.password")
	h := viper.Get("mariadb.hostname")
	port := viper.Get("mariadb.prot")
	dbNme := viper.Get("mariadb.dbname")

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

	return db, nil
}