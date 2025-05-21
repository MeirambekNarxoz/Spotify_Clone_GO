package service

import (
	"Spotify/internal/model"
	"Spotify/internal/repository/mocks"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSongService_CreateSong_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	testSong := &model.Song{
		Title:  "Test Song",
		Artist: "Test Artist",
	}

	mockRepo.EXPECT().Create(testSong).Return(nil)

	err := service.CreateSong(testSong)

	assert.NoError(t, err)
}

func TestSongService_CreateSong_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	// Тест на пустой заголовок
	t.Run("Empty title", func(t *testing.T) {
		err := service.CreateSong(&model.Song{
			Artist: "Test Artist",
		})
		assert.EqualError(t, err, "song title cannot be empty")
	})

	// Тест на пустого исполнителя
	t.Run("Empty artist", func(t *testing.T) {
		err := service.CreateSong(&model.Song{
			Title: "Test Song",
		})
		assert.EqualError(t, err, "song artist cannot be empty")
	})

	// Убедимся, что репозиторий не вызывался
	mockRepo.EXPECT().Create(gomock.Any()).Times(0)
}

func TestSongService_CreateSong_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	testSong := &model.Song{
		Title:  "Test Song",
		Artist: "Test Artist",
	}
	expectedError := errors.New("repository error")

	mockRepo.EXPECT().Create(testSong).Return(expectedError)

	err := service.CreateSong(testSong)

	assert.EqualError(t, err, expectedError.Error())
}

func TestSongService_UpdateSong_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	testSong := &model.Song{
		Title:  "Updated Title",
		Artist: "Updated Artist",
	}

	mockRepo.EXPECT().Update(testSong).Return(nil)

	err := service.UpdateSong(testSong)

	assert.NoError(t, err)
}

func TestSongService_UpdateSong_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	// Тест на пустой заголовок
	t.Run("Empty title", func(t *testing.T) {
		err := service.UpdateSong(&model.Song{
			Artist: "Test Artist",
		})
		assert.EqualError(t, err, "song title cannot be empty")
	})

	// Тест на пустого исполнителя
	t.Run("Empty artist", func(t *testing.T) {
		err := service.UpdateSong(&model.Song{

			Title: "Test Song",
		})
		assert.EqualError(t, err, "song artist cannot be empty")
	})

	// Убедимся, что репозиторий не вызывался
	mockRepo.EXPECT().Update(gomock.Any()).Times(0)
}

func TestSongService_DeleteSong_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	songID := uint(1)

	mockRepo.EXPECT().Delete(songID).Return(nil)

	err := service.DeleteSong(songID)

	assert.NoError(t, err)
}

func TestSongService_DeleteSong_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockSongRepository(ctrl)
	service := NewSongService(mockRepo)

	songID := uint(1)
	expectedError := errors.New("repository error")

	mockRepo.EXPECT().Delete(songID).Return(expectedError)

	err := service.DeleteSong(songID)

	assert.EqualError(t, err, expectedError.Error())
}
