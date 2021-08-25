package dbservice

import m "github.com/GSlon/tgipbotGO/internal/dbservice/models"

func (p *Postgres) CreateAdmin(userid uint, state string) error {
	admin := m.Admin{
		UserID: userid,
		State: state,
	}
	
	result := p.db.Create(&admin)
	return result.Error
}

func (p *Postgres) DeleteAdmin(userid uint) error {
	admin := m.Admin{UserID: userid}
	result := p.db.Delete(&admin)
	return result.Error
}

func (p *Postgres) CheckAdminExists(userid uint) (bool, error) {
	_, err := p.getAdmin(userid)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *Postgres) getAdmin(userid uint) (m.Admin, error) {
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

func (p *Postgres) GetAdminState(userid uint) (string, error) {
	admin, err := p.getAdmin(userid)	
	if err != nil {
		return "", err
	}

	return admin.State, nil
}

func (p *Postgres) SetAdminState(userid uint, state string) error {
	admin, err := p.getAdmin(userid)	
	if err != nil {
		return err
	}

	admin.State = state
	result := p.db.Save(&admin)
	
	return result.Error
}

