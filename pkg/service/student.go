package service

import (
	"class-app/models"
	"class-app/pkg/repository"
)

type StudentService struct {
	repo repository.Student
}

func NewStudentService(repo repository.Student) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student models.Student) (int, error) {
	student.Password = generatePasswordHash(student.Password)
	return s.repo.CreateStudent(student)
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	return s.repo.GetAllStudents()
}

func (s *StudentService) GetStudentById(studentId int) (models.Student, error) {
	return s.repo.GetStudentById(studentId)
}

func (s *StudentService) DeleteAllStudents() error {
	return s.repo.DeleteAllStudents()
}

func (s *StudentService) DeleteStudentById(studentId int) error {
	return s.repo.DeleteStudentById(studentId)
}
