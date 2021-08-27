package service

import (
	"github.com/sirupsen/logrus"

	dbs "github.com/GSlon/tgipbotGO/internal/dbservice"	// custom errors import
	"errors"
	"fmt"
)

func (s *Service) CreateUser(id int, chatid int64, state string) error {
	if err := s.db.CreateUser(id, chatid, state); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("user %d created", id))
	return nil
}

func (s *Service) GetAllUsersChatID() ([]int64, error) {
	chatsID, err := s.db.GetAllUsersChatID()
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return []int64{}, errors.New("user not found")
		}

		return []int64{}, errors.New("internal error")
	}
	
	return chatsID, nil
}

func (s *Service) GetAllUsersInfo() ([]string, error) {
	info, err := s.db.GetAllUsersInfo()
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return []string{}, errors.New("user not found")
		}

		return []string{}, errors.New("internal error")
	}

	return info, nil
}

func (s *Service) GetUserInfo(id int) (string, error) {
	info, err := s.db.GetUserInfo(id)
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return "", errors.New("user not found")
		}

		return "", errors.New("internal error")	
	}

	return info, nil
}

func (s *Service) GetUserState(id int) (string, error) {
	state, err := s.db.GetUserState(id)
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return "", errors.New("user not found")
		}

		return "", errors.New("internal error")
	}

	return state, nil
}

func (s *Service) SetUserState(id int, state string) error {
	if err := s.db.SetUserState(id, state); err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return errors.New("user not found")
		}

		return errors.New("internal error")
	}

	logrus.Info(fmt.Sprintf("user %d -> changed state %s", id, state))
	return nil
}


// UserLog functions
func (s *Service) CreateUserLog(userid int, ip, info string) error {
	if err := s.db.CreateUserLog(userid, ip, info); err != nil {
		s.LogError(err.Error())
		return errors.New("internal error")
	}

	logrus.Info(fmt.Sprintf("user %d log -> created", userid))
	return nil
}

func (s *Service) DeleteUserLog(userid int, ip string) error {
	if err := s.db.DeleteUserLog(userid, ip); err != nil {
		return err
	} 

	logrus.Info(fmt.Sprintf("user %d log -> deleted", userid))
	return nil
}

func (s *Service) GetUserUniqueIPs(userid int) ([]string, error) {
	ip, err := s.db.GetUserUniqueIPs(userid)
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return []string{}, errors.New("user not found")
		}

		if errors.Is(err, &dbs.LogNotFoundError{}) {
			return []string{}, errors.New("log not found")
		}

		return []string{}, errors.New("internal error")
	}

	return ip, nil
}

func (s *Service) GetUserUniqueIPsExt(userid int) (map[string]string, error) {
	logs, err := s.db.GetUserUniqueIPsExt(userid)
	if err != nil {
		s.LogError(err.Error())

		if errors.Is(err, &dbs.UserNotFoundError{}) {
			return map[string]string{}, errors.New("user not found")
		}

		if errors.Is(err, &dbs.LogNotFoundError{}) {
			return map[string]string{}, errors.New("log not found")
		}

		return map[string]string{}, errors.New("internal error")
	}

	return logs, nil	
}
