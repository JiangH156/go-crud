package vo

type PostRequest struct {
	CategoryId uint   `json:"category_id" binding:"required"`
	HeadImg    string `json:"head_img"`
	Title      string `json:"title" binding:"required,max=10"`
	Content    string `json:"content" binding:"required"`
}
