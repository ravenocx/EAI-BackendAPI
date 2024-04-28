package services

import (
	"net/http"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	Register(admin *domain.Admin) (*domain.Admin, error)
	Login(email, password string) (*domain.Admin, error)
}

type adminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService(adminRepo repositories.AdminRepository) *adminService {
	return &adminService{adminRepo: adminRepo}
}

func (s *adminService) Register(admin *domain.Admin) (*domain.Admin, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAdminRepository(conn)

	// check if email already exist
	_, err = repo.FindByEmail(admin.Email)

	if err == nil {
		return nil, &ErrorMessage{
			Message: "Email already exist",
			Code:    http.StatusBadRequest,
		}
	}

	// hash admin password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to hash password",
			Code:    500,
		}
	}

	admin.Password = string(hashedPassword)

	// insert admin to database
	admin, err = repo.InsertToDatabase(admin)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to insert admin to database",
			Code:    500,
		}
	}

	return admin, nil
}

func (s *adminService) Login(email, password string) (*domain.Admin, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAdminRepository(conn)

	admin, err := repo.FindByEmail(email)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Invalid email or password",
			Code:    http.StatusNotFound,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Invalid email or password",
			Code:    http.StatusUnauthorized,
		}
	}

	return admin, nil
}
