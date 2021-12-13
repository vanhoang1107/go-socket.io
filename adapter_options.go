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
	defaultOpts := defaultOptions()
	if opts == nil {
		return defaultOpts
	}
	if opts.Addr == "" {
		opts.Addr = defaultOpts.Addr
	}
	if opts.Prefix == "" {
		opts.Prefix = defaultOpts.Prefix
	}
	if opts.Network == "" {
		opts.Network = defaultOpts.Network
	}
	return opts
}
