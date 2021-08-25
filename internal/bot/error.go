package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleError(chatID int64, err error) {
	messageText := err.Error()

	// log error into db

	msg := tgbotapi.NewMessage(chatID, messageText)
	b.bot.Send(msg)
}