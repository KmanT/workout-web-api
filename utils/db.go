package utils

import (
	"fmt"
	"log"
	"os"

	"../config"
	"../model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	fmt.Println("Connecting to database...")
	var cfg config.Config
	port := os.Getenv("PORT")

	if port == "" {
		cfg = config.InitConfig(true)
	} else {
		cfg = config.InitConfig(false)
	}

	connection := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Dbname, cfg.Database.Password)
	db, err := gorm.Open("postgres", connection)

	if err != nil {
		log.Fatal(err)
	}

	//defer db.Close()

	fmt.Println("Migrating tables...")
	db.AutoMigrate(&model.User{}, &model.Exercise{})

	fmt.Println("Database connected.")
	return db
}
