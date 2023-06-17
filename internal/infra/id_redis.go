package infra

import (
	"github.com/redis/go-redis/v9"
)

func NewIdRedis(redisDsn string) (*redis.Client, error) {
	var options *redis.Options
	var err error

	if options, err = redis.ParseURL(redisDsn); err != nil {
		return &redis.Client{}, err
	}

	conn := redis.NewClient(options)
	return conn, nil
}
