package service

import (
	"class-app/configs"
	"class-app/models"
	"class-app/pkg/repository"
	"github.com/golang-jwt/jwt"
	"time"
)

type TokenClaims struct {
	TeacherId int `json:"teacher_id"`
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateTeacher(teacher models.Teacher) (int, error) {
	teacher.Password = generatePasswordHash(teacher.Password)
	return s.repo.CreateTeacher(teacher)
}

func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	teacher, err := s.repo.GetTeacher(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// generating token here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		TeacherId: teacher.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(configs.SigningKey))
}
