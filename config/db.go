package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB *gorm.DB
)

func GetDB() (*gorm.DB, error) {

	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, name, port,
	)
	log.Println(dsn)
	var err error
	DB, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("GORM bağlantı hatası: %v", err)
	}

	log.Println("PostgreSQL + GORM bağlantısı hazır")

	return DB, nil
}
