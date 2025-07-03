package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=1234 dbname=tprueba port=5432"
var DB *gorm.DB

func Connection(){
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}