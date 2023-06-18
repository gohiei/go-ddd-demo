package user

import (
	"context"

	"github.com/redis/go-redis/v9"

	repo "cypt/internal/user/repository"
)

var _ repo.IDRepository = (*RedisIDRepository)(nil)

type RedisIDRepository struct {
	ctx  context.Context
	conn *redis.Client
}

func NewRedisIDRepository(conn *redis.Client) *RedisIDRepository {
	return &RedisIDRepository{
		ctx:  context.Background(),
		conn: conn,
	}
}

func (r *RedisIDRepository) Incr(step int) (int64, error) {
	v := r.conn.IncrBy(r.ctx, "this.is.my.id", int64(step)).Val()

	return v, nil
}
