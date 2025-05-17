package service

import (
	"github.com/ylmzemre/Schelude-API/models"
	"github.com/ylmzemre/Schelude-API/repository"
)

type LessonService interface {
	Create(*models.Lesson) error
}

type lessonSvc struct{ repo repository.LessonRepository }

func NewLessonService(r repository.LessonRepository) LessonService {
	return &lessonSvc{repo: r}
}
func (s *lessonSvc) Create(l *models.Lesson) error {
	return s.repo.Create(l)
}
