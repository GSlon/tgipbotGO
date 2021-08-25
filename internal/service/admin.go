package service

import (
	"github.com/sirupsen/logrus"
)

func (s *Service) CreateAdmin(id uint, state string) error {
	if err := s.db.CreateAdmin(id, state); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("admin %d created", id))
}

func (s *Service) DeleteAdmin(id uint) error {
	if err := s.db.DeleteAdmin(id); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("admin %d deleted", id))
}

func (s *Service) CheckAdminExists(id uint) (bool, error) {
	exists, err := s.db.CheckAdminExists(id)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *Service) GetAdminState(id uint) (string, error) { 
	state, err := s.db.GetAdminState(id)
	if err != nil {
		return "", err
	}

	return state, nil
}

func (s *Service) SetAdminState(id uint, state string) error {
	if err := s.db.SetAdminState(id, state); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("admin %d -> changed state %s", id, state))
}