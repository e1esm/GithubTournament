package repository

import "gorm.io/gorm"

type ChatRepository struct {
	DB *gorm.DB
}

func NewChatRepository(DB *gorm.DB) *ChatRepository {
	return &ChatRepository{DB: DB}
}
