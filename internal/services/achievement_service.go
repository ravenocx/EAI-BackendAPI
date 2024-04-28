package services

import (
	"mime/multipart"
	"net/http"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/helper"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
)

type AchievementService interface {
	AddAchievement(achievement *domain.Achievement, image *multipart.FileHeader) (*domain.Achievement, error)
	UpdateAchievement(achievement *domain.Achievement, image *multipart.FileHeader) (*domain.Achievement, error)
	DeleteAchievement(id string) error
	GetAchievementByID(id string) (*domain.Achievement, error)
	GetAllAchievements() ([]domain.Achievement, error)
}

type achievementService struct {
	achievementRepository repositories.AchievementRepository
}

func NewAchievementService(achievementRepository repositories.AchievementRepository) *achievementService {
	return &achievementService{achievementRepository: achievementRepository}
}

func (s *achievementService) AddAchievement(achievement *domain.Achievement, image *multipart.FileHeader) (*domain.Achievement, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	//initialize achievement repo
	repo := repositories.NewAchievementRepository(conn)

	// Check if the member exist or not by NameID
	_, err = repo.FindMemberByID(achievement.MemberID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Member not found",
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
	
	achievement.Image = urlImage

	achievement, err = repo.Insert(achievement)
	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to add achievement to database",
			Code:    http.StatusInternalServerError,
		}
	}

	return achievement, nil
}

func (s *achievementService) UpdateAchievement(achievement *domain.Achievement, image *multipart.FileHeader) (*domain.Achievement, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAchievementRepository(conn)

	// Check if the member exist or not by NameID
	_, err = repo.FindMemberByID(achievement.MemberID)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Member not found",
			Code:    http.StatusBadRequest,
		}
	}

	var urlImage string
	if image != nil {
		urlImage, err = helper.UploadImage(image)
	}
	

	if err!= nil {
		return nil, &ErrorMessage{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
	}
	
	achievement.Image = urlImage

	achievement, err = repo.Update(achievement)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to update achievement to database",
			Code:    http.StatusInternalServerError,
		}
	}

	return achievement, nil
}

func (s *achievementService) DeleteAchievement(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAchievementRepository(conn)

	_, err = repo.FindByID(id)

	if err != nil {
		return &ErrorMessage{
			Message: "Member not found",
			Code:    http.StatusBadRequest,
		}
	}

	err = repo.Delete(id)
	if err != nil {
		return &ErrorMessage{
			Message: "Failed to delete achievement from database",
			Code:    http.StatusInternalServerError,
		}
	}
	return err

}

func (s *achievementService) GetAchievementByID(id string) (*domain.Achievement, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAchievementRepository(conn)

	achievement, err := repo.FindByID(id)
	if err != nil {
		return nil, &ErrorMessage{
			Message: "Member not found",
			Code:    http.StatusNotFound,
		}
	}

	return achievement, nil
}

func (s *achievementService) GetAllAchievements() ([]domain.Achievement, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewAchievementRepository(conn)

	achievements, err := repo.GetAllAchievements()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to retrieve achievements data",
			Code:    http.StatusInternalServerError,
		}
	}

	return achievements, nil
}
