package database_postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/teyz/go-svc-template/pkg/constants"
	"github.com/teyz/go-svc-template/pkg/errors"

	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
)

func (d *dbClient) CreateExample(ctx context.Context, description string) (*entities_example_v1.Example, error) {
	exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
	now := time.Now()

	_, err := d.connection.DB.ExecContext(ctx,
		`INSERT INTO 
			examples (
				id,
				description,
				created_at, 
				updated_at
			) 
			VALUES ($1, $2, $3, $4)
		`,
		exampleID, description, now, now)
	if err != nil {
		log.Error().Err(err).
			Msgf("database.postgres.dbClient.CreateExample: failed to create example: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.CreateExample: failed to create description: %v", err.Error()))
	}

	return &entities_example_v1.Example{
		ID:          exampleID,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (d *dbClient) GetExampleByID(ctx context.Context, id string) (*entities_example_v1.Example, error) {
	example := &entities_example_v1.Example{}

	err := d.connection.DB.QueryRowContext(ctx,
		`SELECT
			id,
			description,
			created_at,
			updated_at
		FROM
			examples
		WHERE
			id = $1
		`,
		id).Scan(
		&example.ID,
		&example.Description,
		&example.CreatedAt,
		&example.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error().Err(err).
				Str("id", id).
				Msgf("database.postgres.dbClient.GetExampleByID: example with id: %s not found", id)
			return nil, errors.NewNotFoundError(fmt.Sprintf("database.postgres.dbClient.GetExampleByID: example with id: %s not found", id))
		}

		log.Error().Err(err).
			Str("id", id).
			Msgf("database.postgres.dbClient.GetExampleByID: failed to get example by id: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetExampleByID: failed to get example by id: %v", err.Error()))
	}

	return example, nil
}

func (d *dbClient) GetExamples(ctx context.Context) ([]*entities_example_v1.Example, error) {
	rows, err := d.connection.DB.QueryContext(ctx, `
		SELECT
			id,
			description,
			created_at,
			updated_at
		FROM
			examples
	`)
	if err != nil {
		log.Error().Err(err).
			Msgf("database.postgres.dbClient.GetExamples: failed to get examples: %v", err.Error())
		return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetExamples: failed to get examples: %v", err.Error()))
	}
	defer rows.Close()

	examples := make([]*entities_example_v1.Example, 0)

	for rows.Next() {
		example := &entities_example_v1.Example{}

		err := rows.Scan(
			&example.ID,
			&example.Description,
			&example.CreatedAt,
			&example.UpdatedAt,
		)
		if err != nil {
			log.Error().Err(err).
				Msgf("database.postgres.dbClient.GetExamples: failed to scan example: %v", err.Error())
			return nil, errors.NewInternalServerError(fmt.Sprintf("database.postgres.dbClient.GetExamples: failed to scan example: %v", err.Error()))
		}

		examples = append(examples, example)
	}

	return examples, nil
}
