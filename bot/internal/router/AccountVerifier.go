package router

import (
	models2 "XDaysOfCodeBot/internal/models"
	"fmt"
	"github.com/enescakir/emoji"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"net/http"
)

func (r *Router) VerifyAccount(message tgbotapi.Message) {
	link := "https://github.com/"
	arg := message.CommandArguments()
	res, err := http.Get(link + arg)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, errorsMap[RequestError])
		r.bot.Send(msg)
	}
	if res.StatusCode != 200 {
		msg := tgbotapi.NewMessage(message.Chat.ID, errorsMap[InvalidUsername])
		r.bot.Send(msg)
	} else {
		user := models2.User{Username: arg, Link: link + arg}
		_, isGithubLinked := r.chatService.FindOne(message.Chat.ID)
		_, isGithubTaken := r.tournamentService.FindOne(user)
		if isGithubLinked {
			msg := tgbotapi.NewMessage(message.Chat.ID, errorsMap[AlreadyHaveAccount])
			r.bot.Send(msg)
		} else if isGithubTaken {
			msg := tgbotapi.NewMessage(message.Chat.ID, errorsMap[BusyAccount])
			r.bot.Send(msg)
		} else {
			r.chatService.NewChat(models2.Chat{ChatID: message.Chat.ID, User: user})
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Github account was found and linked with your current Telegram account %v\n", emoji.GreenCircle))
			r.bot.Send(msg)
		}
	}
}
