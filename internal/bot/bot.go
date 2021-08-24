package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	db *dbs.Postgres	// для взаимодействия с бд
}