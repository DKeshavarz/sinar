package userfoodhandler

type UserFoodSingleRequest struct {
	UserID          int    `json:"user_id" example:"1"`
	FoodID          int    `json:"food_id" example:"101"`
	RestaurantID    int    `json:"restaurant_id" example:"12"`
	Price           int    `json:"price" example:"25000"`
	SinarPrice      int    `json:"sinar_price" example:"20000"`
	Code            string `json:"code" example:"ABC123"`
	ExpirationHours int    `json:"expiration_hours" example:"24"`
}

type UserFoodArrayRequest struct {
	UserID       int    `json:"user_id" example:"1"`
	FoodID       int    `json:"food_id" example:"101"`
	RestaurantID int    `json:"restaurant_id" example:"12"`
	Price        int    `json:"price" example:"25000"`
	SinarPrice   int    `json:"sinar_price" example:"20000"`
	Code         string `json:"code" example:"XYZ456"`
	ExpiresAt    string `json:"expires_at" example:"2025-12-31T23:59:59Z"`
}