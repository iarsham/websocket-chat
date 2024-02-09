package controllers

import (
	"github.com/google/uuid"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"github.com/iarsham/websocket-chat/pkg/responses"
	"net/http"
)

func (u *UsersController) UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	userID := u.getFromCtx(r.Context(), constans.UserID).(string)
	if err := u.Service.DeleteUser(userID); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	if err := u.Service.Authenticate(w, r, uuid.UUID{}, false); err != nil {
		responses.Json(w, constans.InternalError, http.StatusInternalServerError)
		return
	}
	responses.Json(w, nil, http.StatusNoContent)
}
