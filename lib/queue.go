package lib

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"

	"go_grpc/lib/logger"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

type Context struct{}
type Queue struct {
	pool *redis.Pool
}

func NewQueue(pool *redis.Pool) *Queue {
	return &Queue{pool: pool}
}

func (queue *Queue) NewWorker() *Worker {
	pool := work.NewWorkerPool(Context{}, 10, "app", queue.pool)
	return &Worker{pool: pool}
}

func (queue *Queue) NewPublisher() *Publisher {
	enqueuer := work.NewEnqueuer("app", queue.pool)
	return &Publisher{enqueuer: enqueuer}
}

type Worker struct {
	pool *work.WorkerPool
}

func (worker *Worker) Register(topic string, fn func(ctx context.Context, payload []byte) error) {
	worker.pool.Job(topic, func(job *work.Job) error {
		return fn(context.Background(), []byte(job.Args["payload"].(string)))
	})
}

func (worker *Worker) Start() {
	worker.pool.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	worker.pool.Stop()
}

type Publisher struct {
	enqueuer *work.Enqueuer
}

func (publisher *Publisher) Publish(ctx context.Context, topic string, payload interface{}) error {
	logger.Info(ctx, "publish job", map[string]interface{}{
		"topic":   topic,
		"payload": payload,
		"tags":    []string{"queue", "publish"},
	})

	encoded, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = publisher.enqueuer.Enqueue(topic, work.Q{"payload": string(encoded)})
	if err != nil {
		logger.Error(ctx, "failed publish job", map[string]interface{}{
			"error":   err,
			"topic":   topic,
			"payload": string(encoded),
			"tags":    []string{"queue", "publish"},
		})

		return err
	}

	return nil
}
