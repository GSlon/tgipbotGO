package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/sirupsen/logrus"

	s "github.com/GSlon/tgipbotGO/internal/service"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	service *s.Service
}

func NewBot(bot *tgbotapi.BotAPI, service *s.Service) *Bot {
	return &Bot{
		bot: bot,
		service: service,
	}
}

func (b *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-message 
			continue
		}

		// commands
		if update.Message.IsCommand() {
			if err := b.handleCommand(update.Message); err != nil {
				b.handleError(update.Message.Chat.ID, err)
			}

			continue
		}

		// regular messages
		if err := b.handleMessage(update.Message); err != nil {
			b.handleError(update.Message.Chat.ID, err)
		}
	}

	return nil
}

// обертка над send для логгирования в бд
func (b *Bot) SendMessage(chatid int64, message string) {
	msg := tgbotapi.NewMessage(chatid, message)
	_, err := b.bot.Send(msg)
	if err != nil {
		b.service.LogError(err.Error())
	}
}