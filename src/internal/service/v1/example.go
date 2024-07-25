package service_v1

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
)

func (s *service) CreateExample(ctx context.Context, description string) (*entities_example_v1.Example, error) {
	example, err := s.store.CreateExample(ctx, description)
	if err != nil {
		return nil, err
	}

	s.cache.Del(ctx, generateExamplesCacheKey())

	return example, nil
}

func (s *service) GetExamples(ctx context.Context) ([]*entities_example_v1.Example, error) {
	key := generateExamplesCacheKey()

	cacheExamples, err := s.cache.Get(ctx, key)
	if err == nil {
		var examples []*entities_example_v1.Example
		err = json.Unmarshal([]byte(cacheExamples), &examples)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetExamples: unable to unmarshal examples")
		} else {
			return examples, nil
		}
	}

	examples, err := s.store.GetExamples(ctx)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(examples)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetExamples: unable to marshal examples")
	} else {
		s.cache.SetEx(ctx, key, bytes, exampleCacheDuration)
	}

	return examples, nil
}

func (s *service) GetExampleByID(ctx context.Context, id string) (*entities_example_v1.Example, error) {
	key := generateExampleCacheKeyWithID(id)

	cacheExample, err := s.cache.Get(ctx, key)
	if err == nil {
		var example *entities_example_v1.Example
		err = json.Unmarshal([]byte(cacheExample), &example)
		if err != nil {
			log.Error().Err(err).
				Msg("service.v1.service.GetExampleByID: unable to unmarshal example")
		} else {
			return example, nil
		}
	}

	example, err := s.store.GetExampleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(example)
	if err != nil {
		log.Error().Err(err).
			Msg("service.v1.service.GetExampleByID: unable to marshal example")
	} else {
		s.cache.SetEx(ctx, key, bytes, exampleCacheDuration)
	}

	return example, nil
}
