package models

//type Student struct {
//	Id        int32  `json:"id"`
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//	Gender    string `json:"gender"`
//	Status    bool   `json:"status"`
//}

type Student struct {
	Id           int32  `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	Status       bool   `json:"status"`
	ClassId      int32  `json:"class_id"`
}
