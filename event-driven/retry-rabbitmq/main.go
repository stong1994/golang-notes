package main

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

const (
	deliverCnt = "x-delivery-count"
	amqpURI    = "amqp://guest:guest@localhost:5672/"
)

var (
	logger        = watermill.NewStdLogger(false, false)
	amqpConfig    = config()
	ErrGoToPoison = errors.New("go to posion")
)

func main() {
	pub := publisher()
	sub := subscriber()
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddMiddleware(poisonFilter(pub))
	router.AddMiddleware(recordDeliverCnt)
	router.AddHandler(
		"simple handler",
		"retry-queue",
		sub,
		"retry-queue",
		pub,
		func(msg *message.Message) ([]*message.Message, error) {
			logger.Info("deliver cnt", map[string]any{
				"cnt": msg.Metadata.Get(deliverCnt),
			})
			time.Sleep(time.Second)
			if msg.Metadata.Get(deliverCnt) > "3" { // if has deliverd more than 3 times, go to poison
				return nil, ErrGoToPoison
			}
			return []*message.Message{msg}, nil
		})

	go router.Run(context.Background())
	<-router.Running()

	err = pub.Publish("retry-queue", message.NewMessage(watermill.NewShortUUID(), []byte("hi")))
	if err != nil {
		panic(err)
	}
	go listenPoison(sub)
	time.Sleep(time.Second * 10)
}

func publisher() message.Publisher {
	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}
	return publisher
}

func subscriber() message.Subscriber {
	subscriber, err := amqp.NewSubscriber(amqpConfig, logger)
	if err != nil {
		panic(err)
	}
	return subscriber
}

func config() amqp.Config {
	config := amqp.NewDurableQueueConfig(amqpURI)
	config.Consume.NoRequeueOnNack = false
	return config
}

func recordDeliverCnt(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		cnt := msg.Metadata.Get(deliverCnt)
		if cnt == "" {
			cnt = "1"
		} else {
			c, _ := strconv.Atoi(cnt)
			cnt = strconv.Itoa(c + 1)
		}
		msg.Metadata.Set(deliverCnt, cnt)
		return h(msg)
	}
}

func poisonFilter(pub message.Publisher) message.HandlerMiddleware {
	mid, err := middleware.PoisonQueueWithFilter(pub, "poison-queue", func(err error) bool {
		logger.Info("going to the poison", nil)
		return errors.Is(err, ErrGoToPoison)
	})
	if err != nil {
		panic(err)
	}
	return mid
}

func listenPoison(sub message.Subscriber) {
	messages, err := sub.Subscribe(context.Background(), "poison-queue")
	if err != nil {
		panic(err)
	}
	for msg := range messages {
		logger.Info("got poison message", map[string]any{
			"deliver cnt": msg.Metadata.Get(deliverCnt),
		})
	}
}
