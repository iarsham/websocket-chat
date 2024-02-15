package entities

type RoomRequest struct {
	Name string `json:"name" validate:"required" example:"Warzone"`
}
