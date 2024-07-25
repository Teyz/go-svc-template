package database_postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/teyz/go-svc-template/internal/database"
)

type dbClient struct {
	connection *sqlx.DB
}

func NewClient(ctx context.Context, db *sqlx.DB) database.Database {
	return &dbClient{
		connection: db,
	}
}
