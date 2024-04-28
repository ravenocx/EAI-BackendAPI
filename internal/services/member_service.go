package services

import (
	"mime/multipart"
	"net/http"

	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/helper"
	"github.com/ravenocx/EAI-BackendAPI/internal/repositories"
)

type MemberService interface {
	AddMember(member *domain.Member, image *multipart.FileHeader) (*domain.Member, error)
	UpdateMember(member *domain.Member, image *multipart.FileHeader) (*domain.Member, error)
	DeleteMember(id string) error
	GetMemberByID(id string) (*domain.Member, error)
	GetAllMembers() ([]domain.Member, error)
	GetMemberByRole(role string) ([]domain.Member, error)
}

type memberService struct {
	memberRepository repositories.MemberRepository
}

func NewMemberService(memberRepository repositories.MemberRepository) *memberService {
	return &memberService{memberRepository: memberRepository}
}

func (s *memberService) AddMember(member *domain.Member, image *multipart.FileHeader) (*domain.Member, error) {
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

	member.Image = urlImage

	repo := repositories.NewMemberRepository(conn)

	member, err = repo.Insert(member)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to add member",
			Code:    http.StatusInternalServerError,
		}
	}

	return member, nil
}

func (s *memberService) UpdateMember(member *domain.Member, image *multipart.FileHeader) (*domain.Member, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewMemberRepository(conn)

	_, err = repo.GetMemberByID(member.ID)

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

	member.Image = urlImage

	member, err = repo.Update(member)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to update member to database",
			Code:    http.StatusInternalServerError,
		}
	}

	return member, nil
}

func (s *memberService) DeleteMember(id string) error {
	conn, err := config.Connect()

	if err != nil {
		return &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewMemberRepository(conn)

	_, err = repo.FindMemberByID(id)

	if err != nil {
		return &ErrorMessage{
			Message: "Member not found",
			Code:    http.StatusBadRequest,
		}
	}

	// check if the user has achievement or not
	achievement, err := repo.FindAchievementByMemberID(id)

	if achievement != nil { // if user does exist in achievement table , then cannot delete member data
		return &ErrorMessage{ // user does exist in achievement table
			Message: "Cannot delete this member as it is referenced by other entities",
			Code:    http.StatusConflict,
		}

	}

	if err != nil { // user doesnt exist in achievement table
		err = repo.Delete(id)

		if err != nil {
			return &ErrorMessage{
				Message: "Failed to delete member from database",
				Code:    http.StatusInternalServerError,
			}
		}
	}
	return err

}

func (s *memberService) GetMemberByID(id string) (*domain.Member, error) {
	conn, err := config.Connect()
	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewMemberRepository(conn)

	member, err := repo.GetMemberByID(id)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Member not found",
			Code:    http.StatusNotFound,
		}
	}

	return member, nil
}

func (s *memberService) GetAllMembers() ([]domain.Member, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewMemberRepository(conn)

	members, err := repo.FindAllMembers()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to retrieve members data",
			Code:    http.StatusInternalServerError,
		}
	}

	return members, nil
}

func (s *memberService) GetMemberByRole(role string) ([]domain.Member, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to connect to database",
			Code:    http.StatusInternalServerError,
		}
	}

	repo := repositories.NewMemberRepository(conn)

	members, err := repo.FindMemberByRole(role)

	if err != nil {
		return nil, &ErrorMessage{
			Message: "Failed to retrieve members data",
			Code:    http.StatusInternalServerError,
		}
	}

	return members, nil
}
