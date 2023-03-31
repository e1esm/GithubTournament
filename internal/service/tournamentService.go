package service

import (
	"XDaysOfCodeBot/internal/models"
	"XDaysOfCodeBot/internal/repository"
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
