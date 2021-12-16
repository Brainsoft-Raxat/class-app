package repository

import (
	"class-app/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type ClassPostgres struct {
	dbPool *pgxpool.Pool
}

func NewClassPostgres(dbPool *pgxpool.Pool) *ClassPostgres {
	return &ClassPostgres{dbPool: dbPool}
}

func (r *ClassPostgres) GetAllClasses() ([]models.Class, error) {
	var classes []models.Class

	rows, err := r.dbPool.Query(context.Background(), "SELECT * FROM classes")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		classes = append(classes, models.Class{
			Id:   int(values[0].(int32)),
			Name: values[1].(string),
		})
	}

	return classes, nil
}

func (r *ClassPostgres) ListClassStudents(classId int) ([]models.Student, error) {
	var students []models.Student

	rows, err := r.dbPool.Query(context.Background(), "SELECT id, email, first_name, last_name, gender, status, class_id FROM students WHERE class_id=$1", classId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		students = append(students, models.Student{
			Id:        int(values[0].(int32)),
			Email:     values[1].(string),
			FirstName: values[2].(string),
			LastName:  values[3].(string),
			Gender:    values[4].(string),
			Status:    values[5].(bool),
			ClassId:   int(values[6].(int32)),
		})
	}

	return students, nil
}
