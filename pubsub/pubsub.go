package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err)
		return
	}
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	topicID := os.Getenv("GOOGLE_PUBSUB_TOPIC_ID")
	t := client.Topic(topicID)

	msg := "Hello: " + fmt.Sprintf("%d", time.Now().UnixMicro())
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	id, err := result.Get(ctx) // get blocked
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Published a message: msg:%s ID:%v\n", msg, id)

	subscriptionID := os.Getenv("GOOGLE_PUBSUB_SUBSCRIPTION_ID")
	sub := client.Subscription(subscriptionID)

	// cctx, cancel := context.WithCancel(ctx)
	// f will be called concurrently with multiple go routines
	// code below will be block forever as passing background context
	// without deadline or cancel
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
		msg.Ack()
	})

	/* the code below will terminate once received 10 messages
	cctx, cancel := context.WithCancel(ctx)
	count := 0
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
		msg.Ack()
		count++
		if count == 5 {
			cancel() // calling context cancel function
		}
	})
	*/
	if err != nil {
		fmt.Println("Received error:", err)
		return
	}
	fmt.Println("done.")
}
