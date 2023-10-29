package models

import (
	"time"

	"gorm.io/gorm"
)

type AlertRule struct {
	ID        int
	Content   string
	AlertID   int
	Alert     Alert
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
