package dbservice

import m "github.com/GSlon/tgipbotGO/internal/dbservice/models"

func (p *Postgres) AddAdmin(id uint) error {
	admin := m.Admin{UserID: id}
	result := p.db.Create(&admin)
	return result.Error
}

func (p *Postgres) RemoveAdmin(id uint) error {
	admin := m.Admin{UserID: id}
	result := p.db.Delete(&admin)
	return result.Error
}

func (p *Postgres) CheckAdminExists(id uint) (bool, error) {
	var admin m.Admin
	result := p.db.First(&admin, id)
	if result.Error != nil {
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		return false, nil
	} 
	return true, nil
}

func (p *Postgres) getAdmin(id uint) (m.Admin, error) {
	var admin m.Admin
	result := p.db.First(&admin, id)
	if result.Error != nil {
		return admin, result.Error
	}

	if result.RowsAffected == 0 {
		return admin, &AdminNotFoundError{}
	}

	return admin, nil
}

func (p *Postgres) GetAdminState(id uint) (string, error) {
	admin, err := p.getAdmin(id)	
	if err != nil {
		return "", err
	}

	return admin.State, nil
}



func (p *Postgres) SetAdminState(id uint, state string) error {
	admin, err := p.getAdmin(id)	
	if err != nil {
		return err
	}

	admin.State = state
	result := db.Save(&admin)
	
	return result.Error
}

