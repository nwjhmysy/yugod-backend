package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"not null;type:varchar(255);comment:姓名"`
	Password string `gorm:"not null;type:varchar(255);comment:密码（登陆）"`
	Email    string `gorm:"unique;not null;type:varchar(255);comment:邮箱（登陆）"`
	Age      uint   `gorm:"type:tinyint;comment:年龄"`
}
