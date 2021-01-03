package app

import (
	"github.com/jinzhu/gorm"
	"hvxahv/pkg/database"
	"log"
)

var db2 *gorm.DB

func InitDB() {
	db, err := database.NewDB()
	if err != nil {
		log.Println("Database Connect Error", err)
	}
	db2 = db
}

