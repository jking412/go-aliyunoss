package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"time"
)

var Logger *zap.Logger

func InitLogger(logPath string, maxSize, maxBackups, maxAge int,
	compress bool, level string) {

	logWriter := getLogWriter(logPath, maxSize, maxBackups, maxAge, compress)
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		log.Println("log level error")
	}

	core := zapcore.NewCore(getEncoder(), logWriter, logLevel)

	Logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	zap.ReplaceGlobals(Logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getLogWriter(logPath string, maxSize, maxBackups, maxAge int,
	compress bool) zapcore.WriteSyncer {
	logName := time.Now().Format("2006-01-02") + ".log"

	lumberjackLogger := &lumberjack.Logger{
		Filename:   logPath + logName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
}

func Error(module string, fields ...zap.Field) {
	Logger.Error(module, fields...)
}

func Info(module string, fields ...zap.Field) {
	Logger.Info(module, fields...)
}

func Debug(module string, fields ...zap.Field) {
	Logger.Debug(module, fields...)
}

func Warn(module string, fields ...zap.Field) {
	Logger.Warn(module, fields...)
}
