package pubsub

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/Sirupsen/logrus"
)

// Config contains configurations to spin up a pubsub publisher.
type Config struct {
	// Pubsub topic name which publisher publishes to
	Topic string `envconfig:"PUBSUB_TOPIC"`
	// Create pubsub topic if not existed
	CreateTopic bool `envconfig:"PUBSUB_CREATE_TOPIC"`
}

// LoadEnvConfig returns a Config object populated from environment variables.
func LoadEnvConfig() *Config {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		logrus.Fatalf("config: Unable to load config for %T: %s", cfg, err)
	}
	return cfg
}
