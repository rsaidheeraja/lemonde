package models

type Customer struct {
	FirstName     string `json:"firstname"`
	LastName      string `json:"lastname"`
	ContactNumber string `json:"contactnumber"`
	Email         string `json:"email"`
}
