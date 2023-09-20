package main

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	_ "time/tzdata"

	app "go_grpc"
	"go_grpc/config"
	"go_grpc/lib/logger"
)

func init() {
	godotenv.Load()
	logger.Init()
}

func main() {
	db, _ := config.NewPG()
	queue := config.NewQueue()
	smtpClient := config.NewSMTPClient()
	publisher := queue.NewPublisher()
	pubsub := config.NewPubSub()
	messaging := config.NewMessaging()
	storage := config.NewStorage()
	app := app.NewBackend(db, publisher, &smtpClient, pubsub, messaging, storage)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	time.Local = loc

	ctx := context.Background()
	ctx = context.WithValue(ctx, "X-Request-ID", uuid.NewString())

	logger.Info(ctx, "cron meeting starting", map[string]interface{}{
		"tags": []string{"cron"},
	})

	app.Usecase.BlastBeforeMeetingMessage(ctx, false)
	app.Usecase.BlastOnGoingMeetingMessage(ctx, false)

	logger.Info(ctx, "cron meeting ended", map[string]interface{}{
		"tags": []string{"cron"},
	})
}
