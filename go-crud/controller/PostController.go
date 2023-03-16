package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"john/gin-curd/common"
	"john/gin-curd/dto"
	"john/gin-curd/models"
	"john/gin-curd/repository"
	"john/gin-curd/response"
	"john/gin-curd/vo"
	"net/http"
)

type IPostController interface {
	RestControoler
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	DB := common.GetDB()
	DB.AutoMigrate(&models.Post{})
	return PostController{DB: DB}
}

func (p PostController) Create(ctx *gin.Context) {
	var postRepository = repository.NewPostRepository()
	// 验证输入数据
	requestPost := vo.PostRequest{}
	if err := ctx.ShouldBind(&requestPost); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "数据验证错误，请重新输入分类名称、文章名称、内容")
		return
	}
	// 此处不应该判断是否存在userid和categoryid，选择文章分类时，应该是选择已存在id

	if err := postRepository.Create(requestPost); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "新增文章失败")
		return
	}

	response.Success(ctx, gin.H{"post": requestPost}, "新增文章成功")

}

func (p PostController) Delete(ctx *gin.Context) {
	var postRepository = repository.NewPostRepository()

	delID := ctx.Params.ByName("id")

	if err := postRepository.Delete(delID); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "删除文章失败")
		return
	}
	response.Success(ctx, gin.H{"postID": delID}, "删除文章成功")

}

func (p PostController) Update(ctx *gin.Context) {
	var postRepository = repository.NewPostRepository()
	updPost := models.Post{}
	// 数据验证
	ctx.ShouldBind(&updPost)

	updID := ctx.Params.ByName("id")
	_, err := postRepository.SelectById(updID)
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "文章不存在")
		return
	}

	updPost.ID = updID
	// 更新文章
	if err := postRepository.Update(updPost); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新文章失败")
		return
	}

	response.Success(ctx, gin.H{"post": dto.ToPostDto(updPost)}, "删除文章成功")
}

func (p PostController) Query(ctx *gin.Context) {
	var postRepository = repository.NewPostRepository()

	delID := ctx.Params.ByName("id")

	selPost, err := postRepository.Query(delID)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "查询文章失败")
		return
	}
	response.Success(ctx, gin.H{"post": dto.ToPostDto(selPost)}, "查询文章成功")
}
