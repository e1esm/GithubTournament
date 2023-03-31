package router

import (
	"XDaysOfCodeBot/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
)

func (r *Router) VerifyAccount(message tgbotapi.Message) {
	link := "https://github.com/"
	arg := message.CommandArguments()
	res, err := http.Get("https://github.com/" + arg)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Try to input another username")
		r.bot.Send(msg)
	} else {
		user := models.User{Username: arg, Link: link + arg}
		r.tournamentService.NewUser(user)
		msg := tgbotapi.NewMessage(message.Chat.ID, "found this account")
		r.bot.Send(msg)
	}
}
