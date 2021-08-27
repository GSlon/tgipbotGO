package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
	"errors"
)

func (p *Postgres) CreateAdmin(userid int, state string) error {
	exists, _ := p.CheckAdminExists(userid)	// ignore check errors
	if exists {
		return &AdminAlreadyExistsError{}
	}

	admin := m.Admin{
		UserID: userid,
		State: state,
	}
	
	result := p.db.Create(&admin)
	return result.Error
}

func (p *Postgres) DeleteAdmin(userid int) error {
	var admin m.Admin
	result := p.db.Where("user_id=?", userid).First(&admin)
	if result.Error != nil {
		return result.Error
	}
	
	res := p.db.Delete(&admin)
	return res.Error
}

func (p *Postgres) CheckAdminExists(userid int) (bool, error) {
	_, err := p.getAdmin(userid)
	if err != nil {
		if errors.Is(err, &AdminNotFoundError{}) { // для этой функции такое - не ошибка
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (p *Postgres) getAdmin(userid int) (m.Admin, error) {
	var admin m.Admin
	result := p.db.Where("user_id=?", userid).Find(&admin)
	if result.Error != nil {
		return admin, result.Error
	}

	if result.RowsAffected == 0 {
		return admin, &AdminNotFoundError{}
	}

	return admin, nil
}

func (p *Postgres) GetAdminState(userid int) (string, error) {
	admin, err := p.getAdmin(userid)	
	if err != nil {
		return "", err
	}

	return admin.State, nil
}

func (p *Postgres) SetAdminState(userid int, state string) error {
	admin, err := p.getAdmin(userid)	
	if err != nil {
		return err
	}

	admin.State = state
	result := p.db.Save(&admin)
	
	return result.Error
}

