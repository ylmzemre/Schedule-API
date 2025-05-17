package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ylmzemre/Schelude-API/handlers"
	"github.com/ylmzemre/Schelude-API/middlewares"
)

func Register(app *fiber.App, hAuth *handlers.AuthHandler,
	hStudent *handlers.StudentHandler, hLesson *handlers.LessonHandler) {
	// Public
	app.Post("/register", hAuth.Register)
	app.Post("/login", hAuth.Login)

	// Protected
	api := app.Group("/api", middlewares.JWTProtected())
	api.Get("/students/:id", hStudent.GetByID)
	api.Post("/lessons", hLesson.Create)
}
