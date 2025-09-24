package usecase

import "github.com/DKeshavarz/sinar/internal/dto"

type UserStore interface {
	GetByStudentNumber(number string)
}

type User interface {
	GetByStudentNumber(number string) (*dto.UserWithUniversity, error)
}


type user struct{

} 

