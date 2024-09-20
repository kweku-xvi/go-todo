package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Deadline    time.Time `gorm:"column:deadline" json:"deadline"`
	Priority    string    `gorm:"column:priority" json:"priority"`
	Status      string    `gorm:"column:status;default:'Pending'" json:"status"`
}
