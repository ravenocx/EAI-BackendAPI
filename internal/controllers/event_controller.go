package controllers

import (
	"strings"
	"time"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *EventController) AddEvent(ctx *fiber.Ctx) (err error) {
	req := dto.EventRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	imgFile, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tempEventDate := date.In(loc)
	eventDate := tempEventDate.Add(-7 * time.Hour)

	event := &domain.Event{
		Name:      req.Name,
		Date:      eventDate,
		Detail:    req.Detail,
		Organizer: req.Organizer,
	}

	event, err = c.eventService.AddEvent(event, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	event, err = c.eventService.GetEventByID(event.ID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	eventResponse := dto.EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Date:      eventDate.String(),
		Detail:    event.Detail,
		Organizer: event.Organizer,
		Image:     event.Image,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success add event",
		"data":    eventResponse,
	})
}

func (c *EventController) GetAllEvents(ctx *fiber.Ctx) (err error) {
	events, err := c.eventService.GetAllEvents()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.EventResponse{}

	for _, event := range events {
		response = append(response, dto.EventResponse{
			ID:        event.ID,
			Name:      event.Name,
			Date:      event.Date.String(),
			Detail:    event.Detail,
			Organizer: event.Organizer,
			Image:     event.Image,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to get all events",
		"data":    response,
	})
}

func (c *EventController) UpdateEvent(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.EventRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	imgFile, _ := ctx.FormFile("image")
	if err != nil && (!strings.Contains(err.Error(), "there is no uploaded file")) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	event, err := c.eventService.GetEventByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	date, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	tempEventDate := date.In(loc)
	eventDate := tempEventDate.Add(-7 * time.Hour)

	event.Name = req.Name
	event.Date = eventDate
	event.Detail = req.Detail
	event.Organizer = req.Organizer

	_, err = c.eventService.UpdateEvent(event, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	event, _ = c.eventService.GetEventByID(id)

	eventResponse := dto.EventResponse{
		ID:        event.ID,
		Name:      event.Name,
		Date:      event.Date.String(),
		Detail:    event.Detail,
		Organizer: event.Organizer,
		Image:     event.Image,
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success update member data",
		"data":    eventResponse,
	})
}

func (c *EventController) DeleteEvent(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = c.eventService.DeleteEvent(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success deleted event",
	})
}
