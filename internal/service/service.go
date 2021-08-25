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

func (s *Service) CreateAdmin(id uint) error {
	if err := db.CreateAdmin(id); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("admin %s created", id))
}

func (p *Postgres) DeleteAdmin(id uint) error {

}