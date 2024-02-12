package controllers

import (
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

// UserLogOutHandler
//	@Summary		Logout account endpoint
//	@Description	If there is a user session, the user can be logout
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.UserLogOutResponse		"Success"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/users/logout [post]
func (u *UsersController) UserLogOutHandler(w http.ResponseWriter, r *http.Request) {
	if err := u.Service.Authenticate(w, r, uuid.UUID{}, false); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, constans.LoggedOut, http.StatusOK)
}
