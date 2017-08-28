package main

import (
	"cloud.google.com/go/pubsub"
	"github.com/Sirupsen/logrus"
	ps "github.com/lnquy/gc-pubsub/subscriber/pubsub"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Configuration
	cfg := ps.LoadEnvConfig()

	// Create pubsub subscriber
	subscriber, err := ps.NewSubscriber(cfg)
	if err != nil {
		logrus.Fatalf("main: unable to create pubsub subscriber: %v", err)
	}

	// Exit channel listens on os.Signal
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt, os.Kill)

	// Receive message from pubsub topic
	go func() {
		err := subscriber.Sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
			logrus.Info("Got message: %q\n", string(msg.Data))
			msg.Ack()
		})
		if err != nil {
			logrus.Errorf("pubsub: Error when receiving message: %s", err.Error())
		}
	}()

	<-exit
	logrus.Info("Exiting...")
}
