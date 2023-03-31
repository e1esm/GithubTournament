package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int `gorm:"primaryKey"`
	Username string
	Link     string
	Commits  int
	Streak   int
	Rating   float64
}
