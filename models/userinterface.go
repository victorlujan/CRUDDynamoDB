package models

type User struct {
	Id string  `json:"id"`
	Email string  `json:"email"`
	Password string 	`json:"password"`
	First_name string  `json:"first_name"`
	Last_name string   `json:"last_name"`
	Phone string  `json:"phone"`
}	
