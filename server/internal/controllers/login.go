package controllers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/internal/entities"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

// UserLoginHandler
//
//	@Summary		Login endpoint
//	@Description	login with username and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			Request	body		entities.UserRequest				true	"login body"
//	@Success		200		{object}	responses.LoginOKResponse			"Success"
//	@Failure		404		{object}	responses.UserNotExistsResponse		"Error"
//	@Failure		401		{object}	responses.PassNotEqualResponse		"Error"
//	@Failure		500		{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/auth/login [post]
func (u *UsersController) UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
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
	if err := u.Service.Authenticate(w, r, user, true); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, user, http.StatusOK)
}
