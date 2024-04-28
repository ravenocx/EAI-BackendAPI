package controllers

import (
	"net/http"
	"strings"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *AchievementController) AddAchievement(ctx *fiber.Ctx) (err error) {
	req := dto.AchievementRequest{}

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

	achievement := &domain.Achievement{
		MemberID:    req.MemberID,
		Achievement: req.Achievement,
	}

	achievement, err = c.achievementService.AddAchievement(achievement, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achievement, err = c.achievementService.GetAchievementByID(achievement.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achievementResponse := dto.AchievementResponse{
		ID:          achievement.ID,
		Name:        achievement.Member.Name,
		Image:       achievement.Image,
		Achievement: achievement.Achievement,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success add Achievement",
		"data":    achievementResponse,
	})
}

func (c *AchievementController) UpdateAchievement(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.AchievementUpdateRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	imgFile, err := ctx.FormFile("image")
	if err != nil && (!strings.Contains(err.Error(), "there is no uploaded file")) {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achievement, err := c.achievementService.GetAchievementByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achievement.MemberID = req.MemberID
	achievement.Achievement = req.Achievement

	_, err = c.achievementService.UpdateAchievement(achievement, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	achievement, _ = c.achievementService.GetAchievementByID(id)

	response := dto.AchievementResponse{
		ID:          achievement.ID,
		Name:        achievement.Member.Name,
		Image:       achievement.Image,
		Achievement: achievement.Achievement,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success update achievement",
		"data":    response,
	})
}

func (c *AchievementController) DeleteAchievement(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = c.achievementService.DeleteAchievement(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success deleted achievement",
	})
}

func (c *AchievementController) GetAllAchievements(ctx *fiber.Ctx) (err error) {
	achievements, err := c.achievementService.GetAllAchievements()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.AchievementResponse{}

	for _, achievement := range achievements {
		response = append(response, dto.AchievementResponse{
			ID:          achievement.ID,
			Name:        achievement.Member.Name,
			Image:       achievement.Image,
			Achievement: achievement.Achievement,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to get all achievements",
		"data":    response,
	})
}
