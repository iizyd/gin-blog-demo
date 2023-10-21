package logger

import (
	"backend-go/internal/pkg/config"
	"fmt"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func InitLogger() {
	// defer sugarLogger.Sync()
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	rootPath, _ := os.Getwd()
	dir := path.Join(rootPath, config.Config.App.LogSavePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// 目录不存在，创建目录
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("无法创建目录:", err)
			os.Exit(1)
		}
	}

	lumberJackLogger := &lumberjack.Logger{
		Filename:   path.Join(dir, config.Config.App.LogFileName),
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func Debugf(s string, args ...any) {
	sugarLogger.Debugf(s, args...)
}

func Errorf(s string, args ...any) {
	sugarLogger.Errorf(s, args...)
}

func Infof(s string, args ...any) {
	sugarLogger.Infof(s, args...)
}
