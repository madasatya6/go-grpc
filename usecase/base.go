package usecase

import (
	"go_grpc/lib"
	"go_grpc/pubsub"
	"go_grpc/repository"
	"go_grpc/storage"
)

type Usecase struct {
	repo      *repository.Repository
	storage   storage.Storage
	publisher *pubsub.Publisher
	messaging *lib.Messaging
}

func NewUsecase(repo *repository.Repository, storage storage.Storage, publisher *pubsub.Publisher, messaging *lib.Messaging) Usecase {
	return Usecase{repo: repo, storage: storage, publisher: publisher, messaging: messaging}
}
