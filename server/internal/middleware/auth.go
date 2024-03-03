package middleware

import (
	"context"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (m *Middleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store, _ := m.store.Get(r, constans.Session)
		if auth := store.Values[constans.AuthKey]; auth != nil && auth != false {
			r = r.WithContext(context.WithValue(r.Context(), constans.UserID, store.Values[constans.UserID]))
			r = r.WithContext(context.WithValue(r.Context(), constans.Username, store.Values[constans.Username]))
			next.ServeHTTP(w, r)
			return
		}
		responses.Json(w, constans.AuthRequired, http.StatusUnauthorized)
		return
	})
}
