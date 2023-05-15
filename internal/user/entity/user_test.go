package user_test

import (
	"cypt/internal/dddcore"
	user "cypt/internal/user/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	u, err := user.NewUser("test1", "password1")

	assert.Nil(t, err)
	assert.Equal(t, "test1", u.GetUsername())
	assert.Equal(t, "password1", u.GetPassword())
	assert.NotNil(t, u.GetId())
}

func TestBuildNewUser(t *testing.T) {
	uuid := dddcore.NewUUID().String()
	u := user.BuildUser(uuid, "test2", "password2")

	assert.Equal(t, "test2", u.GetUsername())
	assert.Equal(t, "password2", u.GetPassword())
	assert.NotNil(t, u.GetId())
	assert.Equal(t, uuid, u.GetId().String())
}

func TestRename(t *testing.T) {
	u, _ := user.NewUser("test3", "password3")

	u.Rename("test4")
	assert.Equal(t, "test4", u.GetUsername())
}
