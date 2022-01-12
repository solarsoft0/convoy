// Code generated by MockGen. DO NOT EDIT.
// Source: event.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	convoy "github.com/frain-dev/convoy"
	models "github.com/frain-dev/convoy/server/models"
	gomock "github.com/golang/mock/gomock"
)

// MockEventRepository is a mock of EventRepository interface.
type MockEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockEventRepositoryMockRecorder
}

// MockEventRepositoryMockRecorder is the mock recorder for MockEventRepository.
type MockEventRepositoryMockRecorder struct {
	mock *MockEventRepository
}

// NewMockEventRepository creates a new mock instance.
func NewMockEventRepository(ctrl *gomock.Controller) *MockEventRepository {
	mock := &MockEventRepository{ctrl: ctrl}
	mock.recorder = &MockEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEventRepository) EXPECT() *MockEventRepositoryMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockEventRepository) CreateEvent(arg0 context.Context, arg1 *convoy.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockEventRepositoryMockRecorder) CreateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEventRepository)(nil).CreateEvent), arg0, arg1)
}

// FindEventByID mocks base method.
func (m *MockEventRepository) FindEventByID(ctx context.Context, id string) (*convoy.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEventByID", ctx, id)
	ret0, _ := ret[0].(*convoy.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindEventByID indicates an expected call of FindEventByID.
func (mr *MockEventRepositoryMockRecorder) FindEventByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEventByID", reflect.TypeOf((*MockEventRepository)(nil).FindEventByID), ctx, id)
}

// LoadAbandonedEventsForPostingRetry mocks base method.
func (m *MockEventRepository) LoadAbandonedEventsForPostingRetry(arg0 context.Context) ([]convoy.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadAbandonedEventsForPostingRetry", arg0)
	ret0, _ := ret[0].([]convoy.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadAbandonedEventsForPostingRetry indicates an expected call of LoadAbandonedEventsForPostingRetry.
func (mr *MockEventRepositoryMockRecorder) LoadAbandonedEventsForPostingRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadAbandonedEventsForPostingRetry", reflect.TypeOf((*MockEventRepository)(nil).LoadAbandonedEventsForPostingRetry), arg0)
}

// LoadEventIntervals mocks base method.
func (m *MockEventRepository) LoadEventIntervals(arg0 context.Context, arg1 string, arg2 models.SearchParams, arg3 convoy.Period, arg4 int) ([]models.EventInterval, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventIntervals", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]models.EventInterval)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventIntervals indicates an expected call of LoadEventIntervals.
func (mr *MockEventRepositoryMockRecorder) LoadEventIntervals(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventIntervals", reflect.TypeOf((*MockEventRepository)(nil).LoadEventIntervals), arg0, arg1, arg2, arg3, arg4)
}

// LoadEventsForPostingRetry mocks base method.
func (m *MockEventRepository) LoadEventsForPostingRetry(arg0 context.Context) ([]convoy.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsForPostingRetry", arg0)
	ret0, _ := ret[0].([]convoy.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventsForPostingRetry indicates an expected call of LoadEventsForPostingRetry.
func (mr *MockEventRepositoryMockRecorder) LoadEventsForPostingRetry(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsForPostingRetry", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsForPostingRetry), arg0)
}

// LoadEventsPaged mocks base method.
func (m *MockEventRepository) LoadEventsPaged(arg0 context.Context, arg1, arg2 string, arg3 models.SearchParams, arg4 models.Pageable) ([]convoy.Event, models.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsPaged", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]convoy.Event)
	ret1, _ := ret[1].(models.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadEventsPaged indicates an expected call of LoadEventsPaged.
func (mr *MockEventRepositoryMockRecorder) LoadEventsPaged(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsPaged", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsPaged), arg0, arg1, arg2, arg3, arg4)
}

// LoadEventsPagedByAppId mocks base method.
func (m *MockEventRepository) LoadEventsPagedByAppId(arg0 context.Context, arg1 string, arg2 models.SearchParams, arg3 models.Pageable) ([]convoy.Event, models.PaginationData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsPagedByAppId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]convoy.Event)
	ret1, _ := ret[1].(models.PaginationData)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadEventsPagedByAppId indicates an expected call of LoadEventsPagedByAppId.
func (mr *MockEventRepositoryMockRecorder) LoadEventsPagedByAppId(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsPagedByAppId", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsPagedByAppId), arg0, arg1, arg2, arg3)
}

// LoadEventsScheduledForPosting mocks base method.
func (m *MockEventRepository) LoadEventsScheduledForPosting(arg0 context.Context) ([]convoy.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadEventsScheduledForPosting", arg0)
	ret0, _ := ret[0].([]convoy.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadEventsScheduledForPosting indicates an expected call of LoadEventsScheduledForPosting.
func (mr *MockEventRepositoryMockRecorder) LoadEventsScheduledForPosting(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadEventsScheduledForPosting", reflect.TypeOf((*MockEventRepository)(nil).LoadEventsScheduledForPosting), arg0)
}
