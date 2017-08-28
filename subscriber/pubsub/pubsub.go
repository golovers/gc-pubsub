package pubsub

import (
	"cloud.google.com/go/pubsub"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

// Subscriber that receives messages from a Google Pub/Sub Topic.
type Subscriber struct {
	Client *pubsub.Client
	Sub    *pubsub.Subscription
}

// New creates and returns a new instance of a pubsub subscriber.
func NewSubscriber(cfg *Config) (*Subscriber, error) {
	// Pubsub Client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, cfg.Topic)
	if err != nil {
		return nil, err
	}

	// Load topic
	t := client.Topic(cfg.Topic)
	_, err = t.Exists(ctx)
	if err != nil {
		return nil, err
	}

	// Load subscription
	sub := client.Subscription(cfg.Subscription)
	ok, err := sub.Exists(context.Background())
	if err != nil {
		logrus.Panicf("pubsub: subscription exists call failed: %v", err)
	}
	if !ok {
		if cfg.CreateSub {
			sub, err = client.CreateSubscription(context.Background(), cfg.Subscription, client.Topic(cfg.Topic), time.Second*20, &pubsub.PushConfig{})
			if err != nil {
				logrus.Panicf("pubsub: failed to create %s subscription", cfg.Subscription)
			}
		} else {
			logrus.Panicf("pubsub: subscription '%s' does not exists", cfg.Subscription)
		}
	}

	return &Subscriber{
		Client: client,
		Sub:    sub,
	}, nil
}

// Close pubsub client.
func (p *Subscriber) Close() error {
	return p.Client.Close()
}
