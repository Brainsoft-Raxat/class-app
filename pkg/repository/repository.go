package repository

import (
	"class-app/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Student
	Authorization
}

type Authorization interface {
	CreateTeacher(teacher models.Teacher) (int, error)
	GetTeacher(email string, passwordHash string) (models.Teacher, error)
}

type Student interface {
	CreateStudent(student models.Student) (int, error)
	GetAllStudents() ([]models.Student, error)
	GetStudentById(studentId int) (models.Student, error)
	DeleteAllStudents() error
	DeleteStudentById(studentId int) error
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Student:       NewStudentPostgres(dbPool),
		Authorization: NewAuthPostgres(dbPool),
	}
}
