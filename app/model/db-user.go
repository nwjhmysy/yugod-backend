package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null;type:varchar(255);comment:账号"`
	Name     string `gorm:"not null;type:varchar(255);comment:姓名"`
	Password string `gorm:"not null;type:varchar(255);comment:密码（登陆）"`
	Auth     uint   `gorm:"not null;type:tinyint;comment:(1,游客;2,用户;)"`
	Email    string `gorm:"type:varchar(255);comment:邮箱"`
	Age      uint   `gorm:"type:tinyint;comment:年龄"`
}
