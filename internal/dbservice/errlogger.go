package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
)

func (p *Postgres) LogError(info string) error {
	log := m.ErrorLog{Info: info}
	result := p.db.Create(log)
	return result.Error	
}