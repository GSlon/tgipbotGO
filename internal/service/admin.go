package service

import (
	"github.com/sirupsen/logrus"
	
	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"	// custom errors import
	"errors"
	"fmt"
)

func (s *Service) CreateAdmin(id int, state string) error {
	if err := s.db.CreateAdmin(id, state); err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.AdminAlreadyExistsError{}) {
			return errors.New("admin already exists")
		}

		return errors.New("internal error")
	}

	logrus.Info(fmt.Sprintf("admin %d created", id))
	return nil
}

func (s *Service) DeleteAdmin(id int) error {
	if err := s.db.DeleteAdmin(id); err != nil {
		s.LogError(err.Error())
		return errors.New("internal error")
	}

	logrus.Info(fmt.Sprintf("admin %d deleted", id))
	return nil
}

func (s *Service) CheckAdminExists(id int) (bool, error) {
	exists, err := s.db.CheckAdminExists(id)
	if err != nil {
		s.LogError(err.Error())
		return false, errors.New("internal error")
	}

	return exists, nil
}

func (s *Service) GetAdminState(id int) (string, error) { 
	state, err := s.db.GetAdminState(id)
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.AdminNotFoundError{}) {
			return "", errors.New("admin not found")
		}

		return "", errors.New("internal error")
	}

	return state, nil
}

func (s *Service) SetAdminState(id int, state string) error {
	if err := s.db.SetAdminState(id, state); err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.AdminNotFoundError{}) {
			return errors.New("admin not found")
		}

		return errors.New("internal error")
	}

	logrus.Info(fmt.Sprintf("admin %d -> changed state %s", id, state))
	return nil
}