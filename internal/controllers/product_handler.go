package controllers

import "github.com/ravenocx/EAI-BackendAPI/internal/services"

type ProductController struct {
	productService services.ProductService
}

func NewProductController (productService services.ProductService) *ProductController{
	return &ProductController{productService: productService}
}