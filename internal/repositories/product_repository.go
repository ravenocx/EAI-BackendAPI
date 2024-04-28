package repositories

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(product *domain.Product) (*domain.Product, error)
	FindByID(id string) (*domain.Product, error)
	FindByStatus(status string) ([]domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	db = db.Debug()
	return &productRepository{db: db}
}

func (r *productRepository) Insert(product *domain.Product) (*domain.Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) FindByID(id string) (*domain.Product, error) {
	var product domain.Product
	if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) FindByStatus(status string) ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Where("on_development = ?", status).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Update(product *domain.Product) (*domain.Product, error) {
	var updatedProduct domain.Product

	if product.ID != "" {
		updatedProduct.ID = product.ID
	}

	if product.Product != "" {
		updatedProduct.Product = product.Product
	}

	if product.Description != "" {
		updatedProduct.Description = product.Description
	}

	if product.Category != "" {
		updatedProduct.Category = product.Category
	}

	if product.OnDevelopment != "" {
		updatedProduct.OnDevelopment = product.OnDevelopment
	}

	if product.Image != "" {
		updatedProduct.Image = product.Image
	}

	if err := r.db.Updates(&updatedProduct).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Delete(id string) error {
	var product domain.Product
	if err := r.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		return err
	}
	return nil
}
