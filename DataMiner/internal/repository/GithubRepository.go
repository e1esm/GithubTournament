package repository

import (
	"gorm.io/gorm"
)

type GithubRepository struct {
	DB *gorm.DB
}

func NewGithubRepository(DB *gorm.DB) *GithubRepository {
	return &GithubRepository{DB: DB}
}
