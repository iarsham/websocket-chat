package responses

type QueueResponse struct {
	RoomID  uint   `json:"room_id"`
	Message string `json:"message"`
}
