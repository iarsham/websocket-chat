package controllers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/entites"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (u *UsersController) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entites.UserRequest)
	if err := common.BindJson(r, data); err != nil {
		responses.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := u.Service.GetUserByUsername(data.Username)
	if errors.Is(err, sql.ErrNoRows) {
		responses.Json(w, constans.UserNotExists, http.StatusNotFound)
		return
	}
	if ok := user.ValidatePassword(data.Password); !ok {
		responses.Json(w, constans.PassNotEqual, http.StatusUnauthorized)
		return
	}
	if err := u.Service.Authenticate(w, r, user.ID, true); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, user, http.StatusOK)
}
