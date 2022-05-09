package event

import (
	"context"
	"fmt"
)

type Event struct {
}

func (*Event) Publish(ctx context.Context, topic string, tag string, data map[string]interface{}) error {
	fmt.Printf("publish event ctx : %v, topic: %s, tag: %s, data: %v", ctx, topic, tag, data)
	return nil
}
