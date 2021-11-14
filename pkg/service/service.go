package service

import (
	"class-app/models"
	"class-app/pkg/repository"
)

type Service struct {
	Student
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
		Student: NewStudentService(repos.Student),
	}
}
