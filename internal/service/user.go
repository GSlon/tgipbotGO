package service

import (
	"github.com/sirupsen/logrus"
)

func (s *Service) CreateUser(id, chatid uint, state string) error {
	if err := s.db.CreateUser(id, chatid, state); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("user %d created", id))
}

func (s *Service) GetAllUsersChatID() ([]uint, error) {
	chatsID, err := s.db.GetAllUsersChatID()
	if err != nil {
		return []uint{}, err
	}
	
	return chatsID, nil
}

func (s *Service) GetAllUsersInfo() ([]string, error) {
	info, err := s.db.GetAllUsersInfo()
	if err != nil {
		return []string{}, err
	}

	return info, nil
}

func (s *Service) GetUserInfo(id uint) (string, error) {
	info, err := s.db.GetUserInfo(id)
	if err != nil {
		return "", err
	}

	return info, nil
}

func (s *Service) GetUserState(id uint) (string, error) {
	state, err := s.db.GetUserState(id)
	if err != nil {
		return "", err
	}

	return state, nil
}

func (s *Service) SetUserState(id uint, state string) error {
	if err := s.db.SetUserState(id, state); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("user %d -> changed state %s", id, state))
}


// UserLog functions
func (s *Service) CreateUserLog(userid uint, ip, info string) error {
	if err := s.db.CreateUserLog(userid, ip, info); err != nil {
		return err
	}

	logrus.Info(fmt.Sprintf("user %d log -> created", userid))
}

func (s *Service) DeleteUserLog(id uint, ip string) error {
	if err := s.db.DeleteUserLog(uint, ip); err != nil {
		return err
	} 

	logrus.Info(fmt.Sprintf("user %d log -> deleted", userid))
}

func (s *Service) GetUserUniqueIPs(userid uint) ([]string, error) {
	ip, err := s.db.GetUserUniqueIPs(userid)
	if err != nil {
		return []string{}, err
	}

	return ip, nil
}

func (s *Service) GetUserUniqueIPsExt(userid uint) (map[string]string, error) {
	logs, err := s.db.GetUserUniqueIPsExt(userid)
	if err != nil {
		return map[string]string, err
	}

	return logs, nil	
}
