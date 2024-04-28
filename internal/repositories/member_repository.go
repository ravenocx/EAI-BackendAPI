package repositories

import (
	"strings"

	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"gorm.io/gorm"
)

type MemberRepository interface {
	Insert(member *domain.Member) (*domain.Member, error)
	Update(member *domain.Member) (*domain.Member, error)
	Delete(id string) error
	FindAllMembers() ([]domain.Member, error)
	FindMemberByID(id string) (*domain.Member, error)
	FindMemberByRole(role string) ([]domain.Member, error)
	FindAchievementByMemberID(memberId string) ([]domain.Achievement, error)
}

type memberRepository struct {
	db *gorm.DB
}

func NewMemberRepository(db *gorm.DB) *memberRepository {
	db = db.Debug()
	return &memberRepository{db: db}
}

func (r *memberRepository) Insert(member *domain.Member) (*domain.Member, error) {
	if err := r.db.Create(member).Error; err != nil {
		return nil, err
	}
	return member, nil
}

func (r *memberRepository) Update(member *domain.Member) (*domain.Member, error) {
	var updatedMember domain.Member

	if member.ID != "" {
		updatedMember.ID = member.ID
	}

	if member.Name != ""{
		updatedMember.Name = member.Name
	}

	if member.Role != ""{
		updatedMember.Role = member.Role
	}

	if member.Status != "" {
		updatedMember.Status = member.Status
	}

	if member.Image != ""{
		updatedMember.Image = member.Image
	}

	if err := r.db.Updates(&updatedMember).Error; err != nil {
		return nil, err
	}
	return member, nil
}

func (r *memberRepository) Delete(id string) error {
	var member domain.Member
	if err := r.db.Where("id = ?", id).Delete(&member).Error; err != nil {
		return err
	}
	return nil
}

func (r *memberRepository) GetMemberByID(id string) (*domain.Member, error) {
	var member domain.Member
	if err := r.db.Where("id = ?", id).First(&member).Error;err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) FindAllMembers() ([]domain.Member, error) {
	var members []domain.Member
	if err := r.db.Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) FindMemberByID(id string) (*domain.Member, error) {
	var member domain.Member

	if err := r.db.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *memberRepository) FindMemberByRole(role string) ([]domain.Member, error) {
	var members []domain.Member

	role = strings.ToLower(role)

	if err := r.db.Where("LOWER(role) LIKE ?", "%"+role+"%").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *memberRepository) FindAchievementByMemberID(memberId string) ([]domain.Achievement, error){
	var achievement []domain.Achievement

	memberId = strings.ToLower(memberId)

	if err := r.db.Where("LOWER(member_id) LIKE ?", "%"+memberId+"%" ).First(&achievement).Error; err != nil {
		return nil, err
	}
	return achievement, nil


}
