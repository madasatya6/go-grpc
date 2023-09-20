package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "time/tzdata"

	app "go_grpc"
	"go_grpc/config"
	"go_grpc/lib"
	"go_grpc/lib/logger"
	"go_grpc/pubsub"
	"go_grpc/wshandler"
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
	pubsubLib := config.NewPubSub()
	messaging := config.NewMessaging()
	storage := config.NewStorage()
	app := app.NewBackend(db, publisher, &smtpClient, pubsubLib, messaging, storage)

	loc, _ := time.LoadLocation("Asia/Jakarta")

	time.Local = loc

	messageChan := make(chan string, 256)

	subscriber := pubsub.Subscriber{Client: pubsubLib}

	subscriber.Subscribe(context.Background(), lib.WebsocketRedisTopic, messageChan)

	hub := wshandler.NewHub(&app)

	handler := wshandler.NewWSHandler(&app, hub)
	go hub.Run()
	go handler.EventBroadcaster(messageChan)

	router := mux.NewRouter()
	router.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(rw, "ok")
	}).Methods("GET")
	router.HandleFunc("/ws/rooms/{id}", handler.PanicMiddlewares(http.HandlerFunc(handler.ServeWS)).ServeHTTP)

	srv := &http.Server{
		Addr:    "0.0.0.0:8081",
		Handler: router,
	}

	srv.ListenAndServe()
}
