package controllers

import "github.com/ravenocx/EAI-BackendAPI/internal/services"

type EventController struct {
	eventService services.EventService
}

func NewEventController(eventService services.EventService) *EventController{
	return &EventController{eventService: eventService}
}