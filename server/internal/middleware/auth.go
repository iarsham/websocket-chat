package middleware

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store, _ := m.store.Get(r, constans.Session)
		if auth := store.Values["authenticated"]; auth != nil && auth != false {
			next.ServeHTTP(w, r)
			return
		}
		responses.Json(w, constans.AuthRequired, http.StatusUnauthorized)
		return
	})
}
