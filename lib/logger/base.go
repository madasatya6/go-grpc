package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	loggerOnce sync.Once
	logger     *zap.Logger
)

func Init() {
	loggerOnce.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.TimeKey = "@timestamp"
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
		logger, _ = config.Build(zap.AddCallerSkip(1))
	})
}

func Info(ctx context.Context, message string, messages map[string]interface{}) {
	if logger == nil {
		return
	}

	logger.Info(message, commonFields(ctx, message, messages)...)
}

func Error(ctx context.Context, message string, messages map[string]interface{}) {
	if logger == nil {
		return
	}

	logger.Error(message, commonFields(ctx, message, messages)...)
}

func Warn(ctx context.Context, message string, messages map[string]interface{}) {
	if logger == nil {
		return
	}

	logger.Warn(message, commonFields(ctx, message, messages)...)
}

func commonFields(ctx context.Context, message string, messages map[string]interface{}) []zap.Field {
	fields := []zap.Field{}

	requestID, ok := ctx.Value("X-Request-ID").(string)
	if ok {
		fields = append(fields, zap.String("request_id", requestID))
	}

	currentUserID, ok := ctx.Value("CurrentUserID").(uint)
	if ok {
		fields = append(fields, zap.Uint("actor_id", currentUserID))
	}

	for key, val := range messages {
		fields = append(fields, zap.Any(key, val))
	}

	return fields
}
