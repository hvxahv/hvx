package db

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type DB interface {
	// InitDB Initialize the PostgreSQL database.
	InitDB() error
	// Create database
	Create(name string) error
}

type db struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}

type Config struct {
	Conn *viper.Viper
}

func NewDb() *db {

	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	dbName := viper.GetString("db.dbName")
	sslMode := viper.GetString("db.sslMode")

	return &db{host: host, port: port, user: user, password: password, dbName: dbName, sslMode: sslMode}
}

var sdb *gorm.DB

func GetDB() *gorm.DB {
	return sdb
}

func (d *db) InitDB() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		d.host,
		d.user,
		d.password,
		d.dbName,
		d.port,
		d.sslMode,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var initError error

	for {
		select {
		case <-ctx.Done():
			return errors.Errorf("Init DB Error: %v", initError)
		default:
			dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				initError = err
			} else {
				log.Println("Initialize the database successfully.")
				sdb = dbs
				return nil
			}
		}
	}
}

func (d *db) Create(name string) error {
	addr := fmt.Sprintf("port=%s user=%s password=%s host=%s sslmode=%s",
		d.port,
		d.user,
		d.password,
		d.host,
		d.sslMode)

	cdb, err := sql.Open("postgres", addr)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := cdb.Exec("CREATE DATABASE " + name); err != nil {
		panic(err)
		return err
	}
	return nil
}
