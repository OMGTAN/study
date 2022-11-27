package models

import (
	"gorm.io/gorm"
)

// swagger:model UserBasic
type UserBasic struct {

	//"id":"1",
	gorm.Model `swaggerignore:"true"`
	Name       string `json:"name" example:"name1"`
	Password   string `json:"password"`
}
