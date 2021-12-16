package service

import (
	"class-app/models"
	"class-app/pkg/repository"
)

type ClassService struct {
	repo repository.Class
}

func NewClassService(repo repository.Class) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) GetAllClasses() ([]models.Class, error) {
	return s.repo.GetAllClasses()
}

func (s *ClassService) ListClassStudents(classId int) ([]models.Student, error) {
	return s.repo.ListClassStudents(classId)
}
