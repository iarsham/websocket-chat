package middleware

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (m *Middleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				responses.Json(w, constans.InternalError, http.StatusBadRequest)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
