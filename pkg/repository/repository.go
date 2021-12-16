package repository

import (
	"class-app/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	Authorization
	Student
	Class
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

type Class interface {
	GetAllClasses() ([]models.Class, error)
	ListClassStudents(classId int) ([]models.Student, error)
}

func NewRepository(dbPool *pgxpool.Pool) *Repository {
	return &Repository{
		Student:       NewStudentPostgres(dbPool),
		Authorization: NewAuthPostgres(dbPool),
		Class:         NewClassPostgres(dbPool),
	}
}
