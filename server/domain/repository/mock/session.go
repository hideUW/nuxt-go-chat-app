// Code generated by MockGen. DO NOT EDIT.
// Source: domain/repository/session.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/hideUW/nuxt-go-chat-app/server/domain/model"
	repository "github.com/hideUW/nuxt-go-chat-app/server/domain/repository"
)

// MockSessionRepository is a mock of SessionRepository interface
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// GetSessionByID mocks base method
func (m_2 *MockSessionRepository) GetSessionByID(m repository.SQLManager, id string) (*model.Session, error) {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "GetSessionByID", m, id)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionByID indicates an expected call of GetSessionByID
func (mr *MockSessionRepositoryMockRecorder) GetSessionByID(m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionByID", reflect.TypeOf((*MockSessionRepository)(nil).GetSessionByID), m, id)
}

// InsertSession mocks base method
func (m_2 *MockSessionRepository) InsertSession(m repository.SQLManager, user *model.Session) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "InsertSession", m, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertSession indicates an expected call of InsertSession
func (mr *MockSessionRepositoryMockRecorder) InsertSession(m, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertSession", reflect.TypeOf((*MockSessionRepository)(nil).InsertSession), m, user)
}

// DeleteSession mocks base method
func (m_2 *MockSessionRepository) DeleteSession(m repository.SQLManager, id string) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "DeleteSession", m, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession
func (mr *MockSessionRepositoryMockRecorder) DeleteSession(m, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessionRepository)(nil).DeleteSession), m, id)
}
