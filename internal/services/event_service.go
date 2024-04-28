package services

import (
	"mime/multipart"
	"net/http"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/helper"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
)

type EventService interface {
	AddEvent(event *domain.Event, image *multipart.FileHeader) (*domain.Event, error)
	GetEventByID(id string) (*domain.Event, error)
	GetAllEvents() ([]domain.Event, error)
	UpdateEvent(event *domain.Event, image *multipart.FileHeader) (*domain.Event, error)
	DeleteEvent(id string) error
}

type eventService struct {
	eventRepository repositories.EventRepository
}

func NewEventService(eventRepository repositories.EventRepository) *eventService {
	return &eventService{eventRepository: eventRepository}
}

func (s *eventService) AddEvent(event *domain.Event, image *multipart.FileHeader) (*domain.Event, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	urlImage, err := helper.UploadImage(image)

	if err!= nil {
		return nil, &ErrorMessage{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	event.Image = urlImage

	repo := repositories.NewEventRepository(conn)

	event, err = repo.Insert(event)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to add event",
			Code:    http.StatusInternalServerError,
		}
	}

	return event, nil
}

func (s *eventService) GetEventByID(id string) (*domain.Event, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewEventRepository(conn)

	event, err := repo.GetEventByID(id)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Event not found",
			Code:    http.StatusNotFound,
		}
	}

	return event, nil
}

func (s *eventService) GetAllEvents() ([]domain.Event, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewEventRepository(conn)

	events, err := repo.GetAllEvents()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to retrieve events data",
			Code:    http.StatusInternalServerError,
		}
	}
	return events, nil
}

func (s *eventService) UpdateEvent(event *domain.Event, image *multipart.FileHeader) (*domain.Event, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewEventRepository(conn)

	_, err = repo.GetEventByID(event.ID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Event not found",
			Code:    http.StatusBadRequest,
		}
	}

	urlImage, err := helper.UploadImage(image)

	if err!= nil {
		return nil, &ErrorMessage{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	event.Image = urlImage

	event, err = repo.Update(event)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to update event to database",
			Code:    http.StatusInternalServerError,
		}
	}

	return event, nil
}

func (s *eventService) DeleteEvent(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessage{
			Message : "Failed to connect to database",
			Code : http.StatusInternalServerError,
		}
	}

	repo := repositories.NewEventRepository(conn)

	_, err = repo.GetEventByID(id)

	if err!= nil {
		return &ErrorMessage{
			Message : "Event not found",
			Code : http.StatusBadRequest,
		}
	}

	err = repo.Delete(id)

	if err != nil {
		return &ErrorMessage{
			Message: "Failed to delete event from database",
			Code:    http.StatusInternalServerError,
		}
	}

	return err
}