package worker

import (
	"context"
	"encoding/json"

	"go_grpc/lib/logger"
	"go_grpc/model"
)

func (worker *Worker) SendVerificationNotification(ctx context.Context, payload []byte) error {
	logger.Info(ctx, "acking send verification notification", map[string]interface{}{
		"payload": string(payload),
		"tags":    []string{"smtp"},
	})

	var verification model.Verification

	err := json.Unmarshal(payload, &verification)
	if err != nil {
		return err
	}

	err = worker.Backend.Usecase.AuthSendVerificationNotification(ctx, verification)
	if err != nil {
		return err
	}

	logger.Info(ctx, "success sent verification notification", map[string]interface{}{
		"payload": string(payload),
		"tags":    []string{"smtp"},
	})

	return nil
}
