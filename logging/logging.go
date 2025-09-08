package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	config2 "nso/config"
	"os"
)

var Logger *zap.Logger

var Sugar *zap.SugaredLogger

func init() {
	var err error
	cfg := config2.GetAppConfig()
	if cfg.Debug {
		Logger, err = zap.NewDevelopment()
		//encoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
		//core := zapcore.NewCore(encoder, os.Stdout, zapcore.DebugLevel)
		//Logger = zap.New(core)
		//Sugar = Logger.Sugar()
	} else {
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		consoleEncoder := zapcore.NewConsoleEncoder(config)
		fileEncoder := zapcore.NewJSONEncoder(config)
		logFile, _ := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		writer := zapcore.AddSync(logFile)
		defaultLogLevel := zapcore.DebugLevel
		core := zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel))
		Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	}
	if err != nil {
		log.Panicln(err)
	}
}
