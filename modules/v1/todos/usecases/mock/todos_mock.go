package mock_usecase

import (
	reflect "reflect"
	domain "todolist-backend/modules/v1/todos/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockTodoAdapter is a mock of TodoAdapter interface.
type MockTodoAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockTodoAdapterMockRecorder
}

// MockTodoAdapterMockRecorder is the mock recorder for MockTodoAdapter.
type MockTodoAdapterMockRecorder struct {
	mock *MockTodoAdapter
}

// NewMockTodoAdapter creates a new mock instance.
func NewMockTodoAdapter(ctrl *gomock.Controller) *MockTodoAdapter {
	mock := &MockTodoAdapter{ctrl: ctrl}
	mock.recorder = &MockTodoAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoAdapter) EXPECT() *MockTodoAdapterMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockTodoAdapter) CreateTodo(arg0 domain.Todos) (domain.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0)
	ret0, _ := ret[0].(domain.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockTodoAdapterMockRecorder) CreateTodo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockTodoAdapter)(nil).CreateTodo), arg0)
}

// DeleteTodo mocks base method.
func (m *MockTodoAdapter) DeleteTodo(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockTodoAdapterMockRecorder) DeleteTodo(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockTodoAdapter)(nil).DeleteTodo), id)
}

// GetAllTodos mocks base method.
func (m *MockTodoAdapter) GetAllTodos() ([]domain.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTodos")
	ret0, _ := ret[0].([]domain.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTodos indicates an expected call of GetAllTodos.
func (mr *MockTodoAdapterMockRecorder) GetAllTodos() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTodos", reflect.TypeOf((*MockTodoAdapter)(nil).GetAllTodos))
}

// GetTodoById mocks base method.
func (m *MockTodoAdapter) GetTodoById(id string) (domain.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoById", id)
	ret0, _ := ret[0].(domain.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoById indicates an expected call of GetTodoById.
func (mr *MockTodoAdapterMockRecorder) GetTodoById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoById", reflect.TypeOf((*MockTodoAdapter)(nil).GetTodoById), id)
}

// GetTodosByGroupId mocks base method.
func (m *MockTodoAdapter) GetTodosByGroupId(group_id string) ([]domain.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodosByGroupId", group_id)
	ret0, _ := ret[0].([]domain.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodosByGroupId indicates an expected call of GetTodosByGroupId.
func (mr *MockTodoAdapterMockRecorder) GetTodosByGroupId(group_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodosByGroupId", reflect.TypeOf((*MockTodoAdapter)(nil).GetTodosByGroupId), group_id)
}

// UpdateTodo mocks base method.
func (m *MockTodoAdapter) UpdateTodo(id string, todos domain.UpdateTodos) (domain.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", id, todos)
	ret0, _ := ret[0].(domain.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockTodoAdapterMockRecorder) UpdateTodo(id, todos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoAdapter)(nil).UpdateTodo), id, todos)
}
