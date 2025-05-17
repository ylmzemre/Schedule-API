package repository

import (
	"github.com/ylmzemre/Schelude-API/config"
	"github.com/ylmzemre/Schelude-API/models"
)

type LessonRepository interface {
	Create(*models.Lesson) error
}

type lessonRepo struct{}

func NewLessonRepo() LessonRepository { return &lessonRepo{} }

func (r *lessonRepo) Create(l *models.Lesson) error {
	return config.DB.Create(l).Error
}
