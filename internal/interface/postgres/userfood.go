package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/DKeshavarz/sinar/internal/domain"
	"github.com/DKeshavarz/sinar/internal/dto"
)

type UserFoodRepository struct {
	DB *sql.DB
}

func NewUserFoodRepository() *UserFoodRepository {
	return &UserFoodRepository{
		DB: getDB(),
	}
}

func (r *UserFoodRepository) GetAll() ([]*dto.UserFood, error) {
	query := `
        SELECT uf.id, u.id, u.first_name, u.last_name, u.phone, u.profile_pic, u.student_num, u.sex, u.university_id,
               r.id, r.university_id, r.name, r.sex, r.color,
               f.id, f.name,
               uf.user_id, uf.food_id, uf.restaurant_id, uf.price, uf.sinar_price, uf.code, uf.created_at, uf.expires_at
        FROM user_foods uf
        JOIN users u ON uf.user_id = u.id
        JOIN restaurants r ON uf.restaurant_id = r.id
        JOIN foods f ON uf.food_id = f.id`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userFoods []*dto.UserFood
	for rows.Next() {
		uf := &dto.UserFood{}
		user := &domain.User{}
		restaurant := &domain.Restaurant{}
		food := &domain.Food{}
		info := &domain.UserFood{}
		err := rows.Scan(
			&uf.Info.ID,
			&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.ProfilePic, &user.StudentNum, &user.Sex, &user.UniversityID,
			&restaurant.ID, &restaurant.UniversityID, &restaurant.Name, &restaurant.Sex, &restaurant.Color,
			&food.ID, &food.Name,
			&info.UserID, &info.FoodID, &info.RestaurantID, &info.Price, &info.SinarPrice, &info.Code, &info.CreatedAt, &info.ExpiresAt,
		)
		if err != nil {
			return nil, err
		}
		uf.User = user
		uf.Restaurant = restaurant
		uf.Food = food
		uf.Info = info
		userFoods = append(userFoods, uf)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if len(userFoods) == 0 {
		return nil, errors.New("no user-food relationships found")
	}
	return userFoods, nil
}

func (r *UserFoodRepository) GetByID(id int) (*dto.UserFood, error) {
	if id < 0 {
		return nil, errors.New("ID cannot be negative")
	}

	query := `
        SELECT uf.id, u.id, u.first_name, u.last_name, u.phone, u.profile_pic, u.student_num, u.sex, u.university_id,
               r.id, r.university_id, r.name, r.sex, r.color,
               f.id, f.name,
               uf.user_id, uf.food_id, uf.restaurant_id, uf.price, uf.sinar_price, uf.code, uf.created_at, uf.expires_at
        FROM user_foods uf
        JOIN users u ON uf.user_id = u.id
        JOIN restaurants r ON uf.restaurant_id = r.id
        JOIN foods f ON uf.food_id = f.id
        WHERE uf.id = $1`
	row := r.DB.QueryRow(query, id)

	uf := &dto.UserFood{}
	user := &domain.User{}
	restaurant := &domain.Restaurant{}
	food := &domain.Food{}
	info := &domain.UserFood{}
	err := row.Scan(
		&uf.Info.ID,
		&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.ProfilePic, &user.StudentNum, &user.Sex, &user.UniversityID,
		&restaurant.ID, &restaurant.UniversityID, &restaurant.Name, &restaurant.Sex, &restaurant.Color,
		&food.ID, &food.Name,
		&info.UserID, &info.FoodID, &info.RestaurantID, &info.Price, &info.SinarPrice, &info.Code, &info.CreatedAt, &info.ExpiresAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("user-food relationship not found")
	}
	if err != nil {
		return nil, err
	}
	uf.User = user
	uf.Restaurant = restaurant
	uf.Food = food
	uf.Info = info
	return uf, nil
}

func (r *UserFoodRepository) Create(userFood *domain.UserFood) error {
	query := `
        INSERT INTO user_foods (user_id, food_id, restaurant_id, price, sinar_price, code, expires_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, created_at`
	return r.DB.QueryRow(query, userFood.UserID, userFood.FoodID, userFood.RestaurantID,
		userFood.Price, userFood.SinarPrice, userFood.Code, userFood.ExpiresAt).Scan(&userFood.ID, &userFood.CreatedAt)
}

func (r *UserFoodRepository) GetActive() ([]*dto.UserFood, error) {
	query := `
        SELECT uf.id, u.id, u.first_name, u.last_name, u.phone, u.profile_pic, u.student_num, u.sex, u.university_id,
               r.id, r.university_id, r.name, r.sex, r.color,
               f.id, f.name,
               uf.user_id, uf.food_id, uf.restaurant_id, uf.price, uf.sinar_price, uf.code, uf.created_at, uf.expires_at
        FROM user_foods uf
        JOIN users u ON uf.user_id = u.id
        JOIN restaurants r ON uf.restaurant_id = r.id
        JOIN foods f ON uf.food_id = f.id
        WHERE uf.expires_at > $1`
	rows, err := r.DB.Query(query, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userFoods []*dto.UserFood
	for rows.Next() {
		uf := &dto.UserFood{}
		user := &domain.User{}
		restaurant := &domain.Restaurant{}
		food := &domain.Food{}
		info := &domain.UserFood{}
		err := rows.Scan(
			&uf.Info.ID,
			&user.ID, &user.FirstName, &user.LastName, &user.Phone, &user.ProfilePic, &user.StudentNum, &user.Sex, &user.UniversityID,
			&restaurant.ID, &restaurant.UniversityID, &restaurant.Name, &restaurant.Sex, &restaurant.Color,
			&food.ID, &food.Name,
			&info.UserID, &info.FoodID, &info.RestaurantID, &info.Price, &info.SinarPrice, &info.Code, &info.CreatedAt, &info.ExpiresAt,
		)
		if err != nil {
			return nil, err
		}
		uf.User = user
		uf.Restaurant = restaurant
		uf.Food = food
		uf.Info = info
		userFoods = append(userFoods, uf)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return userFoods, nil
}

func (r *UserFoodRepository) MarkAsUsed(id int) error {
	if id < 0 {
		return errors.New("ID cannot be negative")
	}

	// Set expiration time to past to mark as used
	query := `UPDATE user_foods SET expires_at = $1 WHERE id = $2`
	result, err := r.DB.Exec(query, time.Now().Add(-time.Hour), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("userfood not found")
	}

	return nil
}
