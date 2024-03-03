package controllers

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"net/http"
)

func (ws *WsController) WsHandler(w http.ResponseWriter, r *http.Request) {
	userID := getFromCtx(r.Context(), constans.UserID).(string)
	userName := getFromCtx(r.Context(), constans.Username).(string)
	ws.Service.ServeWs(w, r, userID, userName)
}
