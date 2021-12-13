package models

type Teacher struct {
	Id        int    `json:"id" db:"id"`
	Email     string `json:"email" db:"email" binding:"required"`
	Password  string `json:"password_hash" db:"password_hash" binding:"required"`
	PhoneNo   string `json:"phone_no" db:"phone_no" binding:"required"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Gender    string `json:"gender" db:"gender" binding:"required"`
}
