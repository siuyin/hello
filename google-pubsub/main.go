package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/iterator"
)

const projectID = "lsy0318"

func main() {
	//createTopic("junk1")

	//printSub()

	printTopics()
}

func printTopics() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	iter := client.Topics(ctx)
	for {
		t, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
	}
}

func createSub(w io.Writer, subID string, topic *pubsub.Topic) error {
	// projectID := "my-project-id"
	// subID := "my-sub"
	// topic of type https://godoc.org/cloud.google.com/go/pubsub#Topic
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub, err := client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic:       topic,
		AckDeadline: 20 * time.Second,
	})
	if err != nil {
		return fmt.Errorf("CreateSubscription: %w", err)
	}
	fmt.Fprintf(w, "Created subscription: %v\n", sub)
	return nil
}

func printSub() {
	subs, err := listSub()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range subs {
		fmt.Printf("%v\n", s)
	}
}

func listSub() ([]*pubsub.Subscription, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	var subs []*pubsub.Subscription
	it := client.Subscriptions(ctx)
	for {
		s, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Next: %w", err)
		}
		subs = append(subs, s)
	}
	return subs, nil
}

func createTopic(name string) {
	ctx := context.Background()

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the id for the new topic.
	topicID := name

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicID)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}
