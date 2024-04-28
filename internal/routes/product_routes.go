package routes

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/controllers"
	"github.com/ravenocx/EAI-BackendAPI/internal/middleware"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(router fiber.Router, productService services.ProductService){
	productController := controllers.NewProductController(productService)

	products := router.Group("/products").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	// product := admin.Group("products")
	products.Post("/add", productController.AddProduct)
	products.Put("/:id", productController.UpdateProduct)
	products.Delete("/:id", productController.DeleteProduct)

	product := router.Group("/product").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	product.Get("", productController.GetProductByStatus)
}