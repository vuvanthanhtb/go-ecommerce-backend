package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello, %s - %d", "Sugar", 2025)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Sugar"), zap.Int("year", 2025))

	// 2. level
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3. custom
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Info("Info log", zap.Int("line", 2))
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1749842308.301941 -> 2025-06-14T02:18:28.301+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> time
	encoderConfig.TimeKey = "time"

	// from info INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:22"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
