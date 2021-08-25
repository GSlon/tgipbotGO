package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"

	"strconv"
	"fmt"
)

func (p *Postgres) CreateUser(id, chatid uint) error {
	user := m.User{
		UserID: id,
		ChatID: chatid,
	}

	result := db.Create(&user)
	return result.Error
}

func (p *Postgres) getUser(id uint) (m.User, error) {
	var user m.User
	result := db.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, &UserNotFoundError{}
	}

	return user, nil
}

func (p *Postgres) getAllUsers() ([]m.User, error) {
	var users []m.User
	result := db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	if result.RowsAffected == 0 {
		return users, &UserNotFoundError{}
	}

	return users, nil
}

func (p *Postgres) GetAllUsersChatID() ([]int, error) {
	users, err := getAllUsers()
	
	if err != nil {
		return []int{}, err
	}

	var chatsID []int
	for user, _ := range users {
		chatsID = append(chatsID, user.ChatID)
	} 
	
	return chatsID, nil
}

func (p *Postgres) GetAllUsersInfo() ([]string, error) {
	users, err := getAllUsers()
	
	if err != nil {
		return []int{}, err
	}

	var usersInfo []string
	for user, _ := range users {
		info := fmt.Sprintf("user_id: %d, chat_id: %d", user.UserID, user.chatsID)
		usersInfo = append(usersInfo, info)
	} 
	
	return usersInfo, nil
}

func (p *Postgres) GetUserInfo(id uint) (string, error) {
	user, err := getUser(id)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf("id: %d, user_id: %d, chat_id: %d, state: %s",
						user.ID, user.UserID, user.ChatID, user.State)

	return info, nil
}

// UserLog functions
func (p *Postgres) CreateUserLog(userid uint, ip, info string) error {
	user, err := getUser(userid)
	if err != nil {
		return "", err
	}

	userlog := m.UserLog{
		User: user,
		IP: ip,
		Info: info,
	}

	result := db.Create(userlog)
	return result.Error
}

// удаляем по полю id в бд
func (p *Postgres) DeleteUserLog(id uint) error {
	userlog := m.UserLog{ID: id}
	result := p.db.Delete(&userlog)
	return result.Error
}

func (p *Postgres) getUniqueUserLogs(userid uint) ([]m.UserLog, error) {
	var userlogs []m.UserLog
	result := db.Distinct("ip").Where("user_id = ?", userid).Find(&userlogs)
	if result.Error != nil {
		return user, result.Error
	}

	if result.RowsAffected == 0 {
		return user, &LogNotFoundError{}
	}
	
	return userlogs, nil
} 

func (p *Postgres) GetUserUniqueIPs(userid uint) ([]string, error) {
	userlogs, err := getUniqueUserLogs(userid)
	if err != nil {
		return []string{}, err
	}

	var res []string 
	for log, _ := range userlogs {
		res = append(res, log.IP)
	}

	return res
}

// map -> ip: info
func (p *Postgres) GetUserUniqueIPsExt(userid uint) (map[string]string, err) {
	userlogs, err := getUniqueUserLogs(userid)
	if err != nil {
		return []string{}, err
	}

	var res map[string]string 
	for log, _ := range userlogs {
		res[log.IP] = log.Info
	}

	return res
}
