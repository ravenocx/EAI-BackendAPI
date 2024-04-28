package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	ID        string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	AdminID   string    `json:"admin_id"`
	Token     string    `json:"token"`
	IssuedAt  time.Time `json:"iat"`
	ExpiredAt time.Time `json:"eat"`
	Admin     Admin     `json:"admin" gorm:"foreignKey:AdminID;references:ID"`
}

func (t *Token) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.NewString()
	return
}
