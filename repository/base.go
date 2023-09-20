package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"go_grpc/lib"
	"go_grpc/lib/logger"
)

type Repository struct {
	db           *lib.Database
	jobPublisher *lib.Publisher
	smtpClient   *lib.SMTPClient
}

func NewRepository(db *lib.Database, jobPublisher *lib.Publisher, smtpClient *lib.SMTPClient) Repository {
	return Repository{db: db, jobPublisher: jobPublisher, smtpClient: smtpClient}
}

func (repo *Repository) Transaction(ctx context.Context, fn func(context.Context) error) error {
	trx := repo.db.Begin()

	ctx = context.WithValue(ctx, "Trx", &lib.Database{DB: trx})
	if err := fn(ctx); err != nil {
		trx.Rollback()
		return err
	}

	return trx.Commit().Error
}

func (repo *Repository) PublishJob(ctx context.Context, topic string, payload interface{}) error {
	return repo.jobPublisher.Publish(ctx, topic, payload)
}

func (repo *Repository) SendEmail(ctx context.Context, req lib.SMTPRequest) error {
	return repo.smtpClient.Send(ctx, req)
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func LogError(ctx context.Context, message string, err error) {
	logger.Error(ctx, message, map[string]interface{}{
		"error": err,
		"tags":  []string{"gorm"},
	})
}

func LogWarn(ctx context.Context, message string) {
	logger.Warn(ctx, message, map[string]interface{}{
		"tags": []string{"gorm"},
	})
}
