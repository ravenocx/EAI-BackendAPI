package repositories

import (
	"time"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"gorm.io/gorm"
)

type EventRepository interface {
	Insert(event *domain.Event) (*domain.Event, error)
	GetEventByID(id string) (*domain.Event, error)
	GetAllEvents() ([]domain.Event, error)
	Update(event *domain.Event) (*domain.Event, error)
	Delete(id string) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	db = db.Debug()
	return &eventRepository{db: db}
}

func (r *eventRepository) Insert(event *domain.Event) (*domain.Event, error) {
	if err := r.db.Create(event).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) GetEventByID(id string) (*domain.Event, error) {
	var event domain.Event
	if err := r.db.Where("id = ?", id).First(&event).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) GetAllEvents() ([]domain.Event, error) {
	var events []domain.Event
	if err := r.db.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *eventRepository) Update(event *domain.Event) (*domain.Event, error) {
	var updatedEvent domain.Event

	if event.ID != ""{
		updatedEvent.ID = event.ID
	}

	if event.Name != ""{
		updatedEvent.Name = event.Name
	}

	if event.Date != (time.Time{}){
		updatedEvent.Date = event.Date
	}

	if event.Detail != ""{
		updatedEvent.Detail = event.Detail
	}

	if event.Organizer != ""{
		updatedEvent.Organizer = event.Organizer
	}

	if event.Image != ""{
		updatedEvent.Image = event.Image
	}

	if err := r.db.Updates(&updatedEvent).Error; err != nil {
		return nil, err
	}
	return event, nil
}

func (r *eventRepository) Delete(id string) error {
	var event domain.Event
	if err := r.db.Where("id = ?",id).Delete(&event).Error; err != nil{
		return err
	}
	return nil
}