package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *Router) PickCommand(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "add":
		r.VerifyAccount(*update.Message)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Invalid user")
		r.bot.Send(msg)
	}
}
