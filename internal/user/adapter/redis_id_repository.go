package user

import (
	"context"

	"github.com/redis/go-redis/v9"

	repo "cypt/internal/user/repository"
)

var _ repo.IDRepository = (*RedisIDRepository)(nil)

// RedisIDRepository is an implementation of IDRepository using Redis as the underlying storage.
type RedisIDRepository struct {
	ctx  context.Context
	conn *redis.Client
}

// NewRedisIDRepository creates a new instance of RedisIDRepository.
func NewRedisIDRepository(conn *redis.Client) *RedisIDRepository {
	return &RedisIDRepository{
		ctx:  context.Background(),
		conn: conn,
	}
}

// Incr increments the value stored in Redis with the specified step and returns the updated value.
func (r *RedisIDRepository) Incr(step int) (int64, error) {
	v := r.conn.IncrBy(r.ctx, "this.is.my.id", int64(step)).Val()

	return v, nil
}
