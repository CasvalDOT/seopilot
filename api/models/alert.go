package models

import (
	"time"

	"gorm.io/gorm"
)

type Alert struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type AlertViewResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AlertViewSimpleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AlertViewAnyResponse struct {
	Data []AlertViewResponse `json:"data"`
}

type AlertCreateRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=255"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

type AlertUpdateRequest struct {
	Name        string `json:"name" binding:"min=2,max=255"`
	Description string `json:"description" binding:"omitempty,max=255"`
}
