package usecase

import (
	"errors"
	"time"

	"github.com/DKeshavarz/sinar/internal/domain"
	"github.com/DKeshavarz/sinar/internal/dto"
)

type UserFoodStore interface {
	GetAll() ([]*dto.UserFood, error)
	GetByID(id int) (*dto.UserFood, error)
	GetActive() ([]*dto.UserFood, error)
	Create(userFood *domain.UserFood) error
	MarkAsUsed(id int) error
}

type UserFood interface {
	GetAll() ([]*dto.UserFood, error)
	GetByID(id int) (*dto.UserFood, error)
	GetActive() ([]*dto.UserFood, error)
	Purchase(userID, foodID, restaurantID, price, sinarPrice int, code string, expirationHours int) (*domain.UserFood, error)
	MarkAsUsed(id int) error
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

func (uf *userFood) Purchase(userID, foodID, restaurantID, price, sinarPrice int, code string, expirationHours int) (*domain.UserFood, error) {
	if userID <= 0 || foodID <= 0 || restaurantID <= 0 {
		return nil, errors.New("invalid user, food, or restaurant ID")
	}
	if price < 0 || sinarPrice < 0 {
		return nil, errors.New("price cannot be negative")
	}
	if code == "" {
		return nil, errors.New("code cannot be empty")
	}
	if expirationHours <= 0 {
		return nil, errors.New("expiration hours must be positive")
	}

	// Calculate expiration time
	expiresAt := time.Now().Add(time.Duration(expirationHours) * time.Hour)

	userFood := &domain.UserFood{
		UserID:       userID,
		FoodID:       foodID,
		RestaurantID: restaurantID,
		Price:        price,
		SinarPrice:   sinarPrice,
		Code:         code,
		ExpiresAt:    expiresAt,
	}

	if err := uf.store.Create(userFood); err != nil {
		return nil, errors.New("failed to create purchase: " + err.Error())
	}

	return userFood, nil
}

func (uf *userFood) GetActive() ([]*dto.UserFood, error) {
	result, err := uf.store.GetActive()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uf *userFood) MarkAsUsed(id int) error {
	if id < 0 {
		return errors.New("ID cannot be negative")
	}
	return uf.store.MarkAsUsed(id)
}
