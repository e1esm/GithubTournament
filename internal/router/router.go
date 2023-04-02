package router

import (
	"XDaysOfCodeBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TournamentRouter interface {
	HandleUpdate(update tgbotapi.Update)
	VerifyAccount(message tgbotapi.Message)
}

type Router struct {
	bot               *tgbotapi.BotAPI
	tournamentService service.TournamentService
	chatService       service.ChatService
}

func NewRouter(bot *tgbotapi.BotAPI, tournamentService service.TournamentService, chatService service.ChatService) *Router {
	return &Router{bot: bot, tournamentService: tournamentService, chatService: chatService}
}

func (r *Router) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil {
		if update.Message.IsCommand() {
			r.PickCommand(update)
		}
	}
}
