package pubsub

import (
	"context"

	"go_grpc/lib"
)

type Publisher struct {
	Client *lib.PubSub
}

func (p *Publisher) Publish(ctx context.Context, topic, meetingID string) error {
	err := p.Client.Redis.Send("PUBLISH", topic, meetingID)
	if err != nil {
		return err
	}

	return p.Client.Redis.Flush()
}
