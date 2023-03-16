package models

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         string `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint   `json:"user_id" gorm:"not null"`
	CategoryId uint   `json:"category_id" gorm:"not null"`
	HeadImg    string `json:"head_img"`
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
	Content    string `json:"content" gorm:"type:text;not null"`
}

// gorm钩子函数，
func (post *Post) BeforeCreate(db *gorm.DB) error {
	post.ID = uuid.NewV4().String()
	fmt.Println(post.ID)
	return nil
}
