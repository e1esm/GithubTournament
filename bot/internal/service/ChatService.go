package service

import (
	"XDaysOfCodeBot/internal/models"
	"XDaysOfCodeBot/internal/repository"
)

type ChatService struct {
	chatRepository repository.ChatRepository
}

func NewChatService(chatRepository repository.ChatRepository) *ChatService {
	return &ChatService{chatRepository: chatRepository}
}

func (c *ChatService) NewChat(chat models.Chat) {
	c.chatRepository.DB.Create(&chat)
}

func (c *ChatService) FindOne(chatId int64) (models.Chat, bool) {
	isStored := true
	chat := models.Chat{ChatID: chatId}
	c.chatRepository.DB.Find(&chat)
	if chat.UserId == 0 {
		isStored = false
	}
	return chat, isStored
}
