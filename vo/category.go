package vo

type CreateCategoryRequest struct { // 对文字分类的所有字段请求进行封装请求
	Name string `json:"name" binding:"required"`
}
