package service

import (
	"class-app/models"
	"class-app/pkg/repository"
	"crypto/sha1"
	"fmt"
	"github.com/spf13/viper"
)

type Service struct {
	Authorization
	Student
	Class
}

type Authorization interface {
	CreateTeacher(teacher models.Teacher) (int, error)
	GenerateToken(email string, password string) (string, error)
}

type Student interface {
	CreateStudent(student models.Student) (int, error)
	GetAllStudents() ([]models.Student, error)
	GetStudentById(studentId int) (models.Student, error)
	DeleteAllStudents() error
	DeleteStudentById(studentId int) error
}

type Class interface {
	GetAllClasses() ([]models.Class, error)
	ListClassStudents(classId int) ([]models.Student, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Student:       NewStudentService(repos.Student),
		Authorization: NewAuthService(repos.Authorization),
		Class:         NewClassService(repos.Class),
	}
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	salt := viper.GetString("salt")
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
