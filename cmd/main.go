package main

import (
    "github.com/spf13/viper"
    "github.com/sirupsen/logrus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

    b "github.com/GSlon/tgipbotGO/internal/bot"
	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"
	s "github.com/GSlon/tgipbotGO/internal/service"

	"os"
)

func initConfig() error {
    viper.AddConfigPath("config")
    viper.SetConfigName("config")
    return viper.ReadInConfig()
}

func main() {
    if err := initConfig(); err != nil {
        logrus.Fatalf(err.Error())
		return
    }

    config := dbs.PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	postgres, err := dbs.NewPostgres(config)
	if err != nil {
		logrus.Fatalf(err.Error())
		return
	}

    if err := postgres.Migrate(); err != nil {
		logrus.Fatalf(err.Error())
	}
	logrus.Info("migrate successfully")

	botApi, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		// log error into db
		
		logrus.Fatalf(err.Error())
		return
	}

	service := s.NewService(postgres)

	bot := b.NewBot(botApi, service)
	if err := bot.Start(); err != nil {
		service.LogError(err.Error()) // log to db
		logrus.Fatalf(err.Error())
		return
	}
    
}

