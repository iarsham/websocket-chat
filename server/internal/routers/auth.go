package routers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/iarsham/websocket-chat/internal/controllers"
	"github.com/iarsham/websocket-chat/internal/services"
	"go.uber.org/zap"
	"net/http"
)

func authGroup(r *mux.Router, db *sql.DB, log *zap.Logger, store *sessions.CookieStore) *mux.Router {
	repo := services.NewUserService(db, log, store)
	c := &controllers.UsersController{
		Service: repo,
	}
	authAPI := r.PathPrefix("/auth").Subrouter()
	authAPI.HandleFunc("/register", c.UserRegisterHandler).Methods(http.MethodPost)
	authAPI.HandleFunc("/login", c.UserLoginHandler).Methods(http.MethodPost)
	return authAPI
}
