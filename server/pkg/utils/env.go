package utils

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"strings"
)

func GetListEnv(envKey string) []string {
	cutBracket := strings.Trim(constans.ORIGINS, "[]")
	return strings.Split(cutBracket, ",")
}
