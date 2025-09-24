package dto

import "github.com/DKeshavarz/sinar/internal/domain"

type UserWithUniversity struct {
	User       *domain.User       `json:"user"`
	University *domain.University `json:"university"`
}

type UserFood struct {
	User       *domain.User       `json:"user"`
	Restaurant *domain.Restaurant `json:"restaurant"`
	Food       *domain.Food       `json:"food"`
	Info       *domain.UserFood   `json:"info"`
}
