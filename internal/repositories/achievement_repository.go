package repositories

import (
	"log"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"gorm.io/gorm"
)

type AchievementRepository interface {
	Insert(achievement *domain.Achievement) (*domain.Achievement, error)
	Update(achievement *domain.Achievement) (*domain.Achievement, error)
	Delete(id string) error
	FindByID(id string) (*domain.Achievement, error)
	FindMemberByID(id string) (*domain.Member, error)
	GetAllAchievements() ([]domain.Achievement, error)
}

type achievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository(db *gorm.DB) *achievementRepository {
	db = db.Debug()
	return &achievementRepository{db: db}
}


func (r *achievementRepository) Insert(achievement *domain.Achievement) (*domain.Achievement, error) {
	if err := r.db.Create(achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (r *achievementRepository) Update(achievement *domain.Achievement) (*domain.Achievement, error) {
	var updatedAchievement domain.Achievement

	if achievement.ID != ""{
		updatedAchievement.ID = achievement.ID
	}

	if achievement.MemberID != "" {
		updatedAchievement.MemberID = achievement.MemberID
		updatedAchievement.Member = achievement.Member
	}

	if achievement.Image != ""{
		updatedAchievement.Image = achievement.Image
	}

	if achievement.Achievement != ""{
		updatedAchievement.Achievement = achievement.Achievement
	}

	log.Printf("Achievement to update : %+v",updatedAchievement)
	if err := r.db.Updates(&updatedAchievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil
}

func (r *achievementRepository) Delete(id string) error {
	var achievement domain.Achievement
	if err := r.db.Where("id = ?", id).Delete(&achievement).Error; err != nil {
		return err
	}
	return nil
}

func (r *achievementRepository) FindByID(id string) (*domain.Achievement, error) {
	var achievement domain.Achievement
	if err := r.db.Preload("Member").Where("id = ?", id).First(&achievement).Error; err != nil {
		return nil, err
	}
	return &achievement, nil
}

func (r *achievementRepository) FindMemberByID(id string) (*domain.Member, error) {
	var member domain.Member

	if err := r.db.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *achievementRepository) GetAllAchievements() ([]domain.Achievement, error) {
	var achievements []domain.Achievement

	if err := r.db.Preload("Name").Find(&achievements).Error; err != nil {
		return nil, err
	}

	return achievements, nil
}

