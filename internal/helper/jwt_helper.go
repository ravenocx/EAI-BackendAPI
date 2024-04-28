package helper

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
	"github.com/ravenocx/EAI-BackendAPI/internal/dto"
)

func GenerateAdminToken(admin *domain.Admin) (string, error) {
	claims := dto.Claims{
		AdminID: admin.ID,
		Email:   admin.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "tel-ulab.ac.id",
			Subject:   admin.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("sUp3rS3cREtK3yf0rw3BpR0f1Le"))

	adminToken := domain.Token{
		AdminID: admin.ID,
		Token:   signedToken,
		IssuedAt: claims.IssuedAt.Time,
		ExpiredAt: claims.ExpiresAt.Time,
	}

	var existingAdminToken domain.Token

	// check if admin already have token or not
	err := config.DB.Where("admin_id = ?", admin.ID).First(&existingAdminToken).Error
	if err != nil { //admin doesnt have token
		err = config.DB.Create(&adminToken).Error
	}else{
		err = config.DB.Model(&existingAdminToken).Updates(&adminToken).Error
	}

	return signedToken, err

}
