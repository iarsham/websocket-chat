package controllers

import (
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (u *UsersController) UserLogOutHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.Service.Authenticate(w, r, uuid.UUID{}, false); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, constans.LoggedOut, http.StatusOK)
}
