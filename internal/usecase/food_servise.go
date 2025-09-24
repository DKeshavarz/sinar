package usecase

import (
	"errors"
	"github.com/DKeshavarz/sinar/internal/domain"
)

type FoodStore interface {
    GetAll() ([]*domain.Food, error)
}

type Food interface {
    GetAllNames() ([]*domain.Food, error)
}

type food struct {
    store FoodStore
}

func NewFood(store FoodStore) Food {
    return &food{store: store}
}

func (f *food) GetAllNames() ([]*domain.Food, error) {
    foods, err := f.store.GetAll()
    if err != nil {
        return nil, err
    }
    if len(foods) == 0 {
        return nil, errors.New("no foods found")
    }
   
    return foods, nil
}