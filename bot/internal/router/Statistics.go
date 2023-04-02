package router

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (r *Router) GetStatistics(message tgbotapi.Message) {
	chat, _ := r.chatService.FindOne(message.Chat.ID)
	user := r.tournamentService.FindByChat(chat)
	userOutput := fmt.Sprintf("*User information*\n===\n*Username:* %s\n*Link:* %s\n*Rating:* %.3f\n*Commits:* %d\n*Streak:* %d\n===",
		user.Username, user.Link, user.Rating, user.Commits, user.Streak)
	msg := tgbotapi.NewMessage(message.Chat.ID, userOutput)
	msg.ParseMode = "markdown"
	r.bot.Send(msg)
}
