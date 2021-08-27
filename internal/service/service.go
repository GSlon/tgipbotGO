package service

import (
	"github.com/sirupsen/logrus"
	
	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"
)

// прослойка между ботом и бд (логгирование и обработка ошибок)
type Service struct {
	db *dbs.Postgres	// для взаимодействия с бд
}

func NewService(db *dbs.Postgres) *Service {
	return &Service{db: db}	
}

func (s *Service) LogError(info string) error {
	if err := s.db.LogError(info); err != nil {
		return err
	}

	logrus.Info("error log to db")
	return nil
}
