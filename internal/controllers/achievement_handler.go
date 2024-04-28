package controllers

import "github.com/ravenocx/EAI-BackendAPI/internal/services"

type AchievementController struct {
	achievementService services.AchievementService
}

func NewAchievementController(achievementService services.AchievementService) *AchievementController{
	return &AchievementController{achievementService: achievementService}
}