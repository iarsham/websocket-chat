package controllers

import (
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

// UserRegisterHandler
//
//	@Summary		Register endpoint
//	@Description	Register with username and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		entities.UserRequest				true	"register body"
//	@Success		201		{object}	responses.RegisterOKResponse		"Success"
//	@Failure		409		{object}	responses.UserAlreadyExistsResponse	"Error"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/auth/register [post]
func (u *UsersController) UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
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
