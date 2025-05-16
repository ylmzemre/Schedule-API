package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/ylmzemre/Schelude-API/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Application struct {
	logger  echo.Logger
	server  *echo.Echo
	handler handlers.Handler
}

func main() {
	e := echo.New()
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	port2 := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, name, port2,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	h := handlers.Handler{
		DB: db,
	}
	app := Application{
		logger:  e.Logger,
		server:  e,
		handler: h,
	}
	e.GET("/", h.LessonScheduleHandler)
	fmt.Println(app)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
