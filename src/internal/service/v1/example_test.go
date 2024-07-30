package service_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	database_mocks "github.com/teyz/go-svc-template/internal/database/mocks"
	entities_example_v1 "github.com/teyz/go-svc-template/internal/entities/example/v1"
	"github.com/teyz/go-svc-template/pkg/constants"
	"github.com/teyz/go-svc-template/pkg/errors"
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

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().Get(gomock.Any(), "teyz").Return(string(channelCachedBytes), nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelBySlug(context.Background(), "teyz")
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
	t.Run("ok - get example by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		created := time.Now()

		mock_database.EXPECT().GetChannelBySlug(gomock.Any(), "teyz").Return(&entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "teyz").Return("", pkgerrors.NewNotFoundError("error"))

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "chnl-store-svc:channel:slug:teyz", channelCachedBytes, time.Hour*24).Return(nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelBySlug(context.Background(), "teyz")
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
	t.Run("nok - get example by id from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_database.EXPECT().GetChannelBySlug(gomock.Any(), "teyz").Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), "teyz").Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelBySlug(context.Background(), "teyz")
		assert.Nil(t, channel)
		assert.Error(t, err)
	})
	t.Run("ok - get example by id when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), "teyz").Return(fakeData, nil)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_database.EXPECT().GetChannelBySlug(gomock.Any(), "teyz").Return(&entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), "chnl-store-svc:channel:slug:teyz", channelCachedBytes, time.Hour*24).Return(nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelBySlug(context.Background(), "teyz")
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
}

func Test_GetExamples(t *testing.T) {
	t.Run("ok - get examples from cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		exampleID := constants.GenerateDataPrefixWithULID(constants.Example)
		created := time.Now()

		mock_cache := cache_mocks.NewMockCache(ctrl)

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().Get(gomock.Any(), channelID).Return(string(channelCachedBytes), nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelByID(context.Background(), channelID)
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
	t.Run("ok - get examples from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		created := time.Now()

		mock_database.EXPECT().GetChannelByID(gomock.Any(), channelID).Return(&entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), channelID).Return("", pkgerrors.NewNotFoundError("error"))

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("chnl-store-svc:channel:id:%+v", channelID), channelCachedBytes, time.Hour*24).Return(nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelByID(context.Background(), channelID)
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
	t.Run("nok - get examples from database", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)

		mock_database.EXPECT().GetChannelByID(gomock.Any(), channelID).Return(nil, pkgerrors.NewNotFoundError("error"))

		mock_cache := cache_mocks.NewMockCache(ctrl)

		mock_cache.EXPECT().Get(gomock.Any(), channelID).Return("", pkgerrors.NewNotFoundError("error"))

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelByID(context.Background(), channelID)
		assert.Nil(t, channel)
		assert.Error(t, err)
	})
	t.Run("ok - get examples when get cache", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mock_database := database_mocks.NewMockDatabase(ctrl)
		mock_cache := cache_mocks.NewMockCache(ctrl)

		channelID := constants.GenerateDataPrefixWithULID(constants.Channel)
		created := time.Now()
		fakeData := `abczd{>`

		mock_cache.EXPECT().Get(gomock.Any(), channelID).Return(fakeData, nil)

		mock_database.EXPECT().GetChannelByID(gomock.Any(), channelID).Return(&entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}, nil)

		channelCached := &entities_channel_v1.Channel{
			ID:        channelID,
			UserID:    "user_id",
			Slug:      "teyz",
			CreatedAt: created,
			UpdatedAt: created,
		}

		channelCachedBytes, _ := json.Marshal(channelCached)

		mock_cache.EXPECT().SetEx(gomock.Any(), fmt.Sprintf("chnl-store-svc:channel:id:%+v", channelID), channelCachedBytes, time.Hour*24).Return(nil)

		s, err := NewChannelStoreService(context.Background(), mock_database, mock_cache)
		assert.NotNil(t, s)
		assert.NoError(t, err)

		channel, err := s.GetChannelByID(context.Background(), channelID)
		assert.NotNil(t, channel)
		assert.NoError(t, err)

		assert.Equal(t, channelID, channel.ID)
		assert.Equal(t, "user_id", channel.UserID)
		assert.Equal(t, "teyz", channel.Slug)
		assert.True(t, channel.CreatedAt.Equal(created))
		assert.True(t, channel.UpdatedAt.Equal(created))
	})
}
