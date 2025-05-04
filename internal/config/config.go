package config

import (
	"fmt"
	"log"
	"os"
	"shorty-link/internal/infrastructure/db/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// 環境変数から取得
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

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
