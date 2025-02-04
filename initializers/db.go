package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/wpcodevo/golang-fiber/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Config env! \n", config.DBUserName, config.DBHost, config.DBPort, config.DBUserPassword)
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	// DB.Exec(`CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"`)
	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	DB.Table("notes").AutoMigrate(&models.Note{})

	log.Println("🚀 Connected Successfully to the Database")
}
