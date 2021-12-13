package service

import (
	"class-app/models"
	"class-app/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const (
	salt       = "ads56d148a9sd1a0sd65asd74"
	SigningKey = "=s90a7df=sda06f9sa=d0fsdaf6"
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
	teacher.Password = s.generatePasswordHash(teacher.Password)
	return s.repo.CreateTeacher(teacher)
}

func (s *AuthService) GenerateToken(email string, password string) (string, error) {
	teacher, err := s.repo.GetTeacher(email, s.generatePasswordHash(password))
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

	return token.SignedString([]byte(SigningKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
