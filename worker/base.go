package worker

import (
	"context"

	"go_grpc"
)

type Worker struct {
	Backend *backend.Backend
}

func NewWorker(backend *backend.Backend) Worker {
	return Worker{
		Backend: backend,
	}
}

func (worker *Worker) Example(ctx context.Context, payload []byte) error {
	return nil
}
