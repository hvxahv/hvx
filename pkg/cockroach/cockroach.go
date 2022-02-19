package cockroach

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Cockroach is a wrapper around a cockroach database connection.
var db *gorm.DB

// GetDB returns the database connection.
func GetDB() *gorm.DB {
	return db
}

type DB interface {
	// InitDB Call the configuration file to initialize and connect to the database, If the connection is incorrect, a custom error is returned.
	// If the connection fails, there will be a 10second retry connection time.
	InitDB() error

	// Create a database method for hvxctl to create a DB when building the environment.
	Create(name string) error
}

// DBAddr is the database address.
type DBAddr struct {
	Addr  string
	Error error
}

// NewDBAddr returns a new DBAddr.
func NewDBAddr() *DBAddr {
	// Get the database address from the configuration file.
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return errors.New("CONNECT_TO_COCKROACH_DB_TIMEOUT")
		default:
			dbs, err := gorm.Open(postgres.Open(d.Addr), &gorm.Config{})
			if err != nil {
				d.Error = err
			} else {
				log.Println("DATABASE_CONNECTION_SUCCESSFUL")
				db = dbs
				return nil
			}
		}
	}
}

func (d *DBAddr) Create(name string) error {

	c, err := sql.Open("postgres", d.Addr)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := c.Exec("CREATE DATABASE " + name); err != nil {
		return err
	}
	return nil
}
