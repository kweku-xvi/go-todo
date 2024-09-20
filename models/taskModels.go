package models

import (
	"time"
)

type Task struct {
	ID          string    `gorm:"column:id;primaryKey;unique" json:"id"`
	Title       string    `gorm:"column:title;not null" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	Deadline    time.Time `gorm:"column:deadline" json:"deadline"`
	Priority    string    `gorm:"column:priority" json:"priority"`
	Status      string    `gorm:"column:status;default:'Pending'" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
