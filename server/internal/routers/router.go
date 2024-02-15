package routers

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/iarsham/websocket-chat/api"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/swaggo/http-swagger"
	"go.uber.org/zap"
	"net/http"
)

func redirectToDocs(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/index.html", http.StatusMovedPermanently)
}

func SetupRoutes(db *sql.DB, rds *redis.Client, log *zap.Logger) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", redirectToDocs).Methods(http.MethodGet)
	store := common.SessionToRedis()
	m := middleware.NewMiddleware(log, store)
	r.Use(m.Recovery)
	r.Use(m.LogMiddleware)
	api := r.PathPrefix("/api").Subrouter()
	authGroup(api, db, log, store)
	usersGroup(api, db, log, store, m)
	roomsGroup(api, db, log, m)
	serveSwagger(r)
	return m.CorsMiddleware().Handler(r)
}

func serveSwagger(r *mux.Router) *mux.Route {
	swagRoute := r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	return swagRoute
}
