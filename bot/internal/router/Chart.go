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
	body := "Chart\n\n===="
	chart := ""
	for _, user := range users {
		chart += "Username: " + user.Username + "Rating: " + fmt.Sprintf("%v\n", user.Rating) + "=====\n"
	}
	body += chart
	msg := tgbotapi.NewMessage(message.Chat.ID, body)
	msg.ReplyMarkup = inlineKeyboard
	r.bot.Send(msg)
}
