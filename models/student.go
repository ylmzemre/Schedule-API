package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `gorm:"uniqueIndex" json:"email"`
	Password  string `json:"-"`

	Lessons []Lesson `gorm:"many2many:student_lessons" json:"lessons,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
