package mock_usecase

import (
	reflect "reflect"
	domain "todolist-backend/modules/v1/activities/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockActivityAdapter is a mock of ActivityAdapter interface.
type MockActivityAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockActivityAdapterMockRecorder
}

// MockActivityAdapterMockRecorder is the mock recorder for MockActivityAdapter.
type MockActivityAdapterMockRecorder struct {
	mock *MockActivityAdapter
}

// NewMockActivityAdapter creates a new mock instance.
func NewMockActivityAdapter(ctrl *gomock.Controller) *MockActivityAdapter {
	mock := &MockActivityAdapter{ctrl: ctrl}
	mock.recorder = &MockActivityAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActivityAdapter) EXPECT() *MockActivityAdapterMockRecorder {
	return m.recorder
}

// CreateActivity mocks base method.
func (m *MockActivityAdapter) CreateActivity(activity domain.InsertActivity) (domain.Activities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActivity", activity)
	ret0, _ := ret[0].(domain.Activities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActivity indicates an expected call of CreateActivity.
func (mr *MockActivityAdapterMockRecorder) CreateActivity(activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActivity", reflect.TypeOf((*MockActivityAdapter)(nil).CreateActivity), activity)
}

// DeleteActivity mocks base method.
func (m *MockActivityAdapter) DeleteActivity(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActivity", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActivity indicates an expected call of DeleteActivity.
func (mr *MockActivityAdapterMockRecorder) DeleteActivity(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActivity", reflect.TypeOf((*MockActivityAdapter)(nil).DeleteActivity), id)
}

// GetActivityByID mocks base method.
func (m *MockActivityAdapter) GetActivityByID(id string) (domain.Activities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivityByID", id)
	ret0, _ := ret[0].(domain.Activities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivityByID indicates an expected call of GetActivityByID.
func (mr *MockActivityAdapterMockRecorder) GetActivityByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivityByID", reflect.TypeOf((*MockActivityAdapter)(nil).GetActivityByID), id)
}

// GetAllActivity mocks base method.
func (m *MockActivityAdapter) GetAllActivity() ([]domain.Activities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllActivity")
	ret0, _ := ret[0].([]domain.Activities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllActivity indicates an expected call of GetAllActivity.
func (mr *MockActivityAdapterMockRecorder) GetAllActivity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllActivity", reflect.TypeOf((*MockActivityAdapter)(nil).GetAllActivity))
}

// UpdateActivity mocks base method.
func (m *MockActivityAdapter) UpdateActivity(id string, activity domain.UpdateActivity) (domain.Activities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActivity", id, activity)
	ret0, _ := ret[0].(domain.Activities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActivity indicates an expected call of UpdateActivity.
func (mr *MockActivityAdapterMockRecorder) UpdateActivity(id, activity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActivity", reflect.TypeOf((*MockActivityAdapter)(nil).UpdateActivity), id, activity)
}
