package pubsub

import (
	"cloud.google.com/go/pubsub"
	"fmt"
	"github.com/rs/xid"
	"golang.org/x/net/context"
	"time"
)

// Publisher that publishes messages to a Google Pub/Sub Topic.
type Publisher struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

// New creates and returns a new instance of a GCP Publisher.
func NewPublisher(cfg *Config) (*Publisher, error) {
	// Pubsub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cfg.Topic)
	if err != nil {
		return nil, err
	}

	// Load topic
	t := client.Topic(cfg.Topic)
	ok, err := t.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ok {
		if !cfg.CreateTopic {
			return nil, fmt.Errorf("gcp: topic '%s' does not exist", cfg.Topic)
		}
		t, err = client.CreateTopic(ctx, cfg.Topic)
		if err != nil {
			return nil, fmt.Errorf("gcp: error creating topic '%s': %v", cfg.Topic, err)
		}
	}
	return &Publisher{
		client: client,
		topic:  t,
	}, nil
}

// Publish message to the pubsub topic.
func (p *Publisher) Publish(ctx context.Context, data []byte) error {
	msg := &pubsub.Message{
		Attributes: map[string]string{
			"timestamp": time.Now().UTC().String(),
			"event_id":  xid.New().String(),
		},
		Data: data,
	}
	p.topic.Publish(ctx, msg)
	return nil // TODO
}

// Close pubsub client.
func (p *Publisher) Close() error {
	return p.client.Close()
}
