package domain

import (
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/internal/entites"
	"github.com/iarsham/websocket-chat/internal/models"
	"net/http"
)

type UserRepository interface {
	CreateUser(req *entites.UserRequest) (*models.Users, error)
	GetUserByID(id int64) (*models.Users, error)
	GetUserByUsername(userName string) (*models.Users, error)
	DeleteUser(userName string) error
	Authenticate(w http.ResponseWriter, r *http.Request, userID uuid.UUID, auth bool) error
}
