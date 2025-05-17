package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBDsn     string
	JWTSecret string
	Port      string
}

var Cfg Config

func Load() {
	// .env’i oku (sessizce geçsin)
	_ = godotenv.Load()
	Cfg = Config{
		DBDsn:     buildDSN(), // <- DSN’i burada inşa ediyoruz
		JWTSecret: env("JWT_SECRET", "secret"),
		Port:      env("APP_PORT", "8080"),
	}
}

func buildDSN() string {
	// Eğer DB_DSN doğrudan verilmişse onu kullan
	if dsn := os.Getenv("DB_DSN"); dsn != "" {
		return dsn
	}
	host := env("DB_HOST", "127.0.0.1")
	port := env("DB_PORT", "5432")
	user := env("DB_USER", "postgres")
	password := env("DB_PASSWORD", "")
	dbname := env("DB_NAME", "postgres")
	sslmode := env("DB_SSLMODE", "disable")
	tz := env("DB_TIMEZONE", "UTC")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, tz)
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
