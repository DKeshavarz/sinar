package postgres

import (
    "database/sql"
    "errors"
    "github.com/DKeshavarz/sinar/internal/domain"
)

type RestaurantRepository struct {
    DB *sql.DB
}

func NewRestaurantRepository() *RestaurantRepository{
	return &RestaurantRepository{
		DB: getDB(),
	}
}


func (r *RestaurantRepository) GetAll(uniID int) ([]*domain.Restaurant, error) {
    if uniID < 0 {
        return nil, errors.New("university ID cannot be negative")
    }

    query := `
        SELECT id, university_id, name, sex, color
        FROM restaurants
        WHERE university_id = $1`
    rows, err := r.DB.Query(query, uniID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var restaurants []*domain.Restaurant
    for rows.Next() {
        restaurant := &domain.Restaurant{}
        err := rows.Scan(&restaurant.ID, &restaurant.UniversityID, &restaurant.Name, &restaurant.Sex, &restaurant.Color)
        if err != nil {
            return nil, err
        }
        restaurants = append(restaurants, restaurant)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    if len(restaurants) == 0 {
        return nil, errors.New("no restaurants found for the given university ID")
    }
    return restaurants, nil
}