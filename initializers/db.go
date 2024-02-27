package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
  var err error
	databaseUrl := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

  if err != nil {
    log.Fatal("Failed to connect to database")
  }
}
