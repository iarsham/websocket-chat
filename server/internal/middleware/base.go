package middleware

import (
	"github.com/gorilla/sessions"
	"go.uber.org/zap"
)

type Middleware struct {
	log   *zap.Logger
	store *sessions.CookieStore
}

func NewMiddleware(log *zap.Logger, store *sessions.CookieStore) *Middleware {
	return &Middleware{
		log:   log,
		store: store,
	}
}
