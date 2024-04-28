package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	ID     string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name   string `json:"name"`
	Role   string `json:"role"` // divisi
	Status string `json:"status"`
	Image  string `json:"image"`
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.NewString()
	return
}
