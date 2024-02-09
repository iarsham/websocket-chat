package routers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/iarsham/websocket-chat/internal/controllers"
	"github.com/iarsham/websocket-chat/internal/middleware"
	"github.com/iarsham/websocket-chat/internal/services"
	"go.uber.org/zap"
	"net/http"
)

func usersGroup(r *mux.Router, db *sql.DB, log *zap.Logger, store *sessions.CookieStore, m *middleware.Middleware) *mux.Router {
	repo := services.NewUserService(db, log, store)
	c := &controllers.UsersController{
		Service: repo,
	}
	usersAPI := r.PathPrefix("/users").Subrouter()
	usersAPI.Use(m.Authenticate)
	usersAPI.HandleFunc("/logout", c.UserLogOutHandler).Methods(http.MethodPost)
	usersAPI.HandleFunc("/delete-account", c.UserDeleteHandler).Methods(http.MethodDelete)
	return usersAPI
}
