package routers

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iarsham/websocket-chat/internal/controllers"
	"github.com/iarsham/websocket-chat/internal/middleware"
	"github.com/iarsham/websocket-chat/internal/services"
	"go.uber.org/zap"
)

func roomsGroup(r *mux.Router, db *sql.DB, log *zap.Logger, m *middleware.Middleware) *mux.Router {
	repo := services.NewRoomsService(db, log)
	c := &controllers.RoomsController{
		Service: repo,
	}
	roomAPI := r.PathPrefix("/rooms").Subrouter()
	roomAPI.Use()
	roomAPI.HandleFunc("/", c.GetAllRoomHandler).Methods(http.MethodGet)
	roomAPI.HandleFunc("/", c.CreateRoomHandler).Methods(http.MethodPost)
	roomAPI.HandleFunc("/{id}", c.GetRoomHandler).Methods(http.MethodGet)
	roomAPI.HandleFunc("/{id}", c.UpdateRoomHandler).Methods(http.MethodPut)
	roomAPI.HandleFunc("/{id}", c.DeleteRoomHandler).Methods(http.MethodDelete)
	return roomAPI
}
