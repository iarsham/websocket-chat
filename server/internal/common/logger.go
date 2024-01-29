package common

import (
	"go.uber.org/zap"
	"os"
	"strconv"
)

func ZapLogger() *zap.Logger {
	Mode, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if Mode {
		return zap.Must(zap.NewDevelopment())
	}
	return zap.Must(zap.NewProduction())
}
