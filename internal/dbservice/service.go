package dbservice

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
    "github.com/sirupsen/logrus"

	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
	"fmt"
)

type PostgresConfig struct {
    Host string
    Port string
    Username string
    Password string
    DBName string
    SSLMode string
}

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(cfg PostgresConfig) (*Postgres, error) {
	var config string
    if cfg.Password == "" {
        config = fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.SSLMode)
    } else {
        config = fmt.Sprintf(
        "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)
    }

    logrus.Info(config)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return &Postgres{}, err
	}

	logrus.Info("connection to postgres establish")

	return &Postgres{db: db}, nil
}

func (p *Postgres) Migrate() error {
	err := p.db.AutoMigrate(&m.User{}, &m.UserLog{}, &m.Admin{}, &m.ErrorLog{})
    return err
}



