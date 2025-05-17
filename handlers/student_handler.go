package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ylmzemre/Schelude-API/service"
)

// HTTP katmanı
type StudentHandler struct {
	svc service.StudentService
}

func NewStudentHandler(s service.StudentService) *StudentHandler {
	return &StudentHandler{svc: s}
}

// GET /api/students/:id
func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	student, err := h.svc.Find(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(student)
}
