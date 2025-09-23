package domain

import "time"

type User struct {
	StudentNumber  string      `json:"student_number"`
	UniversityName string      `json:"university_name"`
	Name           string      `json:"name"`
	PhoneNumber    string      `json:"phone_number"`
	University     *University `json:"university"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type University struct {
	Name          string    `json:"name"`
	LogoURL       string    `json:"logo_url"`
	ProfilePicURL string    `json:"profile_pic_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
