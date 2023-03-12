package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"name" json:"name" form:"name"`
	Password  string `gorm:"password" json:"password" form:"password"`
	Telephone string `gorm:"telephone" json:"telephone" form:"telephone"`
}
