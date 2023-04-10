package router

import (
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

var currentOffset = 0
var inlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", emoji.LeftArrow), "-10"),
		tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%v", emoji.RightArrow), "10"),
	),
)

func (r *Router) Chart(message tgbotapi.Message, callback tgbotapi.CallbackQuery) {
	if callback != (tgbotapi.CallbackQuery{}) {
		decData, _ := strconv.Atoi(callback.Data)
		currentOffset += decData
	}
	users := r.tournamentService.FetchSome(currentOffset)
	header := "Chart\n\n=="
	bodyOfMessage := ""
	for _, user := range users {
		bodyOfMessage += "Username: " + user.Username + "Rating: " + fmt.Sprintf("%v\n", user.Rating) + "=====\n"
	}
	header += bodyOfMessage
	msg := tgbotapi.NewMessage(message.Chat.ID, header)
	msg.ReplyMarkup = inlineKeyboard
	r.bot.Send(msg)
}
