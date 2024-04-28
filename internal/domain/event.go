package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID        string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Detail    string    `json:"detail"`
	Organizer string    `json:"organizer"`
	Image     string    `json:"image"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.NewString()
	return
}
