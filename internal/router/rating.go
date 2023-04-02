package router

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func (r *Router) GetRating(message tgbotapi.Message) {
	strId := strconv.Itoa(int(message.Chat.ID))
	msg := tgbotapi.NewMessage(message.Chat.ID, "hey"+strId+" yo bitch")
	r.bot.Send(msg)
}
