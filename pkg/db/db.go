package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	InitPostgre() error
	InitMysql() error
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

func NewDb(host string, port string, user string, password string, dbName string, sslMode string) *db {
	return &db{host: host, port: port, user: user, password: password, dbName: dbName, sslMode: sslMode}
}

var sdb *gorm.DB

func GetDB() *gorm.DB {
	return sdb
}

// InitPostgre Initialize the PostgreSQL database.
func (d *db) InitPostgre() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai",
		d.host,
		d.user,
		d.password,
		d.dbName,
		d.port,
		d.sslMode,
	)

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	sdb = dbs
	return nil
}

// InitMysql Initialize the MySQL database.
func (d *db) InitMysql() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.user,
		d.password,
		d.host,
		d.dbName,
	)
	dbs, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sdb = dbs
	return nil
}

// Create ...
func (d *db) Create(drive, name string) error {
	var addr string
	switch drive {
	case "mysql":
		addr = fmt.Sprintf("%s:%s@tcp(%s:%s)", d.user, d.password, d.host, d.port)
	case "postgres":
		addr = fmt.Sprintf("port=%s user=%s password=%s host=%s sslmode=%s",
			d.port,
			d.user,
			d.password,
			d.host,
			d.sslMode)
	}

	cdb, err := sql.Open(drive, addr)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := cdb.Exec("CREATE DATABASE " + name); err != nil {
		panic(err)
		return err
	}
	return nil
}
