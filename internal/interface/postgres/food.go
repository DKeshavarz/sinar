package postgres

import (
    "database/sql"
    "errors"
    "github.com/DKeshavarz/sinar/internal/domain"
)

type FoodRepository struct {
    DB *sql.DB
}

func NewFoodRepository() *FoodRepository{
	return &FoodRepository{
		DB: getDB(),
	}
}

func (r *FoodRepository) GetAll() ([]*domain.Food, error) {
    query := `
        SELECT id, name
        FROM foods`
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var foods []*domain.Food
    for rows.Next() {
        food := &domain.Food{}
        err := rows.Scan(&food.ID, &food.Name)
        if err != nil {
            return nil, err
        }
        foods = append(foods, food)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    if len(foods) == 0 {
        return nil, errors.New("no foods found")
    }
    return foods, nil
}