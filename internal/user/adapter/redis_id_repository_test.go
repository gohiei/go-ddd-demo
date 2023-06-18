package user_test

import (
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"

	repo "cypt/internal/user/adapter"
)

func TestIncr(t *testing.T) {
	conn, clientMock := redismock.NewClientMock()

	ret := int64(3)

	clientMock.ExpectIncrBy("this.is.my.id", 1).SetVal(ret)

	r := repo.NewRedisIDRepository(conn)
	id, _ := r.Incr(1)

	assert.Equal(t, ret, id)
}
