package models

import (
	"gorm.io/gorm"
) 


type UserBasic struct{
	gorm.Model
	Name string
	Password string
}

