package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func (c *ProductController) AddProduct(ctx *fiber.Ctx) (err error) {
	req := dto.ProductRequest{}

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

	product := &domain.Product{
		Product:       req.Product,
		Description:   req.Description,
		Category:      req.Category,
		OnDevelopment: req.OnDevelopment,
	}

	product, err = c.productService.AddProduct(product, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, err = c.productService.GetProductByID(product.ID)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	productResponse := dto.ProductResponse{
		ID:            product.ID,
		Product:       product.Product,
		Description:   product.Description,
		Category:      product.Category,
		OnDevelopment: product.OnDevelopment,
		Image:         product.Image,
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success add Product",
		"data":    productResponse,
	})
}

func (c *ProductController) GetProductByStatus(ctx *fiber.Ctx) (err error) {
	status := strings.ToLower(ctx.Query("on_development"))

	if status == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "query product status development is required",
		})
	}

	response := []dto.ProductResponse{}

	if status != "" {
		_, err := strconv.ParseBool(status)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		products, err := c.productService.GetProductByStatus(status)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		for _, product := range products {
			response = append(response, dto.ProductResponse{
				ID:            product.ID,
				Product:       product.Product,
				Description:   product.Description,
				Category:      product.Category,
				OnDevelopment: product.OnDevelopment,
				Image:         product.Image,
			})
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success get products",
		"data":    response,
	})
}

func (c *ProductController) UpdateProduct(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	req := dto.ProductUpdateRequest{}

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

	product, err := c.productService.GetProductByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product.Product = req.Product
	product.Description = req.Description
	product.Category = req.Category
	product.OnDevelopment = req.OnDevelopment

	_, err = c.productService.UpdateProduct(product, imgFile)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	product, _ = c.productService.GetProductByID(id)

	response := dto.ProductResponse{
		ID:            product.ID,
		Product:       product.Product,
		Description:   product.Description,
		Category:      product.Category,
		OnDevelopment: product.OnDevelopment,
		Image:         product.Image,
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success update product data",
		"data":    response,
	})
}

func (c *ProductController) DeleteProduct(ctx *fiber.Ctx) (err error) {
	id := ctx.Params("id")

	err = c.productService.DeleteProduct(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success deleted product",
	})
}
