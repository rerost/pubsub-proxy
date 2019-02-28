package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/srvc/fail"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		fail.New("Please set GOOGLE_CLOUD_PROJECT to env")
	}

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		return fail.Wrap(err)
	}
	pubsubClient, err := pubsub.NewClient(ctx, projectID, option.WithGRPCConn(conn))
	if err != nil {
		return fail.Wrap(err)
	}

	topic := pubsubClient.Topic(os.Getenv("PUBSUB_TOPIC"))
	defer topic.Stop()

	result := topic.Publish(ctx, &pubsub.Message{Data: []byte("test")})
	_, err = result.Get(ctx)
	return fail.Wrap(err)
}
