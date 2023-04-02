package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	User   User `gorm:"foreignKey:UserId"`
	UserId int
	ChatID int64 `gorm:"primaryKey:ChatId"`
}
