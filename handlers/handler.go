package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	DB *gorm.DB
}
type LessonSchedule struct {
	LessonName string `json:"lessonName"`
}

func (h *Handler) LessonScheduleHandler(c echo.Context) error {
	names := LessonSchedule{
		LessonName: "Mathematics",
	}
	return c.JSON(http.StatusOK, names)
}
