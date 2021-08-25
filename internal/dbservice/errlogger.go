package dbservice

import (
	m "github.com/GSlon/tgipbotGO/internal/dbservice/models"
)

func (p *Postgres) LogError(info string) error {
	log := m.ErrorLog{info: info}
	result := db.Create(log)
	return result.Error	
}