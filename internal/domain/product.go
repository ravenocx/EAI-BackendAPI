package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID            string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Product       string `json:"product"`
	Description   string `json:"description"`
	Category      string `json:"category"`
	OnDevelopment string   `json:"on_development"`
	Image         string `json:"image"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}
