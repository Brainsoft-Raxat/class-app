package repository

import (
	"class-app/models"
	"context"
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

	err := r.dbPool.QueryRow(context.Background(), "INSERT INTO students (first_name, last_name, gender, status) VALUES ($1, $2, $3, $4) RETURNING id", student.FirstName, student.LastName, student.Gender, student.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *StudentPostgres) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	rows, err := r.dbPool.Query(context.Background(), "SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		student := models.Student{
			Id:        values[0].(int32),
			FirstName: values[1].(string),
			LastName:  values[2].(string),
			Gender:    values[3].(string),
			Status:    values[4].(bool),
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *StudentPostgres) GetStudentById(studentId int) (models.Student, error) {
	var student models.Student
	rows, err := r.dbPool.Query(context.Background(), "SELECT * FROM students WHERE id=$1", studentId)
	if err != nil {
		return models.Student{}, err
	}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return student, err
		}
		student = models.Student{
			Id:        values[0].(int32),
			FirstName: values[1].(string),
			LastName:  values[2].(string),
			Gender:    values[3].(string),
			Status:    values[4].(bool),
		}
	}

	//if int(student.Id) != studentId {
	//	return student, errors.New(fmt.Sprintf("user with id %d does not exist", studentId))
	//}

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
	_, err := r.dbPool.Exec(context.Background(), "DELETE FROM students WHERE id=$1", studentId)

	return err
}
