package Schedule_API

import (
	"fmt"
	"github.com/joho/godotenv"
)

func buildDSN() string {
	host := env("DB_HOST", "127.0.0.1")
	user := env("DB_USER", "postgres")
	password := env("DB_PASSWORD", "")
	name := env("DB_NAME", "postgres")
	port := env("DB_PORT", "5432")
	sslmode := env("DB_SSLMODE", "disable")
	tz := env("DB_TIMEZONE", "UTC")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, name, port, sslmode, tz)

}

func Load() {
	_ = godotenv.Load()
	Cfg = Config{
		DBDsn:     buildDSN(),
		JWTSecret: env("JWT_SECRET", "secret"),
		Port:      env("APP_PORT", "8080"),
	}
}
