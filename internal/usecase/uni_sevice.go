package usecase

import "github.com/DKeshavarz/sinar/internal/domain"

type UnivercityStore interface {
	Get(id int)(*domain.University, error)
}

type Univercity interface {
	Get(id int)(*domain.University, error)
}

type univercity struct{

} 

