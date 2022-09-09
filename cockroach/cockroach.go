package cockroach

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/hvxahv/hvx/errors"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Cockroach is a wrapper around a cockroach database connection.
var (
	db *gorm.DB
)

type DB interface {
	// Dial Call the configuration file to initialize and connect to the database,
	// If the connection is incorrect, a custom error is returned.
	// If the connection fails, there will be a 10second retry connection time.
	Dial() error

	// Create a database method for hvxctl to create a DB when building the environment.
	Create(name string) error
}

// DBAddr is the database address.
type roach struct {
	addr string
}

// NewRoach ...
func NewRoach() *roach {
	var (
		host     = viper.GetString("cockroach.host")
		port     = viper.GetString("cockroach.port")
		user     = viper.GetString("cockroach.user")
		password = viper.GetString("cockroach.password")
		dbName   = viper.GetString("cockroach.database")
		sslMode  = viper.GetString("cockroach.ssl")
		timeZone = viper.GetString("cockroach.timezone")
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

	return &roach{addr: addr}
}

func (d *roach) Dial() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return errors.New("CONNECT_TO_COCKROACH_DB_TIMEOUT")
		default:
			dbs, err := gorm.Open(postgres.Open(d.addr), &gorm.Config{})
			if err != nil {
				return err
			}
			log.Println("DATABASE_CONNECTION_SUCCESSFUL")
			db = dbs
			return nil
		}
	}
}

func (d *roach) Create(name string) error {
	c, err := sql.Open("postgres", d.addr)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := c.Exec("CREATE DATABASE " + name); err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
