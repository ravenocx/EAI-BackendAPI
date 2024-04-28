package routes

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/controllers"
	"github.com/ravenocx/EAI-BackendAPI/internal/middleware"
	"github.com/ravenocx/EAI-BackendAPI/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupMemberRoutes(router fiber.Router, memberService services.MemberService){
	memberController := controllers.NewMemberController(memberService)

	members := router.Group("/members").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	// member := admin.Group("/members")
	members.Post("/add",memberController.AddMember )
	members.Get("",memberController.GetAllMembers )
	members.Put("/:id",memberController.UpdateMember)
	members.Delete("/:id", memberController.DeleteMember)

	member := router.Group("/member").Use(middleware.AdminAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))
	member.Get("",memberController.GetMembersByRole)

}