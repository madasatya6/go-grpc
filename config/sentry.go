package config

import (
	"context"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"go_grpc/lib/logger"
)

func InitSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         os.Getenv("SENTRY_DSN"),
		Environment: os.Getenv("SENTRY_ENV"),
	})
	if err != nil {
		logger.Error(context.Background(), "Error init sentry", map[string]interface{}{
			"error": err,
			"tags":  []string{"sentry"},
		})
	}
	defer sentry.Flush(2 * time.Second)
	sentry.CaptureMessage("It works!")

	defer func() {
		err := recover()

		if err != nil {
			sentry.CurrentHub().Recover(err)
			sentry.Flush(time.Second * 5)
		}
	}()
}
