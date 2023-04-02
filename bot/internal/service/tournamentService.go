package service

import (
	"XDaysOfCodeBot/bot/internal/models"
	"XDaysOfCodeBot/bot/internal/repository"
)

type TournamentService struct {
	repository repository.UserRepository
}

func NewTournamentService(userRepository repository.UserRepository) *TournamentService {
	return &TournamentService{repository: userRepository}
}

func (t *TournamentService) NewUser(user models.User) {
	t.repository.DB.Create(&user)
}

func (t *TournamentService) FindOne(user models.User) (models.User, bool) {
	isFound := false
	var found models.User
	t.repository.DB.Where("username = ?", user.Username).Find(&found)
	if found.Username != "" {
		isFound = true
	}
	return found, isFound
}

func (t *TournamentService) FindByChat(chat models.Chat) models.User {
	var user models.User
	t.repository.DB.Where("id = ?", chat.UserId).Find(&user)
	return user
}
