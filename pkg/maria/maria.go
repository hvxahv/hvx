package maria

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"hvxahv/pkg/config"
	"log"
)

var (
	MariaDB *gorm.DB
)

type maria struct {
	username string
	password string
	hostname string
	port     string
	dbName   string
}

func newMaria() *maria {
	if err := config.InitConfig(); err != nil {
		log.Println(err)
	}
	return &maria{
		username: viper.GetString("maria.username"),
		password: viper.GetString("maria.password"),
		hostname: viper.GetString("maria.hostname"),
		port: viper.GetString("maria.prot"),
		dbName: viper.GetString("maria.dbname"),
	}
}

// InitDB ... Initialize the mongo
func InitMariaDB() error {
	m := newMaria()

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		m.username,
		m.password,
		m.hostname,
		m.port,
		m.dbName,
	)

	db, err := gorm.Open("mysql", addr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to MariaDB: %s", err))
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
