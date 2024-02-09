package controllers

import (
	"github.com/iarsham/websocket-chat/internal/domain"
)

type UsersController struct {
	Service domain.UserRepository
}
