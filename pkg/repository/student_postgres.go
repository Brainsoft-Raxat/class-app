package repository

import (
	"class-app/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StudentPostgres struct {
	dbPool *pgxpool.Pool
}

func NewStudentPostgres(dbPool *pgxpool.Pool) *StudentPostgres {
	return &StudentPostgres{dbPool: dbPool}
}

func (r *StudentPostgres) CreateStudent(student models.Student) (int, error) {
	var id int

	err := r.dbPool.QueryRow(context.Background(), "INSERT INTO students (email, password_hash, first_name, last_name, gender, status, class_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", student.Email, student.Password, student.FirstName, student.LastName, student.Gender, student.Status, student.ClassId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *StudentPostgres) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	rows, err := r.dbPool.Query(context.Background(), "SELECT id, email, first_name, last_name, gender, status, class_id FROM students")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		student := models.Student{
			Id:        int(values[0].(int32)),
			Email:     values[1].(string),
			FirstName: values[2].(string),
			LastName:  values[3].(string),
			Gender:    values[4].(string),
			Status:    values[5].(bool),
			ClassId:   int(values[6].(int32)),
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *StudentPostgres) GetStudentById(studentId int) (models.Student, error) {
	var student models.Student
	var id int32
	var classId int32

	if err := r.dbPool.QueryRow(context.Background(), "SELECT id, email, first_name, last_name, gender, status, class_id FROM students WHERE id=$1", studentId).Scan(
		&id,
		&student.Email,
		&student.FirstName,
		&student.LastName,
		&student.Gender,
		&student.Status,
		&classId,
	); err != nil {
		return models.Student{}, err
	}
	student.Id = int(id)
	student.ClassId = int(classId)

	return student, nil
}

func (r *StudentPostgres) DeleteAllStudents() error {
	_, err := r.dbPool.Exec(context.Background(), "TRUNCATE TABLE students RESTART IDENTITY CASCADE")
	return err
}

func (r *StudentPostgres) DeleteStudentById(studentId int) error {
	//_, err := r.GetStudentById(studentId)
	//if err != nil {
	//	return err
	//}
	var id int
	err := r.dbPool.QueryRow(context.Background(), "DELETE FROM students WHERE id=$1 RETURNING ID", studentId).Scan(&id)
	if err != nil {
		return errors.New(fmt.Sprintf("no user with id=%d", studentId))
	}
	return err
}
