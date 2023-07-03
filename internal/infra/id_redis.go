package infra

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

// NewIdRedis creates a new Redis client with the provided Redis DSN.
func NewIdRedis(redisDsn string) (*redis.Client, error) {
	var options *redis.Options
	var err error

	if options, err = redis.ParseURL(redisDsn); err != nil {
		fmt.Println(redisDsn, err)
		panic(err)
	}

	conn := redis.NewClient(options)
	return conn, nil
}
