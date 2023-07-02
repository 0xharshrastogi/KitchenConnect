package logger

import (
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *zap.Logger {
	l, _ := zap.NewDevelopment()
	return l
}
