package services

import (
	"mime/multipart"
	"net/http"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/helper"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
)

type ProductService interface {
	AddProduct(product *domain.Product, image *multipart.FileHeader) (*domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	GetProductByStatus(status string) ([]domain.Product, error)
	UpdateProduct(product *domain.Product, image *multipart.FileHeader) (*domain.Product, error)
	DeleteProduct(id string) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func (s *productService) AddProduct(product *domain.Product, image *multipart.FileHeader) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	urlImage, err := helper.UploadImage(image)

	if err!= nil {
		return nil, &ErrorMessage{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	product.Image = urlImage
	repo := repositories.NewProductRepository(conn)

	product, err = repo.Insert(product)
	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to add product",
			Code:    http.StatusInternalServerError,
		}
	}

	return product, nil
}

func (s *productService) GetProductByID(id string) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductRepository(conn)

	product, err := repo.FindByID(id)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Product not found",
			Code:    http.StatusNotFound,
		}
	}

	return product, nil

}

func (s *productService) GetProductByStatus(status string) ([]domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductRepository(conn)

	products, err := repo.FindByStatus(status)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Product not found",
			Code:    http.StatusNotFound,
		}
	}

	return products, nil
}

func (s *productService) UpdateProduct(product *domain.Product, image *multipart.FileHeader) (*domain.Product, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductRepository(conn)

	_, err = repo.FindByID(product.ID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Product not found",
			Code:    http.StatusBadRequest,
		}
	}

	urlImage, err := helper.UploadImage(image)

	if err!= nil {
		return nil, &ErrorMessage{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}

	product.Image = urlImage

	product, err = repo.Update(product)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to update product to database",
			Code:    http.StatusNotFound,
		}
	}

	return product, nil
}

func (s *productService) DeleteProduct(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessage{
			Message : "Failed to connect to database",
			Code : http.StatusInternalServerError,
		}
	}

	repo := repositories.NewProductRepository(conn)

	_, err = repo.FindByID(id)

	if err!= nil {
		return &ErrorMessage{
			Message : "Product not found",
			Code : http.StatusBadRequest,
		}
	}

	err = repo.Delete(id)

	if err != nil {
		return &ErrorMessage{
			Message: "Failed to delete product from database",
			Code:    http.StatusInternalServerError,
		}
	}

	return err
}
