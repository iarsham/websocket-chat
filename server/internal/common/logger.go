package common

import (
	"github.com/iarsham/websocket-chat/pkg/constans"
	"go.uber.org/zap"
)

func ZapLogger() *zap.Logger {
	if constans.Mode {
		return zap.Must(zap.NewDevelopment())
	}
	return zap.Must(zap.NewProduction())
}
