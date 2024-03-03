package domain

import (
	"net/http"
)

type WsRepository interface {
	ServeWs(w http.ResponseWriter, r *http.Request, userID, username string)
}
