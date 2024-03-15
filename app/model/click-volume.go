package model

import "gorm.io/gorm"

type ClickVolume struct {
	gorm.Model
	Web  uint `gorm:"not null;type:bigint;comment:点击次数"`
	Note uint `gorm:"not null;type:bigint;comment:笔记浏览量"`
}
