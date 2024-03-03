package controllers

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

// UserDeleteHandler
//
//	@Summary		Delete account endpoint
//	@Description	If there is a user session, the user can be deleted
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		204	{object}	responses.RecordDeletedResponse		"Success"
//	@Failure		500	{object}	responses.InterServerErrorResponse	"Error"
//	@Router			/users/delete-account [delete]
func (u *UsersController) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	userID := getFromCtx(r.Context(), constans.UserID).(string)
	if err := u.Service.DeleteUser(userID); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	if err := u.Service.Authenticate(w, r, nil, false); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, nil, http.StatusNoContent)
}
