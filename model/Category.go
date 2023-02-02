package model

type Category struct { // 分类结构体字段
	ID   uint   `json:"id" gorm:"primary_key"` // gorm 默认ID是自增字段
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	// 下面的两个字段gorm默认结构体里面有字段 CreateAdd UpdateAt 就会自动赋值
	//但是由于数据库版本是8.0.27问题 对时间类型做序列化和反序列号处理
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"`
}
