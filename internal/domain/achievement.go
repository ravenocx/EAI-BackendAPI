package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Achievement struct {
	ID          string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	MemberID    string `json:"member_id"`
	Member      Member `json:"name" gorm:"foreignKey:MemberID;references:ID"`
	Image       string `json:"image"`
	Achievement string `json:"achievement"`
}

func (a *Achievement) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return
}
