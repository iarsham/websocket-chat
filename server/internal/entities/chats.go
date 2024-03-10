package entities

type QueueRequest struct {
	RoomID uint   `json:"room_id"`
	Code   string `json:"code"`
}
