package usecase

import (
	"errors"

	"github.com/DKeshavarz/sinar/internal/dto"
)

type UserStore interface {
    GetByStudentNumber(number string) (*dto.UserWithUniversity, error)
}

type User interface {
    GetByStudentNumber(number string) (*dto.UserWithUniversity, error)
}

type user struct {
    store UserStore
}

func NewUser(store UserStore) User {
    return &user{store: store}
}

func (u *user) GetByStudentNumber(number string) (*dto.UserWithUniversity, error) {
    if number == "" {
        return nil, errors.New("student number cannot be empty")
    }
    return u.store.GetByStudentNumber(number)
}