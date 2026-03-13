package models

import "github.com/jinzhu/gorm"

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"size:100;not null" json:"username" example:"joy"`
	Email    string `gorm:"size:100;not null;unique" json:"email" example:"joy@example.com"`
	Password string `gorm:"size:100;not null" json:"password" example:"123456"`
	Role     string `gorm:"size:20;not null;default:'user'" json:"role" example:"user"`
	Posts    []Post    `json:"posts"`
	Comments []Comment `json:"comments"`
}

// Post represents a blog post
type Post struct {
	gorm.Model
	Title    string `gorm:"size:255;not null" json:"title" example:"My First Post"`
	Content  string `gorm:"type:text;not null" json:"content" example:"This is my first post content"`
	UserID   uint   `json:"user_id" example:"1"`
	Comments []Comment `json:"comments"`
}

// Comment represents a comment on a post
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content" example:"Nice post!"`
	UserID  uint   `json:"user_id" example:"1"`
	PostID  uint   `json:"post_id" example:"1"`
	User    User   `json:"user"`
}