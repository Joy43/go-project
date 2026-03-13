package models

import "github.com/jinzhu/gorm"

// User model to represent users in the database
type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null" json:"username"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
	Role     string `gorm:"size:20;not null;default:'user'" json:"role"`
	Posts    []Post `json:"posts"`
	Comments []Comment `json:"comments"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"size:255;not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `json:"user_id"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	User    User   `json:"user"`
}
