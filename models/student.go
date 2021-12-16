package models

//type Student struct {
//	Id        int32  `json:"id"`
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//	Gender    string `json:"gender"`
//	Status    bool   `json:"status"`
//}

type Student struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Status    bool   `json:"status"`
	ClassId   int    `json:"class_id"`
}
