package models

import "gorm.io/gorm"

type Lesson struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	DayOfWeek   int            `json:"day_of_week"`                                         // 1=Mon ... 7=Sun
	StartTime   string         `json:"start_time"`                                          // "14:00"
	EndTime     string         `json:"end_time"`                                            // "15:30"
	Students    []Student      `json:"students,omitempty" gorm:"many2many:student_lessons"` // <- burası
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
