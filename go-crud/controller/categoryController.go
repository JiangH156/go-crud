package controller

import (
	"github.com/gin-gonic/gin"
	"john/gin-curd/models"
	"john/gin-curd/repository"
	"john/gin-curd/response"
	"john/gin-curd/vo"
	"net/http"
	"strconv"
)

// 分类的顶级抽象接口
type ICategoryController interface {
	RestControoler
}

// 分类功能的实现载体
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() CategoryController {
	repository := repository.NewICategoryRepository()
	repository.DB.AutoMigrate(&models.Category{})
	return CategoryController{
		Repository: repository,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var category = models.Category{}
	categoryRepository := repository.NewICategoryRepository()
	// 绑定数据
	var requestCategory = vo.CreateCategoryRequest{}

	// 验证数据
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "数据验证有误，分类名称必填")
		return
	}

	c.Repository.DB.Where("name = ?", category.Name).First(&category)
	if category.ID != 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "分类已存在，请重新输入分类")
		return
	}

	if err := categoryRepository.Create(category.Name); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "新增分类失败")
		return
	}

	response.Success(ctx, gin.H{"category": category}, "新增分类成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	var category = models.Category{}
	categoryRepository := repository.NewICategoryRepository()
	// URL中的路径是 string 类型
	categoryId := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(categoryId)

	if err := categoryRepository.Delete(id); err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "删除失败")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "删除成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var category = models.Category{}
	categoryRepository := repository.NewICategoryRepository()
	ctx.ShouldBind(&category)
	if len(category.Name) == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "数据输入有误，请重新输入")
		return
	}

	// URL中的路径是 string 类型
	categoryId := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(categoryId)
	if _, err := categoryRepository.SelectById(id); err != nil {
		response.Response(ctx, http.StatusNotFound, 404, nil, "分类不存在")
		return
	}

	Qcategory, err := categoryRepository.UpdateById(id, category.Name)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新失败")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"category": Qcategory}, "更新成功")
}

func (c CategoryController) Query(ctx *gin.Context) {
	categoryRepository := repository.NewICategoryRepository()

	categorytmp := ctx.Params.ByName("id")
	id, _ := strconv.Atoi(categorytmp)

	category, err := categoryRepository.Query(id)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "更新失败")
		return
	}
	response.Response(ctx, http.StatusOK, 200, gin.H{"category": category}, "查询成功")
}
