package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	BaseModel BaseModel `gorm:"embedded"`

	Name   string    `gorm:"unique_index:idx_name_phone"`   // 名称
	Sex    bool      // 性别 男0女1
	Bir    time.Time // 生日
	Img    string    // 图片链接
	Phone  string    `gorm:"unique_index:idx_name_phone"`  // 电话
	UserID string      // uuid

	RegisteredTime time.Time // 注册时间
	LoginTime      time.Time // 登录时间
	LogoutTime     time.Time // 退出时间
}
