package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var lg *zap.Logger

type Config struct {
	Encoding string
	Level    string
	DevMode  bool
	Caller   bool
}
type appLogger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func GetSugar() *zap.SugaredLogger {
	return logger
}

func GetGlobalLogger() *zap.Logger {
	return lg
}

func (l *appLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func WithDefault(appName string) {
	(&Config{
		Encoding: "json",
		Level:    "info",
		DevMode:  false,
		Caller:   false,
	}).SetConfiguration(appName)
}

func (config *Config) SetConfiguration(appName string) {
	appLogger := &appLogger{level: config.Level, devMode: config.DevMode, encoding: config.Encoding}
	logLevel := appLogger.getLoggerLevel()
	logWriter := zapcore.AddSync(os.Stdout)

	var encoderCfg zapcore.EncoderConfig
	if appLogger.devMode {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.NameKey = "service"
	encoderCfg.TimeKey = "time"
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "line"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	if appLogger.encoding == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoderCfg.FunctionKey = "caller"
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))

	lg = zap.New(core)
	if config.Caller {
		lg = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	}

	logger = lg.Sugar()
	logger = logger.Named(appName)
	lg = lg.Named(appName)
	logger.Sync()
	lg.Sync()
}

func Debug(args ...any) {
	logger.Debug(args...)
}

func Info(args ...any) {
	logger.Info(args...)
}

func Warn(args ...any) {
	logger.Warn(args...)
}

func Error(args ...any) {
	logger.Error(args...)
}

func DPanic(args ...any) {
	logger.DPanic(args...)
}

func Panic(args ...any) {
	logger.Panic(args...)
}

func Fatal(args ...any) {
	logger.Fatal(args...)
}

func DebugF(template string, args ...any) {
	logger.Debugf(template, args...)
}

func InfoF(template string, args ...any) {
	logger.Infof(template, args...)
}

func WarnF(template string, args ...any) {
	logger.Warnf(template, args...)
}

func ErrorF(template string, args ...any) {
	logger.Errorf(template, args...)
}

func DPanicF(template string, args ...any) {
	logger.DPanicf(template, args...)
}

func PanicF(template string, args ...any) {
	logger.Panicf(template, args...)
}

func FatalF(template string, args ...any) {
	logger.Fatalf(template, args...)
}

func DebugW(msg string, keysAndValues ...any) {
	logger.Debugw(msg, keysAndValues)
}

func InfoW(msg string, keysAndValues ...any) {
	logger.Info(msg, keysAndValues)
}

func WarnW(msg string, keysAndValues ...any) {
	logger.Warnw(msg, keysAndValues)
}

func ErrorW(msg string, keysAndValues ...any) {
	logger.Errorw(msg, keysAndValues)
}

func DPanicW(msg string, keysAndValues ...any) {
	logger.DPanicw(msg, keysAndValues)
}

func PanicW(msg string, keysAndValues ...any) {
	logger.Info(msg, keysAndValues)
}

func FatalW(msg string, keysAndValues ...any) {
	logger.Fatal(msg, keysAndValues)
}
