package vo

type CreatePostRequst struct {
	CategoryId uint   `json:"category_id" binding:"required"` // binding:"required" 用于表单验证
	Title      string `json:"title" binding:"required,max=10"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" binding:"required"`
}
