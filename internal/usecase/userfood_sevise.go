package usecase

import (
	"github.com/DKeshavarz/sinar/internal/domain"
	"github.com/DKeshavarz/sinar/internal/dto"
)

type UserFoodStore interface {
	GetAll() (domain.UserFood, error)
	GetByID(id int) (domain.UserFood, error)
}

type UserFood interface {
	GetAll() (dto.UserFood, error)
	GetByID(id int) (dto.UserFood, error)
}
