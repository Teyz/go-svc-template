package service_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	database_mocks "github.com/teyz/go-svc-template/internal/database/mocks"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
	cache_mocks "github.com/teyz/go-svc-template/pkg/cache/mocks"
	"github.com/teyz/go-svc-template/pkg/constants"
	"github.com/teyz/go-svc-template/pkg/errors"
	"go.uber.org/mock/gomock"
)

func Test_CreateExample(t *testing.T) {
	t.Run("ok - create example", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_database.EXPECT().CreateExample(gomock.Any(), "hello world !").Return(&entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}, nil)

		mock_cache.EXPECT().Del(gomock.Any(), generateExamplesCacheKey())

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.CreateExample(context.Background(), "hello world !")
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.Equal(t, exampleID, example.ID)
		assert.Equal(t, "hello world !", example.Description)
		assert.True(t, example.CreatedAt.Equal(created))
		assert.True(t, example.UpdatedAt.Equal(created))
	})
	t.Run("nok - create example", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().CreateExample(gomock.Any(), "hello world !").Return(nil, errors.NewInternalServerError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.CreateExample(context.Background(), "hello world !")
		assert.Nil(t, example)
		assert.Error(t, err)
	})
}

func Test_GetExampleByID(t *testing.T) {
	t.Run("ok - get example by id from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		exampleCached := &entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}

		exampleCachedBytes, _ := json.Marshal(exampleCached)

		mock_cache.EXPECT().Get(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", exampleID)).Return(string(exampleCachedBytes), nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.GetExampleByID(context.Background(), exampleID)
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.Equal(t, exampleID, example.ID)
		assert.Equal(t, "hello world !", example.Description)
		assert.True(t, example.CreatedAt.Equal(created))
		assert.True(t, example.UpdatedAt.Equal(created))
	})
	t.Run("ok - get example by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_database.EXPECT().GetExampleByID(gomock.Any(), exampleID).Return(&entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}, nil)

		mock_cache.EXPECT().Get(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", exampleID)).Return("", errors.NewNotFoundError("error"))

		exampleCached := &entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}

		exampleCachedBytes, _ := json.Marshal(exampleCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", exampleID), exampleCachedBytes, time.Hour*24).Return(nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.GetExampleByID(context.Background(), exampleID)
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.Equal(t, exampleID, example.ID)
		assert.Equal(t, "hello world !", example.Description)
		assert.True(t, example.CreatedAt.Equal(created))
		assert.True(t, example.UpdatedAt.Equal(created))
	})
	t.Run("nok - get example by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_database.EXPECT().GetExampleByID(gomock.Any(), "id").Return(nil, errors.NewNotFoundError("error"))

		mock_cache.EXPECT().Get(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", "id")).Return("", errors.NewNotFoundError("error"))

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.GetExampleByID(context.Background(), "id")
		assert.Nil(t, example)
		assert.Error(t, err)
	})
	t.Run("ok - get example by id when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_cache.EXPECT().Get(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", exampleID)).Return(fakeData, nil)

		mock_database.EXPECT().GetExampleByID(gomock.Any(), exampleID).Return(&entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}, nil)

		exampleCached := &entities_example_v1.Example{
			ID:          exampleID,
			Description: "hello world !",
			CreatedAt:   created,
			UpdatedAt:   created,
		}

		exampleCachedBytes, _ := json.Marshal(exampleCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("go-svc-template:example:id:%v", exampleID), exampleCachedBytes, time.Hour*24).Return(nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.GetExampleByID(context.Background(), exampleID)
		assert.NotNil(t, example)
		assert.NoError(t, err)

		assert.Equal(t, exampleID, example.ID)
		assert.Equal(t, "hello world !", example.Description)
		assert.True(t, example.CreatedAt.Equal(created))
		assert.True(t, example.UpdatedAt.Equal(created))
	})
}

func Test_GetExamples(t *testing.T) {
	t.Run("ok - get examples from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		examplesCached := []*entities_example_v1.Example{
			{
				ID:          exampleID,
				Description: "hello world !",
				CreatedAt:   created,
				UpdatedAt:   created,
			},
		}

		exampleCachedBytes, _ := json.Marshal(examplesCached)

		mock_cache.EXPECT().Get(gomock.Any(), "go-svc-template:examples").Return(string(exampleCachedBytes), nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		examples, err := s.FetchExamples(context.Background())
		assert.NotNil(t, examples)
		assert.NoError(t, err)

		for _, example := range examples {
			assert.Equal(t, exampleID, example.ID)
			assert.Equal(t, "hello world !", example.Description)
			assert.True(t, example.CreatedAt.Equal(created))
			assert.True(t, example.UpdatedAt.Equal(created))
		}
	})
	t.Run("ok - get examples from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		examplesResults := []*entities_example_v1.Example{
			{
				ID:          exampleID,
				Description: "hello world !",
				CreatedAt:   created,
				UpdatedAt:   created,
			},
		}

		mock_database.EXPECT().FetchExamples(gomock.Any()).Return(examplesResults, nil)

		mock_cache.EXPECT().Get(gomock.Any(), "go-svc-template:examples").Return("", errors.NewNotFoundError("error"))

		exampleCachedBytes, _ := json.Marshal(examplesResults)

		mock_cache.EXPECT().SetEx(gomock.Any(), "go-svc-template:examples", exampleCachedBytes, time.Hour*24).Return(nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		examples, err := s.FetchExamples(context.Background())
		assert.NotNil(t, examples)
		assert.NoError(t, err)

		for _, example := range examples {
			assert.Equal(t, exampleID, example.ID)
			assert.Equal(t, "hello world !", example.Description)
			assert.True(t, example.CreatedAt.Equal(created))
			assert.True(t, example.UpdatedAt.Equal(created))
		}
	})
	t.Run("nok - get examples from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_database.EXPECT().FetchExamples(gomock.Any()).Return(nil, errors.NewNotFoundError("error"))

		mock_cache.EXPECT().Get(gomock.Any(), "go-svc-template:examples").Return("", errors.NewNotFoundError("error"))

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		example, err := s.FetchExamples(context.Background())
		assert.Nil(t, example)
		assert.Error(t, err)
	})
	t.Run("ok - get examples when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_cache.EXPECT().Get(gomock.Any(), "go-svc-template:examples").Return(fakeData, nil)

		examplesResults := []*entities_example_v1.Example{
			{
				ID:          exampleID,
				Description: "hello world !",
				CreatedAt:   created,
				UpdatedAt:   created,
			},
		}

		mock_database.EXPECT().FetchExamples(gomock.Any()).Return(examplesResults, nil)

		exampleCachedBytes, _ := json.Marshal(examplesResults)

		mock_cache.EXPECT().SetEx(gomock.Any(), "go-svc-template:examples", exampleCachedBytes, time.Hour*24).Return(nil)

		s, err := NewExampleStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		examples, err := s.FetchExamples(context.Background())
		assert.NotNil(t, examples)
		assert.NoError(t, err)

		for _, example := range examples {
			assert.Equal(t, exampleID, example.ID)
			assert.Equal(t, "hello world !", example.Description)
			assert.True(t, example.CreatedAt.Equal(created))
			assert.True(t, example.UpdatedAt.Equal(created))
		}
	})
}
