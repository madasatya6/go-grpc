package pubsub

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"go_grpc/lib"
	"go_grpc/lib/logger"
)

type Subscriber struct {
	Client *lib.PubSub
}

func (s *Subscriber) Subscribe(ctx context.Context, topic string, message chan<- string) error {
	psc := redis.PubSubConn{Conn: s.Client.Redis}
	err := psc.Subscribe(topic)
	if err != nil {
		return err
	}

	go func() {
		for {
			switch v := psc.Receive().(type) {
			case redis.Message:
				message <- string(v.Data)
			case error:
				logger.Error(context.Background(), "error subscription", map[string]interface{}{
					"error": v,
					"tags":  []string{"websocket"},
				})
			}
		}
	}()

	return nil
}
