package main

import (
	"log"
	"net"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "time/tzdata"

	app "go_grpc"
	"go_grpc/config"
	h "go_grpc/handler"
	"go_grpc/lib/logger"
	"go_grpc/model/proto/handler"
)

func init() {
	godotenv.Load()
	logger.Init()
}

func main() {
	db, _ := config.NewPG()
	queue := config.NewQueue()
	publisher := queue.NewPublisher()
	smtpClient := config.NewSMTPClient()
	pubsub := config.NewPubSub()
	messaging := config.NewMessaging()
	storage := config.NewStorage()
	app := app.NewBackend(db, publisher, &smtpClient, pubsub, messaging, storage)
	handlerService := h.NewHandler(&app)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	time.Local = loc

	config.InitSentry()

	// Create TCP Server on localhost:8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}

	// Create new gRPC server handler
	server := grpc.NewServer(
		grpc.UnaryInterceptor(handlerService.UnaryServerInterceptor),
		grpc.StreamInterceptor(handlerService.StreamServerInterceptor),
	)

	// register gRPC UserService to gRPC server handler
	handler.RegisterHandlerServiceServer(server, &handlerService)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("error serve: %v", err)
	}
}
