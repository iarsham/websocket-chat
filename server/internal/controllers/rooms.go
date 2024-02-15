package controllers

import (
	"database/sql"
	"errors"
	"github.com/gorilla/mux"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (rc *RoomsController) GetAllRoomHandler(w http.ResponseWriter, r *http.Request) {
	rooms, err := rc.Service.GetAllRooms()
	if err != nil {
		responses.Json(w, constans.RoomNotFound, http.StatusNotFound)
		return
	}
	responses.Json(w, rooms, http.StatusOK)
}

func (rc *RoomsController) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)[constans.ID]
	room, err := rc.Service.GetRoomByID(roomID)
	if errors.Is(err, sql.ErrNoRows) {
		responses.Json(w, constans.RoomNotFound, http.StatusNotFound)
		return
	}
	responses.Json(w, room, http.StatusOK)
}

func (rc *RoomsController) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.RoomRequest)
	if err := common.BindJson(r, data); err != nil {
		responses.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	room, err := rc.Service.CreateRoom(data)
	if err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, room, http.StatusCreated)
}

func (rc *RoomsController) UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)[constans.ID]
	data := new(entities.RoomRequest)
	if err := common.BindJson(r, data); err != nil {
		responses.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := rc.Service.GetRoomByID(roomID); errors.Is(err, sql.ErrNoRows) {
		responses.Json(w, constans.RoomNotFound, http.StatusNotFound)
		return
	}
	room, err := rc.Service.UpdateRoom(data, roomID)
	if err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, room, http.StatusOK)
}

func (rc *RoomsController) DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)[constans.ID]
	if _, err := rc.Service.GetRoomByID(roomID); errors.Is(err, sql.ErrNoRows) {
		responses.Json(w, constans.RoomNotFound, http.StatusNotFound)
		return
	}
	if err := rc.Service.DeleteRoom(roomID); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, nil, http.StatusNoContent)
}
