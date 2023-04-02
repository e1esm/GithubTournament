package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int     `gorm:"primaryKey:id"`
	Username string  `gorm:"column:username"`
	Link     string  `gorm:"column:link"`
	Commits  int     `gorm:"column:commits"`
	Streak   int     `gorm:"column:streak"`
	Rating   float64 `gorm:"column:rating"`
}
