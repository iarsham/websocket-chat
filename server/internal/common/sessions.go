package common

import (
	"github.com/gorilla/sessions"
	"github.com/iarsham/websocket-chat/pkg/constans"
)

func SessionToRedis() *sessions.CookieStore {
	store := sessions.NewCookieStore([]byte(constans.Key))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   constans.SessionExpire * 60 * 60,
	}
	return store
}
