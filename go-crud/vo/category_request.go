package vo

// 视图层
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}
