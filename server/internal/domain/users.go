package domain

import (
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/internal/models"
	"net/http"
)

type UserRepository interface {
	CreateUser(req *entities.UserRequest) (*models.Users, error)
	GetUserByID(id string) (*models.Users, error)
	GetUserByUsername(userName string) (*models.Users, error)
	DeleteUser(userID string) error
	Authenticate(w http.ResponseWriter, r *http.Request, userID uuid.UUID, auth bool) error
}
