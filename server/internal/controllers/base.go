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

type WsController struct {
	Service domain.WsRepository
}

func getFromCtx(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
