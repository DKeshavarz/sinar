package usecase

import "github.com/DKeshavarz/sinar/internal/domain"

type UnivercityStore interface {
	Get(id int)(*domain.University, error)
}

type Univercity interface {
	Get(id int)(*domain.University, error)
}

type univercity struct{
	uniRepo UnivercityStore
} 

func NewUnivercity(storage UnivercityStore) Univercity{
	return &univercity{
		uniRepo: storage,
	}
}

func (u *univercity)Get(id int)(*domain.University, error){
	return u.uniRepo.Get(id)
}

