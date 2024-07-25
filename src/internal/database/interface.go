package database

import (
	"context"

	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
)

//go:generate mockgen -source interface.go -destination mocks/mock_database.go -package database_mocks
type Database interface {
	CreateExample(ctx context.Context, description string) (*entities_example_v1.Example, error)
	GetExampleByID(ctx context.Context, id string) (*entities_example_v1.Example, error)
	GetExamples(ctx context.Context) ([]*entities_example_v1.Example, error)
}
