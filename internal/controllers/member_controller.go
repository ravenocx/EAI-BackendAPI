package controllers

import (
	"net/http"
	"strings"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *MemberController) AddMember(ctx *fiber.Ctx) (err error) {
	req := dto.MemberRequest{}

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

	member := &domain.Member{
		Name:   req.Name,
		Role:   req.Role,
		Status: req.Status,
	}

	member, err = c.memberService.AddMember(member, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	member, err = c.memberService.GetMemberByID(member.ID)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	memberResponse := dto.MemberResponse{
		ID:     member.ID,
		Name:   member.Name,
		Role:   member.Role,
		Status: member.Status,
		Image:  member.Image,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success add member",
		"data":    memberResponse,
	})
}

func (c *MemberController) UpdateMember(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.MemberRequest{}

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

	member, err := c.memberService.GetMemberByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	member.Name = req.Name
	member.Role = req.Role
	member.Status = req.Status

	_, err = c.memberService.UpdateMember(member, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	member, _ = c.memberService.GetMemberByID(id)

	response := dto.MemberResponse{
		ID:     member.ID,
		Name:   member.Name,
		Role:   member.Role,
		Status: member.Status,
		Image:  member.Image,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success update member data",
		"data":    response,
	})
}

func (c *MemberController) DeleteMember(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = c.memberService.DeleteMember(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success deleted member",
	})
}

func (c *MemberController) GetAllMembers(ctx *fiber.Ctx) (err error) {
	members, err := c.memberService.GetAllMembers()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := []dto.MemberResponse{}

	for _, member := range members {
		response = append(response, dto.MemberResponse{
			ID:     member.ID,
			Name:   member.Name,
			Role:   member.Role,
			Status: member.Status,
			Image:  member.Image,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to get all members",
		"data":    response,
	})
}

func (c *MemberController) GetMembersByRole(ctx *fiber.Ctx) (err error) {
	role := ctx.Query("role")

	if role == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query string member role is required",
		})
	}

	response := []dto.MemberResponse{}

	if role != "" {
		members, err := c.memberService.GetMemberByRole(role)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		for _, member := range members {
			response = append(response, dto.MemberResponse{
				ID:     member.ID,
				Name:   member.Name,
				Role:   member.Role,
				Status: member.Status,
				Image:  member.Image,
			})
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get members",
		"data":    response,
	})
}
