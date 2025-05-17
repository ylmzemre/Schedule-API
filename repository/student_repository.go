package repository

import (
	"github.com/ylmzemre/Schelude-API/config"
	"github.com/ylmzemre/Schelude-API/models"
)

type StudentRepository interface {
	Create(*models.Student) error
	GetByEmail(string) (*models.Student, error)
	Find(uint) (*models.Student, error)
}

type studentRepo struct{}

func NewStudentRepo() StudentRepository {
	return &studentRepo{}
}

func (r *studentRepo) Create(s *models.Student) error {
	return config.DB.Create(s).Error
}

func (r *studentRepo) GetByEmail(email string) (*models.Student, error) {
	var s models.Student
	err := config.DB.Where("email = ?", email).First(&s).Error
	return &s, err
}

func (r *studentRepo) Find(id uint) (*models.Student, error) {
	var s models.Student
	err := config.DB.Preload("Lessons").First(&s, id).Error
	return &s, err
}
