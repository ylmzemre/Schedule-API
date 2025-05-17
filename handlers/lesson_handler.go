package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ylmzemre/Schelude-API/models"
	"github.com/ylmzemre/Schelude-API/service"
)

type LessonHandler struct {
	svc service.LessonService
}

func NewLessonHandler(s service.LessonService) *LessonHandler {
	return &LessonHandler{svc: s}
}

// POST /api/lessons
func (h *LessonHandler) Create(c *fiber.Ctx) error {
	var l models.Lesson
	if err := c.BodyParser(&l); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if err := h.svc.Create(&l); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(l)
}
