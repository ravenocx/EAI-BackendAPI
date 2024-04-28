package routes

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/controllers"
	"github.com/ravenocx/EAI-BackendAPI/internal/middleware"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(router fiber.Router, eventService services.EventService) {
	eventController := controllers.NewEventController(eventService)

	event := router.Group("/events").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	// event := admin.Group("/events")
	event.Post("/add", eventController.AddEvent)
	event.Get("", eventController.GetAllEvents)
	event.Put("/:id", eventController.UpdateEvent)
	event.Delete("/:id", eventController.DeleteEvent)
}
