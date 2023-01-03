package models

import (
	_ "gorm.io/gorm"
)

type Group struct {
	BaseModel BaseModel `gorm:"embedded"`

	Name string // 名称  `gorm:"unique_index:idx_name"`
	Img  string // 图片链接
}
