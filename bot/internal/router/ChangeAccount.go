package router

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (r *Router) ChangeAccount(message tgbotapi.Message) {
	r.VerifyAccount(message)
}
