package usecase

import "github.com/DKeshavarz/sinar/internal/domain"

type RestaurantStore interface {
	GetAll(uniID int)(*domain.Restaurant, error)
}

type Restaurant interface {
	GetAll(uniID int)(*domain.Restaurant, error)
}