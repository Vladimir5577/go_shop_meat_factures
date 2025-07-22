package model

type User struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Is_active bool   `json:"is_active"`
	Created   string `json:"created_at"`
	Updated   string `json:"updated_at"`
}

type UserRegistration struct {
	Name     string `json:"name" validate:"required,min=3,max=20"`
	Password string `json:"password" validate:"required,min=3,max=20"`
	Phone    string `json:"phone" validate:"required,min=5,max=20"`
	Address  string `json:"address" validate:"min=3,max=150"`
}

type UserLogin struct {
	Name     string `json:"name"`
	Password string `json:"password" validate:"required,min=3,max=20"`
	Phone    string `json:"phone" validate:"required,min=5,max=20"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
