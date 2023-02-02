package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 文章结构体字段

type Post struct {
	ID     uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserID uint      `json:"user_id" gorm:"not null"`
	//Category 的外键 要复合 Category + Id = CategoryId 否则要重新指向具体的外键 具体看官网
	CategoryId uint `json:"category_id" gorm:"not null"` //Category 的外键 要复合 Category + Id = CategoryId
	Category   *Category
	Title      string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string `json:"head_img"`
	Content    string `json:"content" gorm:"type:text;not null"`
	CreatedAt  Time   `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) error { // 创建之前给id赋值
	post.ID = uuid.NewV4()
	return nil
}
