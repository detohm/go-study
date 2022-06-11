package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "my-kafka-topic-multi"
	broker1Address = "localhost:9093"
	broker2Address = "localhost:9094"
	broker3Address = "localhost:9095"
)

func main() {
	runtime.GOMAXPROCS(4)
	ctx := context.Background()
	go produce(ctx)
	consume(ctx)
}

func produce(ctx context.Context) {

	l := log.New(os.Stdout, "kafka writer: ", 0)

	i := 0
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{
			broker1Address,
			broker2Address,
			broker3Address,
		},
		Topic:        topic,
		BatchSize:    10,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: 1,
		Logger:       l,
	})

	for {
		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message" + strconv.Itoa(i)),
		})
		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes:", i)
		i++
		time.Sleep(time.Second)
	}
}

func consume(ctx context.Context) {

	l := log.New(os.Stdout, "kafka reader: ", 0)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{
			broker1Address,
			broker2Address,
			broker3Address,
		},
		Topic:       topic,
		GroupID:     "my-group3",
		MinBytes:    20,
		MaxBytes:    1e6,
		MaxWait:     10 * time.Second,
		StartOffset: kafka.FirstOffset,
		Logger:      l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}
