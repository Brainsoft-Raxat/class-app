package repository

import (
	"class-app/models"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthPostgres struct {
	dbPool *pgxpool.Pool
}

func NewAuthPostgres(dbPool *pgxpool.Pool) *AuthPostgres {
	return &AuthPostgres{dbPool: dbPool}
}

func (r *AuthPostgres) CreateTeacher(teacher models.Teacher) (int, error) {
	var id int

	err := r.dbPool.QueryRow(context.Background(), "INSERT INTO teachers (email, password_hash, first_name, last_name, gender, phone_no) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", teacher.Email, teacher.Password, teacher.FirstName, teacher.LastName, teacher.Gender, teacher.PhoneNo).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetTeacher(email string, passwordHash string) (models.Teacher, error) {
	var teacher models.Teacher
	var id int32

	if err := r.dbPool.QueryRow(context.Background(), "SELECT * FROM teachers WHERE email=$1 AND password_hash=$2", email, passwordHash).Scan(
		&id,
		&teacher.Email,
		&teacher.Password,
		&teacher.FirstName,
		&teacher.LastName,
		&teacher.Gender,
		&teacher.PhoneNo,
	); err != nil {
		return models.Teacher{}, err
	}
	teacher.Id = int(id)

	return teacher, nil
}
