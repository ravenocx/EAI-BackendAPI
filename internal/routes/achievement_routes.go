package routes

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/controllers"
	"github.com/ravenocx/EAI-BackendAPI/internal/middleware"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAchievementRoutes(router fiber.Router, achievementService services.AchievementService) {
	achievementController := controllers.NewAchievementController(achievementService)

	achievement := router.Group("/achievements").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	// achievement := admin.Group("/achievements")
	achievement.Post("/add", achievementController.AddAchievement)
	achievement.Get("", achievementController.GetAllAchievements)
	achievement.Put("/:id", achievementController.UpdateAchievement)
	achievement.Delete("/:id", achievementController.DeleteAchievement)
}
