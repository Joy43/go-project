package models

import "github.com/jinzhu/gorm"

// User model to represent users in the database
type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}
