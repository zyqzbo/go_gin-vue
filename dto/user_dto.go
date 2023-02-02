package dto

import (
	"go_gin+vue/model"
)

type UserDto struct { // 重新定义一个结构体来过滤请求数据的password
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDto(user model.User) UserDto { // 定义一个值返用户名和用户手机号码的方法
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
