package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID        int         `gorm:"column:id" csv:"id" json:"id"`
	Firstname interface{} `gorm:"column:firstname" csv:"firstname" json:"firstname"`
	Lastname  interface{} `gorm:"column:lastname" csv:"lastname" json:"lastname"`
}

// TableName for model
func (User) TableName() string {
	return "user"
}
