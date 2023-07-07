package user_test

import (
	"errors"
	"testing"

	"cypt/internal/dddcore"
	adapter "cypt/internal/user/adapter"
	entity "cypt/internal/user/entity"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	UserSchema = []string{"id", "username", "password", "user_id"}
)

func InitializeDatabase() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormdb, _ := gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})

	return gormdb, mock
}

func TestAdapterGet(t *testing.T) {
	db, mock := InitializeDatabase()
	uuid := dddcore.NewUUID()

	rows := sqlmock.NewRows(UserSchema)
	rows = rows.AddRow(uuid.String(), "test1", "password1", 1)

	mock.ExpectQuery("SELECT").WithArgs(uuid.String()).WillReturnRows(rows)

	r := adapter.NewMySQLUserRepository(db)
	u, err := r.Get(uuid)

	assert.Nil(t, err)
	assert.Equal(t, uuid, u.GetID())
	assert.Equal(t, "test1", u.GetUsername())
	assert.Equal(t, "password1", u.GetPassword())
}

func TestAdapterGetWithDatabaseError(t *testing.T) {
	db, mock := InitializeDatabase()
	uuid, _ := dddcore.BuildUUID("7f8394c4-4267-41c7-a01b-66b45c656bff")

	mock.ExpectQuery("SELECT").WithArgs(uuid.String()).WillReturnError(errors.New("fake error"))

	r := adapter.NewMySQLUserRepository(db)
	u, err := r.Get(uuid)

	assert.NotNil(t, err)
	assert.Equal(t, "failed to get (10003)", err.Error())
	assert.Empty(t, u.GetUsername())
}

func TestAdapterGetWithErrUserNotFound(t *testing.T) {
	db, mock := InitializeDatabase()
	uuid := dddcore.NewUUID()

	rows := sqlmock.NewRows(UserSchema)
	mock.ExpectQuery("SELECT").WithArgs(uuid.String()).WillReturnRows(rows)

	r := adapter.NewMySQLUserRepository(db)
	u, err := r.Get(uuid)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found (10002)", err.Error())
	assert.Empty(t, u.GetUsername())
}

func TestAdapterAdd(t *testing.T) {
	db, mock := InitializeDatabase()
	u, _ := entity.NewUser("test2", "password2", 2)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").
		WithArgs(u.GetID().String(), u.GetUsername(), u.GetPassword(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	r := adapter.NewMySQLUserRepository(db)
	err := r.Add(u)

	assert.Nil(t, err)
}

func TestAdapterAddWithDatabaseError(t *testing.T) {
	db, mock := InitializeDatabase()
	u, _ := entity.NewUser("test2", "password2", 2)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT").WillReturnError(gorm.ErrInvalidDB)
	mock.ExpectRollback()

	r := adapter.NewMySQLUserRepository(db)
	err := r.Add(u)

	assert.NotNil(t, err)
	assert.Equal(t, "failed to add (10004)", err.Error())
}

func TestAdapterRename(t *testing.T) {
	db, mock := InitializeDatabase()
	u, _ := entity.NewUser("test3", "password3", 3)
	uid := u.GetID().String()

	rows := sqlmock.NewRows(UserSchema)
	rows = rows.AddRow(uid, u.GetUsername(), u.GetPassword(), u.GetUserID())

	mock.ExpectQuery("SELECT").WithArgs(uid).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").
		WithArgs(u.GetUsername(), sqlmock.AnyArg(), uid).
		WillReturnResult(sqlmock.NewResult(2, 1))
	mock.ExpectCommit()

	r := adapter.NewMySQLUserRepository(db)
	err := r.Rename(u)

	assert.Nil(t, err)
}

func TestAdapterRenameWithErrUserNotFound(t *testing.T) {
	db, mock := InitializeDatabase()
	u, _ := entity.NewUser("test3", "password3", 3)
	uid := u.GetID().String()

	rows := sqlmock.NewRows(UserSchema)
	mock.ExpectQuery("SELECT").WithArgs(uid).WillReturnRows(rows)

	r := adapter.NewMySQLUserRepository(db)
	err := r.Rename(u)

	assert.NotNil(t, err)
	assert.Equal(t, "user not found (10005)", err.Error())
}

func TestAdapterRenameWithDatabaseErrror(t *testing.T) {
	db, mock := InitializeDatabase()
	u, _ := entity.NewUser("test3", "password3", 3)
	uid := u.GetID().String()

	rows := sqlmock.NewRows(UserSchema)
	rows = rows.AddRow(uid, u.GetUsername(), u.GetPassword(), u.GetUserID())
	mock.ExpectQuery("SELECT").WithArgs(uid).WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnError(gorm.ErrInvalidValue)
	mock.ExpectRollback()

	r := adapter.NewMySQLUserRepository(db)
	err := r.Rename(u)

	assert.NotNil(t, err)
	assert.Equal(t, "failed to rename (10004)", err.Error())
}
