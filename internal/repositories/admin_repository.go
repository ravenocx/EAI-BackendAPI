package repositories

import (
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"gorm.io/gorm"
)

type AdminRepository interface {
	InsertToDatabase(admin *domain.Admin) (*domain.Admin, error)
	FindByEmail(email string) (*domain.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *adminRepository {
	db = db.Debug()
	return &adminRepository{db: db}
}

func (r *adminRepository) InsertToDatabase(admin *domain.Admin) (*domain.Admin, error) {
	if err := r.db.Create(admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func (r *adminRepository) FindByEmail(email string) (*domain.Admin, error) {
	var admin domain.Admin
	if err := r.db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
