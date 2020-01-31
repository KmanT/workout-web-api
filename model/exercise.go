package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Exercise struct {
	gorm.Model
	UserId uint      `json:"user_id"`
	Date   time.Time `json:"date"`
	Name   string    `gorm:"type:varchar(100);not_null" json:"name"`
	Sets   uint      `gorm:"not_null" json:"sets"`
	Reps   uint      `json:"reps"`
	Weight float32   `json:"weight"`
	Unit   string    `json:"unit"`
	Time   uint      `json:"time"` // milliseconds
}
