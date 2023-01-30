package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);no null"`
	Phone    string `gorm:"varchar(100;not null;unique"`
	Password string `gorm:"size:255;not null"`
}
