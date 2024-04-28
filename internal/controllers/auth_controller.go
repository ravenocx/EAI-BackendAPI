package controllers

import (
	"net/http"
	"time"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
	"github.com/ravenocx/EAI-BackendAPI/internal/helper"
	"github.com/gofiber/fiber/v2"
)

func (c *AdminController) Register(ctx *fiber.Ctx) error {
	req := dto.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// insert admin to database
	admin := &domain.Admin{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// save admin to database
	admin, err := c.adminService.Register(admin)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// return Response
	response := dto.RegisterResponse{
		ID:    admin.ID,
		Name:  admin.Name,
		Email: admin.Email,
	}

	// return api
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Admin created successfully",
		"data":    response,
	})
}

func (c *AdminController) Login(ctx *fiber.Ctx) error {
	req := dto.LoginRequest{}
	
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	admin,err := c.adminService.Login(req.Email, req.Password)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token,err := helper.GenerateAdminToken(admin)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	response := dto.LoginResponse{
		ID : admin.ID,
		Name : admin.Name,
		Email : admin.Email,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "User logged in successfully",
		"data":    response,
		"token":   token,
	})
}
