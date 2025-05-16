package models

import (
	"gorm.io/gorm"
	"time"
)

type Lessons struct {
	gorm.Model
	lessonName  string    `gorm:"type:varchar(30); not null;"`
	lessonID    int       `gorm:"type:int; not null; unique;"`
	lessonStart time.Time `gorm:"type:timestamp; not null;"`
	lessonEnd   time.Time `gorm:"type:timestamp; not null;"`
}

func (receiver Lessons) TableName() string {
	return "lessons"
}
