package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique" json:"username"`
	Email    string `gorm:"size:255;not_null;unique" json:"email"`
	Password string `gorm:"size:255;not_null" json:"password"`
	Exercise []Exercise
}
