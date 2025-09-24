package usecase

import (
	"errors"
	"github.com/DKeshavarz/sinar/internal/dto"
)

type UserFoodStore interface {
    GetAll() ([]*dto.UserFood, error)
    GetByID(id int) (*dto.UserFood, error)
}

type UserFood interface {
    GetAll() ([]*dto.UserFood, error)
    GetByID(id int) (*dto.UserFood, error)
}

type userFood struct {
    store UserFoodStore
}

func NewUserFood(store UserFoodStore) UserFood {
    return &userFood{store: store}
}

func (uf *userFood) GetAll() ([]*dto.UserFood, error) {
    result, err := uf.store.GetAll()
    if err != nil {
        return nil, err
    }
    if len(result) == 0 {
        return nil, errors.New("no user-food relationships found")
    }
    return result, nil
}

func (uf *userFood) GetByID(id int) (*dto.UserFood, error) {
    if id < 0 {
        return nil, errors.New("ID cannot be negative")
    }
    result, err := uf.store.GetByID(id)
    if err != nil {
        return nil, err
    }
    if result == nil {
        return nil, errors.New("user-food relationship not found")
    }
    return result, nil
}