package main

import (
    "github.com/spf13/viper"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"

	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"

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
    }

    if err := godotenv.Load(); err != nil {
		logrus.Fatalf(err.Error())
	}

    // вызвать Postgres.Migrate
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
	}

	if err := postgres.Migrate(); err != nil {
		logrus.Fatalf(err.Error())
	}
	logrus.Info("migrate successfully")
}
