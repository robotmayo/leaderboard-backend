package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  sql.NullString
	TwitterID string `gorm:"unique"`
}

type UserIdentifier struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type UserPersonal struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserRegister struct {
	Username        string `json:"username" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"eqfield=Password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}