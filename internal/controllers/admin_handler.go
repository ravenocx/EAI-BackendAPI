package controllers

import "github.com/ravenocx/EAI-BackendAPI/internal/services"

type AdminController struct {
	adminService services.AdminService
}

func NewAdminController(adminService services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}