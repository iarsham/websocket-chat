package routers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/middleware"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

func SetupRoutes(db *sql.DB, rds *redis.Client, log *zap.Logger) http.Handler {
	r := mux.NewRouter()
	store := common.SessionToRedis()
	m := middleware.NewMiddleware(log, store)
	r.Use(m.Recovery)
	r.Use(m.LogMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	usersGroup(api, db, log, store)
	return m.CorsMiddleware().Handler(r)
}
