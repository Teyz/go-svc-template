package database_postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/teyz/go-svc-template/pkg/constants"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_CreateChannel(t *testing.T) {
	t.Run("ok - create channel", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.Example)

		mock.ExpectExec("INSERT INTO channels").WithArgs(sqlmock.AnyArg(), userID, "teyz", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		channel, err := sqlxDB.CreateChannel(context.Background(), userID, "teyz")
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.True(t, constants.Channel.IsValid(channel.ID))
		assert.True(t, constants.User.IsValid(channel.UserID))
		assert.Equal(t, "teyz", channel.Slug)
		assert.False(t, channel.CreatedAt.IsZero())
		assert.False(t, channel.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - create channel", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.User)

		mock.ExpectExec("INSERT INTO channels").WithArgs(sqlmock.AnyArg(), userID, "teyz", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(pkgerrors.NewInternalServerError("error"))

		channel, err := sqlxDB.CreateChannel(context.Background(), userID, "teyz")
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetChannelByID(t *testing.T) {
	t.Run("ok - get channel by id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		userID := constants.GenerateDataPrefixWithULID(constants.User)

		rows := sqlmock.NewRows([]string{"id", "user_id", "slug", "created_at", "updated_at"}).
			AddRow(channelID, userID, "teyz", time.Now(), time.Now())

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE id = $1").WithArgs(channelID).WillReturnError(nil).WillReturnRows(rows)

		channel, err := sqlxDB.GetChannelByID(context.Background(), channelID)
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.True(t, constants.Channel.IsValid(channel.ID))
		assert.True(t, constants.User.IsValid(channel.UserID))
		assert.Equal(t, "teyz", channel.Slug)
		assert.False(t, channel.CreatedAt.IsZero())
		assert.False(t, channel.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE id = $1").WithArgs(channelID).WillReturnError(pkgerrors.NewInternalServerError("error"))

		channel, err := sqlxDB.GetChannelByID(context.Background(), channelID)
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by id - no rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE id = $1").WithArgs(channelID).WillReturnError(sql.ErrNoRows)

		channel, err := sqlxDB.GetChannelByID(context.Background(), channelID)
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetChannelByUserID(t *testing.T) {
	t.Run("ok - get channel by user_id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		userID := constants.GenerateDataPrefixWithULID(constants.User)

		rows := sqlmock.NewRows([]string{"id", "user_id", "slug", "created_at", "updated_at"}).
			AddRow(channelID, userID, "teyz", time.Now(), time.Now())

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE user_id = $1").WithArgs(userID).WillReturnError(nil).WillReturnRows(rows)

		channel, err := sqlxDB.GetChannelByUserID(context.Background(), userID)
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.True(t, constants.Channel.IsValid(channel.ID))
		assert.True(t, constants.User.IsValid(channel.UserID))
		assert.Equal(t, "teyz", channel.Slug)
		assert.False(t, channel.CreatedAt.IsZero())
		assert.False(t, channel.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by user_id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.User)

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE user_id = $1").WithArgs(userID).WillReturnError(pkgerrors.NewInternalServerError("error"))

		channel, err := sqlxDB.GetChannelByUserID(context.Background(), userID)
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by user_id - no rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		userID := constants.GenerateDataPrefixWithULID(constants.User)

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE user_id = $1").WithArgs(userID).WillReturnError(sql.ErrNoRows)

		channel, err := sqlxDB.GetChannelByUserID(context.Background(), userID)
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetChannelBySlug(t *testing.T) {
	t.Run("ok - get channel by slug", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		userID := constants.GenerateDataPrefixWithULID(constants.User)

		rows := sqlmock.NewRows([]string{"id", "user_id", "slug", "created_at", "updated_at"}).
			AddRow(channelID, userID, "teyz", time.Now(), time.Now())

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE slug = $1").WithArgs("teyz").WillReturnError(nil).WillReturnRows(rows)

		channel, err := sqlxDB.GetChannelBySlug(context.Background(), "teyz")
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.True(t, constants.Channel.IsValid(channel.ID))
		assert.True(t, constants.User.IsValid(channel.UserID))
		assert.Equal(t, "teyz", channel.Slug)
		assert.False(t, channel.CreatedAt.IsZero())
		assert.False(t, channel.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by slug", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE slug = $1").WithArgs("teyz").WillReturnError(pkgerrors.NewInternalServerError("error"))

		channel, err := sqlxDB.GetChannelBySlug(context.Background(), "teyz")
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get channel by slug - no rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectQuery("SELECT id, user_id, slug, created_at, updated_at FROM channels WHERE slug = $1").WithArgs("teyz").WillReturnError(sql.ErrNoRows)

		channel, err := sqlxDB.GetChannelBySlug(context.Background(), "teyz")
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
