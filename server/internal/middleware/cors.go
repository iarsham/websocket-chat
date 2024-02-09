package middleware

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/utils"
	"github.com/rs/cors"
	"net/http"
)

func (m *Middleware) CorsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: utils.GetListEnv(constans.ORIGINS),
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowCredentials: true,
		Debug:            constans.Mode,
		MaxAge:           300,
	})
}
