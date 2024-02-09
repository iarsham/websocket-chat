package responses

import (
	"encoding/json"
	"github.com/iarsham/websocket-chat/internal/common"
	"github.com/iarsham/websocket-chat/pkg/constans"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}, status int) {
	var result []byte
	switch data.(type) {
	case string, int, float64:
		result, _ = json.Marshal(common.M{constans.Response: data})
	default:
		result, _ = json.Marshal(data)
	}
	w.Header().Set(constans.ContentType, constans.JsonContentType)
	w.WriteHeader(status)
	w.Write(result)
}
