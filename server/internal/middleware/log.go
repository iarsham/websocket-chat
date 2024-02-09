package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func (m *Middleware) LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer func() {
			strLog := fmt.Sprintf("%s -  %s  -  (%v)", r.Method, r.URL.Path, time.Since(now))
			m.log.Info(strLog)
		}()
		next.ServeHTTP(w, r)
	})
}
