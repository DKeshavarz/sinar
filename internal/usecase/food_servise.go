package usecase

import (
	"github.com/DKeshavarz/sinar/internal/domain"
)

type FoodStore interface {
	GetAll() (domain.Food, error)
}

type Food interface {
	GetAllNames() (domain.Food, error)
}
