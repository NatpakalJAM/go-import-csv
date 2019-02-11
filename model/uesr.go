package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID        int         `gorm:"column:id" json:"id"`
	Firstname interface{} `gorm:"column:firstname" json:"firstname"`
	Lastname  interface{} `gorm:"column:lastname" json:"lastname"`
}

// TableName for model
func (User) TableName() string {
	return "user"
}
