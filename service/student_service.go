package service

import (
	"errors"
	"github.com/ylmzemre/Schelude-API/models"
	"github.com/ylmzemre/Schelude-API/repository"
	"github.com/ylmzemre/Schelude-API/utils"
)

type StudentService interface {
	Register(*models.Student) (*models.Student, error)
	Find(id uint) (*models.Student, error)
	Login(email, password string) (*models.Student, error)
}

type studentSvc struct {
	repo repository.StudentRepository
}

func NewStudentService(r repository.StudentRepository) StudentService {
	return &studentSvc{repo: r}
}

func (s *studentSvc) Register(st *models.Student) (*models.Student, error) {
	// email uniqueness kontrolü örnek
	if _, err := s.repo.GetByEmail(st.Email); err == nil {
		return nil, errors.New("email already exists")
	}
	st.Password = utils.HashPassword(st.Password)
	if err := s.repo.Create(st); err != nil {
		return nil, err
	}
	return st, nil
}

func (s *studentSvc) Login(email, password string) (*models.Student, error) {
	st, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHash(password, st.Password) {
		return nil, errors.New("invalid credentials")
	}
	return st, nil
}

func (s *studentSvc) Find(id uint) (*models.Student, error) {
	return s.repo.Find(id)
}
