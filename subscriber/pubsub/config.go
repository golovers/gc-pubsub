package pubsub

import (
	"github.com/Sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
)

// Config contains all configurations to spin up a pubsub subscriber.
type Config struct {
	// Pubsub topic name which publisher publishes to
	Topic string `envconfig:"PUBSUB_TOPIC"`
	// Pubsub subscription name for subscriber
	Subscription string `envconfig:"PUBSUB_SUBSCRIPTION"`
	// Create pubsub subscription if not existed
	CreateSub bool `envconfig:"PUBSUB_CREATE_SUBSCRIPTION"`
}

// LoadEnvConfig returns a Config object populated from environment variables.
func LoadEnvConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		logrus.Fatalf("config: Unable to load config for %T: %s", cfg, err)
	}
	return cfg
}
