package main

import (
	configuration "XDaysOfCodeBot/cmd/config"
	"XDaysOfCodeBot/internal/models"
	"XDaysOfCodeBot/internal/repository"
	"XDaysOfCodeBot/internal/router"
	"XDaysOfCodeBot/internal/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("./cmd/bot/.env")
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	dbconfig, err := configuration.NewConfig()

	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbconfig.DBHost, dbconfig.DBUser, dbconfig.DBPassword, dbconfig.DBName, dbconfig.DBPort)), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	repository := repository.NewUserRepository(db)
	tournamentService := service.NewTournamentService(*repository)
	rtr := router.NewRouter(bot, *tournamentService)

	for update := range updates {
		rtr.HandleUpdate(update)
	}
}
