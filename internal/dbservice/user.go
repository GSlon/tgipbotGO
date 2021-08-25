package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
)

func (p *Postgres) AddUser(id, chatid uint) error {
	user := m.User{
		UserID: id,
		ChatID: chatid,
	}

	result := db.Create(&user)
	return result.Error
}

func (p *Postgres) GetAllUsersChatID() ([]int, error) {
	var users []m.User
	result := db.Select("chat_id").Find(&users)
	
	if result.Error != nil {
		return []int{}, result.Error
	}

	var chatsID []int
	for user, _ := range users {
		chatsID = append(chatsID, user.ChatID)
	} 
	
	return chatsID, nil
}

func (p *Postgres) getUsers() ([]m.User, error) {

}

func (p *Postgres) GetAllUsersInfo() []string {

}

func (p *Postgres) GetUserInfo(id uint) string {

}

// UserLog functions

