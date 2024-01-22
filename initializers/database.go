package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func ConnectToDB(){
    var err error
	
	dsn := "host=babar.db.elephantsql.com user=zaaoeawg password=VKmcbCroIUdwwB6ESnwRF6y9mJ0e7c7W dbname=zaaoeawg port=5432 sslmode=disable"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    
	 if err != nil{
		log.Fatal("Failed to connect to database")
	 }

}