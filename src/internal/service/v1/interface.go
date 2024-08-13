package service_v1

import (
	"context"

	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
)

type ExampleStoreService interface {
	CreateExample(ctx context.Context, description string) (*entities_example_v1.Example, error)
	FetchExamples(ctx context.Context) ([]*entities_example_v1.Example, error)
	GetExampleByID(ctx context.Context, id string) (*entities_example_v1.Example, error)
}
