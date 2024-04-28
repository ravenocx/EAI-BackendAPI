package routes

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/controllers"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAdminRoutes(router fiber.Router, adminService services.AdminService) {
	adminController := controllers.NewAdminController(adminService)

	router.Post("/register", adminController.Register)
	router.Post("/login", adminController.Login)

}
