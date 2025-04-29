package mocks

import (
	"Spotify/internal/model"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockSongRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSongRepositoryMockRecorder
}

type MockSongRepositoryMockRecorder struct {
	mock *MockSongRepository
}

func NewMockSongRepository(ctrl *gomock.Controller) *MockSongRepository {
	mock := &MockSongRepository{ctrl: ctrl}
	mock.recorder = &MockSongRepositoryMockRecorder{mock}
	return mock
}

func (m *MockSongRepository) EXPECT() *MockSongRepositoryMockRecorder {
	return m.recorder
}

func (m *MockSongRepository) Create(song *model.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", song)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSongRepositoryMockRecorder) Create(song interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSongRepository)(nil).Create), song)
}

func (m *MockSongRepository) Update(song *model.Song) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", song)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSongRepositoryMockRecorder) Update(song interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockSongRepository)(nil).Update), song)
}

func (m *MockSongRepository) Delete(songID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", songID)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockSongRepositoryMockRecorder) Delete(songID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSongRepository)(nil).Delete), songID)
}

// Add the missing methods

func (m *MockSongRepository) GetByID(id uint) (*model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSongRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockSongRepository)(nil).GetByID), id)
}

func (m *MockSongRepository) GetByAlbumID(albumID uint) ([]model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAlbumID", albumID)
	ret0, _ := ret[0].([]model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSongRepositoryMockRecorder) GetByAlbumID(albumID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAlbumID", reflect.TypeOf((*MockSongRepository)(nil).GetByAlbumID), albumID)
}

func (m *MockSongRepository) GetAll() ([]model.Song, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Song)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSongRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockSongRepository)(nil).GetAll))
}
