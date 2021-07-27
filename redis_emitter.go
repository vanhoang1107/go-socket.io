package socketio

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisEmitter struct {
	client *redis.Client

	namespace string
	uid       string
	key       string
}

func NewRedisEmitter(namespace string, opts *RedisAdapterOptions) (_ *RedisEmitter, err error) {
	if namespace == "/" {
		namespace = ""
	}
	var (
		client = redis.NewClient(&opts.Options)
		uid    = newV4UUID()
	)
	rbc := RedisEmitter{
		client:    client,
		uid:       uid,
		namespace: namespace,
		key:       fmt.Sprintf("%s#%s#%s", opts.Prefix, namespace, uid),
	}
	return &rbc, nil
}

func (re *RedisEmitter) Send(ctx context.Context, room, event string, args ...interface{}) error {
	return re.publish(ctx, room, event, args...)
}

func (re *RedisEmitter) SendAll(ctx context.Context, event string, args ...interface{}) error {
	return re.publish(ctx, "", event, args...)
}

func (re *RedisEmitter) publish(
	ctx context.Context,
	room string, event string, args ...interface{},
) error {
	message := map[string][]interface{}{
		"opts": {room, event},
		"args": args,
	}
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("socket.io emitter: encode message: %w", err)
	}
	result := re.client.Publish(ctx, re.key, messageJSON)
	if err := result.Err(); err != nil {
		return fmt.Errorf("socket.io emitter: send message: %w", err)
	}
	return nil
}
