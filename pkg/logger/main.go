package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func loadEncoderConfig(devMode bool, timeFormat string) zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()

	if devMode {
		cfg = zap.NewDevelopmentEncoderConfig()
	}

	cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(timeFormat))
	}

	return cfg
}

type zapLogger struct {
	zap *zap.SugaredLogger
}

func New(options Options) Logger {
	globalLogLevel := levelMap[options.LogLevel]

	encoderCfg := loadEncoderConfig(options.DevMode, options.DateTimeFormat)

	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(globalLogLevel),
		Development:   options.DevMode,
		Encoding:      "json",
		EncoderConfig: encoderCfg,
	}

	zapLog, err := loggerConfig.Build(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(zapLog)()

	return &zapLogger{
		zap: zapLog.Named(options.ProjectName).Sugar(),
	}
}

func (l *zapLogger) Debug(message string, fields ...zap.Field) {
	l.zap.Desugar().Debug(message, fields...)
}

func (l *zapLogger) Info(message string, fields ...zap.Field) {
	l.zap.Desugar().Info(message, fields...)
}

func (l *zapLogger) Warn(message string, fields ...zap.Field) {
	l.zap.Desugar().Warn(message, fields...)
}

func (l *zapLogger) Error(message string, fields ...zap.Field) {
	l.zap.Desugar().Error(message, fields...)
}

func (l *zapLogger) Fatal(message string, fields ...zap.Field) {
	l.zap.Desugar().Fatal(message, fields...)
}

func (l *zapLogger) Debugf(format string, a ...any) {
	l.zap.Debugf(format, a...)
}

func (l *zapLogger) Infof(format string, a ...any) {
	l.zap.Infof(format, a...)
}

func (l *zapLogger) Warnf(format string, a ...any) {
	l.zap.Warnf(format, a...)
}

func (l *zapLogger) Errorf(format string, a ...any) {
	l.zap.Errorf(format, a...)
}

func (l *zapLogger) Fatalf(format string, a ...any) {
	l.zap.Fatalf(format, a...)
}

func (l *zapLogger) Sync() error {
	return l.zap.Sync()
}
