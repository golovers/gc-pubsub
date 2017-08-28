package main

import (
	"context"
	"fmt"
	"github.com/Sirupsen/logrus"
	ps "github.com/lnquy/gc-pubsub/publisher/pubsub"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Configuration
	cfg := ps.LoadEnvConfig()

	// Create pubsub publisher
	publisher, err := ps.NewPublisher(cfg)
	if err != nil {
		logrus.Fatalf("main: unable to create pubsub publisher: %v", err)
	}

	// Exit channel listens on os.Signal
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)

	// Send message to pubsub topic every 10 seconds
	ticker := time.NewTicker(10 * time.Second)
	counter := 1
	for {
		select {
		case <-ticker.C:
			s := fmt.Sprintf("Publish message #%d at %s", counter, time.Now().UTC().String())
			err := publisher.Publish(context.Background(), []byte(s))
			if err != nil {
				logrus.Errorf("pubsub: Failed to publish message: %s", err.Error())
			} else {
				logrus.Info(s)
			}
			counter++
		case <-exit:
			logrus.Info("Exiting...")
			return
		}
	}
}
