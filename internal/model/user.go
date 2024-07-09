package model

import (
	"time"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
	Username  string     `json:"username" gorm:"unique;not null" example:"exampleuser"`
	Password  string     `json:"password" gorm:"not null" example:"examplepassword"`
}

type LoginInput struct {
	Username string `json:"username" validate:"required" example:"exampleuser"`
	Password string `json:"password" validate:"required" example:"examplepassword"`
}
