package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"cypt/internal/dddcore"
	user "cypt/internal/user/entity"
)

func TestNewUser(t *testing.T) {
	u, err := user.NewUser("test1", "password1")

	assert.Nil(t, err)
	assert.Equal(t, "test1", u.GetUsername())
	assert.Equal(t, "password1", u.GetPassword())
	assert.NotNil(t, u.GetID())
}

func TestBuildNewUser(t *testing.T) {
	uuid := dddcore.NewUUID().String()
	u := user.BuildUser(uuid, "test2", "password2")

	assert.Equal(t, "test2", u.GetUsername())
	assert.Equal(t, "password2", u.GetPassword())
	assert.NotNil(t, u.GetID())
	assert.Equal(t, uuid, u.GetID().String())
}

func TestRename(t *testing.T) {
	u, _ := user.NewUser("test3", "password3")

	u.Rename("test4")
	assert.Equal(t, "test4", u.GetUsername())
}
