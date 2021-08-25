package service

import (
	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"
)

type Service struct {
	db *dbs.Postgres	// для взаимодействия с бд
}

func NewService(db *dbs.Postgres) *Service {
	return &Service{db: db}	
}

func (s *Service) AddAdmin() {
	
}