package models

import "gorm.io/gorm"

type Student struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Email     string         `json:"email" gorm:"uniqueIndex"`
	Password  string         `json:"-"`                 // hashed
	Lessons   []Lesson       `json:"lessons,omitempty"` // many-to-many
	CreatedAt int64          `json:"created_at"`
	UpdatedAt int64          `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
