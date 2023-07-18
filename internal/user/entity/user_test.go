package entity_test

import (
	"testing"

	"cypt/internal/dddcore"
	"cypt/internal/user/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u, err := entity.NewUser("test1", "password1", 2)

	assert.Nil(t, err)
	assert.Equal(t, "test1", u.GetUsername())
	assert.Equal(t, "password1", u.GetPassword())
	assert.Equal(t, int64(2), u.GetUserID())
	assert.NotNil(t, u.GetID())
}

func TestBuildNewUser(t *testing.T) {
	uuid := dddcore.NewUUID().String()
	u := entity.BuildUser(uuid, "test2", "password2", 3)

	assert.Equal(t, "test2", u.GetUsername())
	assert.Equal(t, "password2", u.GetPassword())
	assert.Equal(t, int64(3), u.GetUserID())
	assert.NotNil(t, u.GetID())
	assert.Equal(t, uuid, u.GetID().String())
}

func TestRename(t *testing.T) {
	u, _ := entity.NewUser("test3", "password3", 5)

	u.Rename("test4")
	assert.Equal(t, "test4", u.GetUsername())
}
