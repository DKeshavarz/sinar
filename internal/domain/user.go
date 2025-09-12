package domain

import "time"

type User struct {
	StudentNumber  string    `json:"student_number"`
	UniversityName string    `json:"university_name"`
	Name           string    `json:"name"`
	PhoneNumber    string    `json:"phone_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type otp struct {
	Key  string `json:"key"`
	Code string `json:"code"`
}
