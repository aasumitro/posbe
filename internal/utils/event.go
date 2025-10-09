package utils

import (
	"context"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	mu sync.Mutex

	redisPublisher  *redis.Client
	redisSubscriber *redis.Client
	subscriptions   map[string]*redis.PubSub
)

func InitEventStream(rp, rs *redis.Client) {
	if rp == nil || rs == nil {
		panic("Redis PubSub connection is required")
	}
	redisPublisher = rp
	redisSubscriber = rs
	subscriptions = make(map[string]*redis.PubSub)
}

// PublishEvent publishes a message to a specified channel in Redis pub/sub.
// Returns an error if the publish operation fails.
//
// usage:
//
// import "github.com/aasumitro/posbe/internal/utils/event"
//
//	if err := event.Publish(ctx, "notify", []byte("test")); err != nil {
//		return err
//	}
func PublishEvent(ctx context.Context, channel string, message interface{}) error {
	return redisPublisher.Publish(ctx, channel, message).Err()
}

// SubscribeEvent to messages on a specified channel in Redis pub/sub.
// It invokes the provided handler function for each received message asynchronously.
// Returns the PubSub instance for managing the subscription.
//
// usage:
//
//	 import "github.com/aasumitro/posbe/internal/utils/event"
//
//		event.Subscribe(ctx, "notify", func(message *redis.Message) {
//			fmt.Println(message.Payload)
//		})
func SubscribeEvent(ctx context.Context, channel string, handler func(*redis.Message)) {
	subscription := redisSubscriber.Subscribe(ctx, channel)

	mu.Lock()
	subscriptions[channel] = subscription
	mu.Unlock()

	ch := subscription.Channel()
	for {
		select {
		case <-ctx.Done():
			_ = subscription.Unsubscribe(
				context.Background(), channel)
			_ = subscription.Close()

			mu.Lock()
			delete(subscriptions, channel)
			mu.Unlock()
			
			return
		case msg, ok := <-ch:
			if !ok {
				return
			}
			handler(msg)
		}
	}
}

func CloseEventStream(ctx context.Context) {
	mu.Lock()
	for channel, pb := range subscriptions {
		delete(subscriptions, channel)
		if err := pb.Unsubscribe(ctx, channel); err != nil {
			log.Printf("Error unsubscribing from channel %s during cleanup: %v", channel, err)
		}
		if err := pb.Close(); err != nil {
			log.Printf("Error closing PubSub for channel %s during cleanup: %v", channel, err)
		}
	}
	mu.Unlock()
}
