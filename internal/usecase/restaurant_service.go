package usecase

import (
	"errors"

	"github.com/DKeshavarz/sinar/internal/domain"
)

type RestaurantStore interface {
    GetAll(uniID int) ([]*domain.Restaurant, error)
}

type Restaurant interface {
    GetAll(uniID int) ([]*domain.Restaurant, error)
}

type restaurant struct {
    store RestaurantStore
}

func NewRestaurant(store RestaurantStore) Restaurant {
    return &restaurant{store: store}
}

func (r *restaurant) GetAll(uniID int) ([]*domain.Restaurant, error) {
    if uniID < 0 {
        return nil, errors.New("university ID cannot be negative")
    }
    return r.store.GetAll(uniID)
}