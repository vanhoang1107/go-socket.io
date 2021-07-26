package socketio

import (
	"github.com/go-redis/redis/v8"
)

// RedisAdapterOptions is configuration to create new adapter
type RedisAdapterOptions struct {
	redis.Options
	Prefix string
}

func defaultOptions() *RedisAdapterOptions {
	return &RedisAdapterOptions{
		Options: redis.Options{
			Addr:    "127.0.0.1:6379",
			Network: "tcp",
		},
		Prefix: "socket.io",
	}
}

func getOptions(opts *RedisAdapterOptions) *RedisAdapterOptions {
	options := defaultOptions()
	if opts != nil {
		if opts.Addr != "" {
			options.Addr = opts.Addr
		}
		if opts.Prefix != "" {
			options.Prefix = opts.Prefix
		}
		if opts.Network != "" {
			options.Network = opts.Network
		}
	}
	return options
}
