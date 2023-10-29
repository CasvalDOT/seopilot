package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        int
	Name      string
	Users     []User `gorm:"many2many:users_have_roles;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type RoleSimple struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
