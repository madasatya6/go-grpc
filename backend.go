package backend

import (
	"go_grpc/lib"
	"go_grpc/pubsub"
	"go_grpc/repository"
	"go_grpc/storage"
	"go_grpc/usecase"
)

type Backend struct {
	Usecase *usecase.Usecase
}

func NewBackend(db *lib.Database, jobPub *lib.Publisher, smtpClient *lib.SMTPClient, pubsubLib *lib.PubSub, messaging *lib.Messaging, storage storage.Storage) Backend {
	repository := repository.NewRepository(db, jobPub, smtpClient)
	publisher := &pubsub.Publisher{Client: pubsubLib}
	usecase := usecase.NewUsecase(&repository, storage, publisher, messaging)

	return Backend{
		Usecase: &usecase,
	}
}
