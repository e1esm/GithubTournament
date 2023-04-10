package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *Router) PickCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "add":
		r.VerifyAccount(*update.Message)
	case "statistics":
		r.GetStatistics(*update.Message)
	case "change":
		r.ChangeAccount(*update.Message)
	case "chart":
		r.Chart(*update.Message, tgbotapi.CallbackQuery{})
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid user")
		r.bot.Send(msg)
	}
}
