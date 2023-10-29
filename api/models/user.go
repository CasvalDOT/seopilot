package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Roles     []Role    `db:"roles" gorm:"many2many:users_have_roles;"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt gorm.DeletedAt
}

type UserViewResponse struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Roles     []RoleSimple `json:"roles"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type UserViewAnyResponse struct {
	Data []UserViewResponse `json:"data"`
}

type UserCreateRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=255"`
	Email string `json:"email" binding:"required,email,max=255"`
	Roles []int  `json:"roles" binding:"required,unique,gt=0"`
}

type UserUpdateRequest struct {
	Name  string `json:"name" binding:"min=2,max=255"`
	Email string `json:"email" binding:"email,max=255"`
	Roles []int  `json:"roles" binding:"unique,gt=0"`
}
