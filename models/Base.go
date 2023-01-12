package models

import (
	"gorm.io/gorm"
)

type BaseModel struct {
	gorm.Model
}


// func (model  *BaseModel) FindById(id string) BaseModel{

// 	return nil
// }