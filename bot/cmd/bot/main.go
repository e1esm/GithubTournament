package main

import (
	configuration "XDaysOfCodeBot/cmd/config"
	models2 "XDaysOfCodeBot/internal/models"
	repository2 "XDaysOfCodeBot/internal/repository"
	"XDaysOfCodeBot/internal/router"
	service2 "XDaysOfCodeBot/internal/service"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	token := os.Getenv("TOKEN")
	dbconfig, err := configuration.NewConfig()

	if err != nil {
		log.Fatal(err)
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbconfig.DBUser, dbconfig.DBPassword, dbconfig.DBHost, dbconfig.DBPort, dbconfig.DBName)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models2.User{}, &models2.Chat{})

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	userRepository := repository2.NewUserRepository(db)
	chatRepository := repository2.NewChatRepository(db)
	chatService := service2.NewChatService(*chatRepository)
	tournamentService := service2.NewTournamentService(*userRepository)
	rtr := router.NewRouter(bot, *tournamentService, *chatService)

	for update := range updates {
		rtr.HandleUpdate(update)
	}

}
