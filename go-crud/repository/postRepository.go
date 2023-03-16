package repository

import (
	"errors"
	"gorm.io/gorm"
	"john/gin-curd/common"
	"john/gin-curd/models"
	"john/gin-curd/vo"
)

type IPostRepository interface {
	Create(post vo.PostRequest) error
	Delete(id string) error
	Update(post models.Post) error
	Query(id string) error
}

type PostRepository struct {
	DB *gorm.DB
}

func (p *PostRepository) Delete(id string) error {
	if err := p.DB.Where("id", id).Delete(&models.Post{}).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostRepository) Update(post models.Post) error {
	if err := p.DB.Where("id = ?", post.ID).Updates(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p *PostRepository) Query(id string) (models.Post, error) {
	var post = models.Post{}
	p.DB.Where("id = ?", id).First(&post)
	if len(post.ID) == 0 {
		return post, errors.New("文章不存在")
	}
	return post, nil
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		DB: common.GetDB(),
	}
}

func (p *PostRepository) Create(post vo.PostRequest) error {
	var CPost = models.Post{
		CategoryId: post.CategoryId,
		HeadImg:    post.HeadImg,
		Title:      post.Title,
		Content:    post.Content,
	}

	if err := p.DB.Create(&CPost).Error; err != nil {
		return err
	}
	return nil

}

func (p *PostRepository) SelectById(id string) (*models.Post, error) {
	var post = models.Post{}
	if err := p.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
