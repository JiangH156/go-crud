package dto

import "john/gin-curd/models"

type PostDto struct {
	HeadImg string `json:"head_img"`
	Title   string `json:"title" gorm:"type:varchar(50);not null"`
	Content string `json:"content" gorm:"type:text;not null"`
}

func ToPostDto(post models.Post) PostDto {
	return PostDto{
		HeadImg: post.HeadImg,
		Title:   post.Title,
		Content: post.Content,
	}
}
