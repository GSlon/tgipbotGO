package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"

	"fmt"
	"errors"
	_ "strconv"
)

func (p *Postgres) CreateUser(userid int, chatid int64, state string) error {
	exists, _ := p.CheckUserExists(userid)	// ignore check errors
	if exists {
		return &UserAlreadyExistsError{}
	}

	user := m.User{
		UserID: userid,
		ChatID: chatid,
		State: state,
	}

	result := p.db.Create(&user)
	return result.Error
}

func (p *Postgres) getUser(userid int) (m.User, error) {
	var user m.User
	result := p.db.Where("user_id=?", userid).Find(&user)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, &UserNotFoundError{}
	}

	return user, nil
}

func (p *Postgres) CheckUserExists(userid int) (bool, error) {
	_, err := p.getUser(userid)
	if err != nil {
		if errors.Is(err, &UserNotFoundError{}) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *Postgres) getAllUsers() ([]m.User, error) {
	var users []m.User
	result := p.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	if result.RowsAffected == 0 {
		return users, &UserNotFoundError{}
	}

	return users, nil
}

func (p *Postgres) GetAllUsersChatID() ([]int64, error) {
	users, err := p.getAllUsers()
	
	if err != nil {
		return []int64{}, err
	}

	var chatsID []int64
	for _, user := range users {
		chatsID = append(chatsID, user.ChatID)
	} 
	
	return chatsID, nil
}

func (p *Postgres) GetAllUsersInfo() ([]string, error) {
	users, err := p.getAllUsers()
	
	if err != nil {
		return []string{}, err
	}

	var usersInfo []string
	for _, user := range users {
		info := fmt.Sprintf("user_id: %d, chat_id: %d", user.UserID, user.ChatID)
		usersInfo = append(usersInfo, info)
	} 
	
	return usersInfo, nil
}

func (p *Postgres) GetUserInfo(id int) (string, error) {
	user, err := p.getUser(id)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("id: %d, user_id: %d, chat_id: %d, state: %s",
						user.ID, user.UserID, user.ChatID, user.State)

	return info, nil
}

func (p *Postgres) GetUserState(id int) (string, error) {
	user, err := p.getUser(id)	
	if err != nil {
		return "", err
	}

	return user.State, nil
}

func (p *Postgres) SetUserState(userid int, state string) error {
	user, err := p.getUser(userid)	
	if err != nil {
		return err
	}

	user.State = state
	result := p.db.Save(&user)
	
	return result.Error
}

// UserLog functions
func (p *Postgres) CreateUserLog(userid int, ip, info string) error {
	user, err := p.getUser(userid)
	if err != nil {
		return err
	}

	userlog := m.UserLog{
		User: user,
		IP: ip,
		Info: info,
	}

	result := p.db.Create(&userlog)
	return result.Error
}

// удаляем по id юзера и ip
func (p *Postgres) DeleteUserLog(userid int, ip string) error {
	user, err := p.getUser(userid)
	if err != nil {
		return err
	}

	var userlog m.UserLog 
	result := p.db.Where("user_id = ? AND ip = ?", fmt.Sprint(user.ID), ip).First(&userlog)
	if result.Error != nil {
		return result.Error
	}
	
	res := p.db.Delete(&userlog)
	return res.Error
}

func (p *Postgres) getUniqueUserLogs(userid int) ([]m.UserLog, error) {
	user, err := p.getUser(userid)
	if err != nil {
		return []m.UserLog{}, err
	}

	var userlogs []m.UserLog
	result := p.db.Distinct("ip").Where("user_id = ?", user.ID).Find(&userlogs)
	if result.Error != nil {
		return userlogs, result.Error
	}

	if result.RowsAffected == 0 {
		return userlogs, &LogNotFoundError{}
	}
	
	return userlogs, nil
} 

// only ip's
func (p *Postgres) GetUserUniqueIPs(userid int) ([]string, error) {
	userlogs, err := p.getUniqueUserLogs(userid)
	if err != nil {
		return []string{}, err
	}

	var res []string 
	for _, log := range userlogs {
		res = append(res, log.IP)
	}

	return res, nil
}

// map -> ip: info
func (p *Postgres) GetUserUniqueIPsExt(userid int) (map[string]string, error) {
	userlogs, err := p.getUniqueUserLogs(userid)
	if err != nil {
		return map[string]string{}, err
	}

	var res map[string]string 
	for _, log := range userlogs {
		res[log.IP] = log.Info
	}

	return res, nil
}
