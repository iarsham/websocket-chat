package controllers

import (
	"context"
	"github.com/iarsham/websocket-chat/internal/domain"
)

type UsersController struct {
	Service domain.UserRepository
}

type RoomsController struct {
	Service domain.RoomsRepository
}

func (u *UsersController) getFromCtx(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
