package router

import (
	"XDaysOfCodeBot/internal/models"
	"fmt"
	"github.com/enescakir/emoji"
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
		doesExist := r.chatService.FindOne(message.Chat.ID)
		if doesExist {
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("You already have linked account %v\n", emoji.RedCircle))
			r.bot.Send(msg)
		} else {
			r.tournamentService.NewUser(user)
			r.chatService.NewChat(models.Chat{ChatID: message.Chat.ID, User: user})
			msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Github account was found %v\n", emoji.GreenCircle))
			r.bot.Send(msg)
		}
	}
}
