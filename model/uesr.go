package model

import "time"

type User struct {
	ID        int       `gorm:"column:id;primary_key:true;AUTO_INCREMENT" csv:"id" json:"id"`
	Firstname string    `gorm:"column:firstname;size:50;index" csv:"firstname" json:"firstname"`
	Lastname  string    `gorm:"column:lastname;size:50;index" csv:"lastname" json:"lastname"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

// TableName for model
func (User) TableName() string {
	return "user"
}
