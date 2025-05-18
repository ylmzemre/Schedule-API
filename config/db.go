package config

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	sqlDB *sql.DB // underline connection pool
)

// InitDB zaten bağlanıyor; burada ping de attıracağız.
func InitDB() {
	dsn := Cfg.DBDsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm open error: %v", err)
	}

	// sql.DB handle
	sqlHandle, err := db.DB()
	if err != nil {
		log.Fatalf("sql.DB handle error: %v", err)
	}

	sqlHandle.SetMaxIdleConns(10)
	sqlHandle.SetMaxOpenConns(100)
	sqlHandle.SetConnMaxLifetime(time.Hour)

	DB = db
	sqlDB = sqlHandle

	// Bağlantıyı hemen test et
	if err := Ping(); err != nil {
		log.Fatalf("database ping error: %v", err)
	}

	log.Println("✅  Database connection established")
}

// Ping dışarıdan da çağrılabilsin
func Ping() error {
	return sqlDB.Ping()
}
