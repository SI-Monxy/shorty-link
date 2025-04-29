package config

import (
	"log"
	"shorty-link/internal/infrastructure/db/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := "root:password@tcp(mysql:3306)/shorty?charset=utf8mb4&parseTime=True&loc=Local"

	var db *gorm.DB
	var err error

	// リトライ最大10回
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("waiting for db to be ready...")
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("failed to connect database after retries: %v", err)
	}

	if err := db.AutoMigrate(&model.ShortURL{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	return db
}
