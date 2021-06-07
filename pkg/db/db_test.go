package db

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"
)

func TestInitDB(t *testing.T) {
	viper.SetConfigFile("./db.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbName := viper.GetString("db.dbName")
	sslMode := viper.GetString("db.sslMode")

	nd := NewDb(host, port, user, password, dbName, sslMode)
	if err := nd.InitPostgre(); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}

func TestCreateDB(t *testing.T) {
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

	nd := NewDb(host, port, user, password, dbName, sslMode)

	drive := "postgres"
	name := "hvxahv"

	if err := nd.Create(drive, name); err != nil {
		t.Errorf("Failed to initialize PostgreSQL : %s", err)
	} else {
		t.Logf("Initialize PostgreSQL success.")
	}
}
