package cockroach

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

	// New a database, receive the name and return an error.
	New(name string) error
}

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// DBAddr ...
type DBAddr struct {
	Addr string
}

func NewDBAddr() *DBAddr {
	var (
		host     = viper.GetString("cockroach.host")
		port     = viper.GetString("cockroach.port")
		user     = viper.GetString("cockroach.user")
		password = viper.GetString("cockroach.password")
		dbName   = viper.GetString("cockroach.dbName")
		sslMode  = viper.GetString("cockroach.sslMode")
		timeZone = viper.GetString("cockroach.timeZone")
	)

	addr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbName,
		port,
		sslMode,
		timeZone,
	)
	return &DBAddr{Addr: addr}
}

func (d *DBAddr) InitDB() error {
	var initError error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return errors.Errorf("init db error: %v", initError)
		default:
			dbs, err := gorm.Open(postgres.Open(d.Addr), &gorm.Config{})
			if err != nil {
				initError = err
			} else {
				log.Println("initialize the database successfully.")
				db = dbs
				return nil
			}
		}
	}
}

func (d *DBAddr) New(name string) error {

	c, err := sql.Open("postgres", d.Addr)
	if err != nil {
		fmt.Println(err)
	}

	if _, err2 := c.Exec("CREATE DATABASE " + name); err2 != nil {
		panic(err2)
		return err2
	}
	return nil
}

