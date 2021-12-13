package service

import (
	"class-app/models"
	"class-app/pkg/repository"
)

type Service struct {
	Student
	Authorization
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Student:       NewStudentService(repos.Student),
		Authorization: NewAuthService(repos.Authorization),
	}
}
