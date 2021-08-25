package service

import (
	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"
	"github.com/sirupsen/logrus"
)

// прослойка между ботом и бд
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
}
