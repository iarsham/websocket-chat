package responses

import "github.com/iarsham/websocket-chat/internal/models"

type LoginOKResponse *models.Users
type RegisterOKResponse *models.Users

type InterServerErrorResponse struct {
	Response string `example:"Internal Server Error"`
}

type UserNotExistsResponse struct {
	Response string `example:"User not found"`
}

type PassNotEqualResponse struct {
	Response string `example:"Password is incorrect"`
}

type UserAlreadyExistsResponse struct {
	Response string `example:"Password is incorrect"`
}

type RecordDeletedResponse struct{}

type UserLogOutResponse struct {
	Response string `example:"User logged out"`
}
