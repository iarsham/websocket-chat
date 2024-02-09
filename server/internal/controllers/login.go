package controllers

import (
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/entites"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (u *UsersController) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entites.UserRequest)
	if err := common.BindJson(r, data); err != nil {
		responses.Json(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := u.Service.GetUserByUsername(data.Username); err != nil {
		user, err := u.Service.CreateUser(data)
		if err != nil {
			responses.Json(w, constans.InternalError, http.StatusInternalServerError)
			return
		}
		responses.Json(w, user, http.StatusCreated)
		return
	}
	responses.Json(w, constans.UserExists, http.StatusConflict)
}
