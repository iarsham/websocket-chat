package entities

type UserRequest struct {
	Username string `json:"username" validate:"required" example:"ana"`
	Password string `json:"password" validate:"required,min=8" example:"geckodriver123"`
}
