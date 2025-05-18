package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ylmzemre/Schelude-API/config"
	"github.com/ylmzemre/Schelude-API/handlers"
	"github.com/ylmzemre/Schelude-API/models"
	"github.com/ylmzemre/Schelude-API/repository"
	"github.com/ylmzemre/Schelude-API/routes"
	"github.com/ylmzemre/Schelude-API/service"
	"log"
)

func main() {
	config.Load()
	config.InitDB()
	// Auto migrate
	config.DB.AutoMigrate(&models.Student{}, &models.Lesson{})

	app := fiber.New()

	// DI
	lessonRepo := repository.NewLessonRepo()
	lessonSvc := service.NewLessonService(lessonRepo)
	lessonH := handlers.NewLessonHandler(lessonSvc)

	studentRepo := repository.NewStudentRepo()
	studentSvc := service.NewStudentService(studentRepo)
	studentH := handlers.NewStudentHandler(studentSvc)
	authH := handlers.NewAuthHandler(studentSvc)

	routes.Register(app, authH, studentH, lessonH)
	port := config.Cfg.Port
	log.Printf("🚀  API is running on http://localhost:%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Fiber Listen error: %v", err)
	}

	log.Fatal(app.Listen(":" + config.Cfg.Port))

}
