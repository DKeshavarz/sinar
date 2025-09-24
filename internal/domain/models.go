package domain

type User struct {
	ID         int        `json:"id"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Phone      string     `json:"phone"`
	ProfilePic string     `json:"profile_pic"`
	StudentNum string     `json:"student_num"`
	Sex        bool       `json:"sex"`
	University University `json:"university"`
}

type University struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Logo     string `json:"logo"`
}

type Restaurant struct {
	ID           int    `json:"id"`
	UniversityID int    `json:"university_id"`
	Name         string `json:"name"`
	Sex          bool   `json:"sex"`
	Color        string `json:"color"`
}

type Food struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserFood struct {
	ID           int     `json:"id"`
	UserID       int     `json:"user_id"`
	FoodID       int     `json:"food_id"`
	RestaurantID int     `json:"Restaurant_id"`
	Price        float64 `json:"price"`
	Code         string  `json:"code"`
	TTL          int     `json:"ttl"`
}
