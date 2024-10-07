package logger

import (
	"os"

	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/setting"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// the LoggerZap struct is embedding a pointer to a zap.Logger
// gives the embedding struct (in this case LoggerZap) all the methods and behavior of the embedded struct (in this case, *zap.Logger)
type LoggerZap struct {
	*zap.Logger
}

func NewLogger(config setting.LoggerSetting) *LoggerZap {
	logLevel := config.LogLevel
	// debug -> info -> warn -> error -> fatal -> panic
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}
	encoder := getEncoderLog()
	hook := lumberjack.Logger{
		Filename:   config.FileName, // "/var/log/myapp/foo.log",
		MaxSize:    config.MaxSize,  // megabytes
		MaxBackups: config.MaxBackups,
		MaxAge:     config.MaxAge,   //days
		Compress:   config.Compress, // disabled by default
	}
	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		level)
	return &LoggerZap{zap.New(core, zap.AddCaller())}
}

// format logs
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1727941184.9466841 -> 2024-10-03T14:39:44.946+0700
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> Time
	encoderConfig.TimeKey = "time"
	// info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:25"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
