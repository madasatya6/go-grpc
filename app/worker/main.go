package main

import (
	"time"

	"github.com/joho/godotenv"

	_ "time/tzdata"

	app "go_grpc"
	"go_grpc/config"
	"go_grpc/lib/logger"
	"go_grpc/worker"
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

	worker := worker.NewWorker(&app)

	queueWorker := queue.NewWorker()
	queueWorker.Register("send-verification-notification", worker.SendVerificationNotification)
	queueWorker.Register("send-meeting-invitation", worker.SendMeetingInvitation)
	queueWorker.Start()
}
