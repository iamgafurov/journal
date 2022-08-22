package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: "./app.log",
		MaxSize:  100,
		MaxAge:   60,
	})

	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "datetime"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.DebugLevel,
	)

	Logger = zap.New(core)
}
