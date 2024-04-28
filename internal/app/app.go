package app

import (
	"os"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
	"github.com/ravenocx/EAI-BackendAPI/internal/routes"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartApplication() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	app.Get("/ping",func(c *fiber.Ctx) error {
		return c.SendString("PING!")
	})
	// connect to database
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	// init routes
	api := app.Group("/api")

	// init repositories
	adminRepo := repositories.NewAdminRepository(db)

	// init service
	adminService := services.NewAdminService(adminRepo)

	routes.SetupAdminRoutes(api, adminService)

	// init member repo and service
	memberRepository := repositories.NewMemberRepository(db)
	memberService := services.NewMemberService(memberRepository)
	routes.SetupMemberRoutes(api, memberService)

	// init achievement repo and service
	achievementRepository := repositories.NewAchievementRepository(db)
	achievementService := services.NewAchievementService(achievementRepository)
	routes.SetupAchievementRoutes(api, achievementService)

	// init product repo and service
	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	routes.SetupProductRoutes(api, productService)

	// init event repo and service
	eventRepository := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepository)
	routes.SetupEventRoutes(api, eventService)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
