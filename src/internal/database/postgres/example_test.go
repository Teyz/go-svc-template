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
	"github.com/teyz/go-svc-template/pkg/errors"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_CreateExample(t *testing.T) {
	t.Run("ok - create example", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectExec("INSERT INTO examples").WithArgs(sqlmock.AnyArg(), "hello world !", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

		example, err := sqlxDB.CreateExample(context.Background(), "hello world !")
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.True(t, constants.Example.IsValid(example.ID))
		assert.Equal(t, "hello world !", example.Description)
		assert.False(t, example.CreatedAt.IsZero())
		assert.False(t, example.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - create example", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectExec("INSERT INTO examples").WithArgs(sqlmock.AnyArg(), "hello world !", sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(errors.NewInternalServerError("error"))

		channel, err := sqlxDB.CreateExample(context.Background(), "hello world !")
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetExampleByID(t *testing.T) {
	t.Run("ok - get example by id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)

		rows := sqlmock.NewRows([]string{"id", "description", "created_at", "updated_at"}).
			AddRow(exampleID, "hello world !", time.Now(), time.Now())

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples WHERE id = $1").WithArgs(exampleID).WillReturnError(nil).WillReturnRows(rows)

		example, err := sqlxDB.GetExampleByID(context.Background(), exampleID)
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.True(t, constants.Example.IsValid(example.ID))
		assert.Equal(t, "hello world !", example.Description)
		assert.False(t, example.CreatedAt.IsZero())
		assert.False(t, example.UpdatedAt.IsZero())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get example by id", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples WHERE id = $1").WithArgs(exampleID).WillReturnError(errors.NewInternalServerError("error"))

		example, err := sqlxDB.GetExampleByID(context.Background(), exampleID)
		assert.Nil(t, example)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get example by id - no rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples WHERE id = $1").WithArgs(exampleID).WillReturnError(sql.ErrNoRows)

		channel, err := sqlxDB.GetExampleByID(context.Background(), exampleID)
		assert.Nil(t, channel)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func Test_GetExamples(t *testing.T) {
	t.Run("ok - get examples", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)

		rows := sqlmock.NewRows([]string{"id", "description", "created_at", "updated_at"}).
			AddRow(exampleID, "hello world !", time.Now(), time.Now())

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples").WillReturnError(nil).WillReturnRows(rows)

		examples, err := sqlxDB.FetchExamples(context.Background())
		assert.NotNil(t, examples)
		assert.NoError(t, err)

		for _, example := range examples {
			assert.True(t, constants.Example.IsValid(example.ID))
			assert.Equal(t, "hello world !", example.Description)
			assert.False(t, example.CreatedAt.IsZero())
			assert.False(t, example.UpdatedAt.IsZero())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get examples", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples").WillReturnError(errors.NewInternalServerError("error"))

		examples, err := sqlxDB.FetchExamples(context.Background())
		assert.Nil(t, examples)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
	t.Run("nok - get examples - no rows", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		assert.NoError(t, err)
		defer db.Close()

		sqlxDB := &dbClient{
			connection: sqlx.NewDb(db, "sqlmock"),
		}

		mock.ExpectQuery("SELECT id, description, created_at, updated_at FROM examples").WillReturnError(sql.ErrNoRows)

		examples, err := sqlxDB.FetchExamples(context.Background())
		assert.Nil(t, examples)
		assert.Error(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
