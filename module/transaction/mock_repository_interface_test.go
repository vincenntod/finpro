// Code generated by mockery v2.30.1. DO NOT EDIT.

package transaction

import (
	mock "github.com/stretchr/testify/mock"
)

// MockRepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type MockRepositoryInterface struct {
	mock.Mock
}

type MockRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoryInterface) EXPECT() *MockRepositoryInterface_Expecter {
	return &MockRepositoryInterface_Expecter{mock: &_m.Mock}
}

// GetAllTransaction provides a mock function with given fields:
func (_m *MockRepositoryInterface) GetAllTransaction() ([]Transaction, error) {
	ret := _m.Called()

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Transaction, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Transaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetAllTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransaction'
type MockRepositoryInterface_GetAllTransaction_Call struct {
	*mock.Call
}

// GetAllTransaction is a helper method to define mock.On call
func (_e *MockRepositoryInterface_Expecter) GetAllTransaction() *MockRepositoryInterface_GetAllTransaction_Call {
	return &MockRepositoryInterface_GetAllTransaction_Call{Call: _e.mock.On("GetAllTransaction")}
}

func (_c *MockRepositoryInterface_GetAllTransaction_Call) Run(run func()) *MockRepositoryInterface_GetAllTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransaction_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetAllTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransaction_Call) RunAndReturn(run func() ([]Transaction, error)) *MockRepositoryInterface_GetAllTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByDate provides a mock function with given fields: start, end
func (_m *MockRepositoryInterface) GetAllTransactionByDate(start string, end string) ([]Transaction, error) {
	ret := _m.Called(start, end)

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]Transaction, error)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string) []Transaction); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetAllTransactionByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByDate'
type MockRepositoryInterface_GetAllTransactionByDate_Call struct {
	*mock.Call
}

// GetAllTransactionByDate is a helper method to define mock.On call
//   - start string
//   - end string
func (_e *MockRepositoryInterface_Expecter) GetAllTransactionByDate(start interface{}, end interface{}) *MockRepositoryInterface_GetAllTransactionByDate_Call {
	return &MockRepositoryInterface_GetAllTransactionByDate_Call{Call: _e.mock.On("GetAllTransactionByDate", start, end)}
}

func (_c *MockRepositoryInterface_GetAllTransactionByDate_Call) Run(run func(start string, end string)) *MockRepositoryInterface_GetAllTransactionByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByDate_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByDate_Call) RunAndReturn(run func(string, string) ([]Transaction, error)) *MockRepositoryInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatus provides a mock function with given fields: status
func (_m *MockRepositoryInterface) GetAllTransactionByStatus(status string) ([]Transaction, error) {
	ret := _m.Called(status)

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]Transaction, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) []Transaction); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetAllTransactionByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatus'
type MockRepositoryInterface_GetAllTransactionByStatus_Call struct {
	*mock.Call
}

// GetAllTransactionByStatus is a helper method to define mock.On call
//   - status string
func (_e *MockRepositoryInterface_Expecter) GetAllTransactionByStatus(status interface{}) *MockRepositoryInterface_GetAllTransactionByStatus_Call {
	return &MockRepositoryInterface_GetAllTransactionByStatus_Call{Call: _e.mock.On("GetAllTransactionByStatus", status)}
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatus_Call) Run(run func(status string)) *MockRepositoryInterface_GetAllTransactionByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatus_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatus_Call) RunAndReturn(run func(string) ([]Transaction, error)) *MockRepositoryInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatusDate provides a mock function with given fields: status, start, end
func (_m *MockRepositoryInterface) GetAllTransactionByStatusDate(status string, start string, end string) ([]Transaction, error) {
	ret := _m.Called(status, start, end)

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) ([]Transaction, error)); ok {
		return rf(status, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) []Transaction); ok {
		r0 = rf(status, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(status, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetAllTransactionByStatusDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatusDate'
type MockRepositoryInterface_GetAllTransactionByStatusDate_Call struct {
	*mock.Call
}

// GetAllTransactionByStatusDate is a helper method to define mock.On call
//   - status string
//   - start string
//   - end string
func (_e *MockRepositoryInterface_Expecter) GetAllTransactionByStatusDate(status interface{}, start interface{}, end interface{}) *MockRepositoryInterface_GetAllTransactionByStatusDate_Call {
	return &MockRepositoryInterface_GetAllTransactionByStatusDate_Call{Call: _e.mock.On("GetAllTransactionByStatusDate", status, start, end)}
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatusDate_Call) Run(run func(status string, start string, end string)) *MockRepositoryInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatusDate_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetAllTransactionByStatusDate_Call) RunAndReturn(run func(string, string, string) ([]Transaction, error)) *MockRepositoryInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionByDate provides a mock function with given fields: req, input
func (_m *MockRepositoryInterface) GetTransactionByDate(req FilterByDate, input FilterLimit) ([]Transaction, error) {
	ret := _m.Called(req, input)

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(FilterByDate, FilterLimit) ([]Transaction, error)); ok {
		return rf(req, input)
	}
	if rf, ok := ret.Get(0).(func(FilterByDate, FilterLimit) []Transaction); ok {
		r0 = rf(req, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(FilterByDate, FilterLimit) error); ok {
		r1 = rf(req, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetTransactionByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionByDate'
type MockRepositoryInterface_GetTransactionByDate_Call struct {
	*mock.Call
}

// GetTransactionByDate is a helper method to define mock.On call
//   - req FilterByDate
//   - input FilterLimit
func (_e *MockRepositoryInterface_Expecter) GetTransactionByDate(req interface{}, input interface{}) *MockRepositoryInterface_GetTransactionByDate_Call {
	return &MockRepositoryInterface_GetTransactionByDate_Call{Call: _e.mock.On("GetTransactionByDate", req, input)}
}

func (_c *MockRepositoryInterface_GetTransactionByDate_Call) Run(run func(req FilterByDate, input FilterLimit)) *MockRepositoryInterface_GetTransactionByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(FilterByDate), args[1].(FilterLimit))
	})
	return _c
}

func (_c *MockRepositoryInterface_GetTransactionByDate_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetTransactionByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetTransactionByDate_Call) RunAndReturn(run func(FilterByDate, FilterLimit) ([]Transaction, error)) *MockRepositoryInterface_GetTransactionByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionByStatusAndDate provides a mock function with given fields: req, input
func (_m *MockRepositoryInterface) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) ([]Transaction, error) {
	ret := _m.Called(req, input)

	var r0 []Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(FilterByStatusDate, FilterLimit) ([]Transaction, error)); ok {
		return rf(req, input)
	}
	if rf, ok := ret.Get(0).(func(FilterByStatusDate, FilterLimit) []Transaction); ok {
		r0 = rf(req, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(FilterByStatusDate, FilterLimit) error); ok {
		r1 = rf(req, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_GetTransactionByStatusAndDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionByStatusAndDate'
type MockRepositoryInterface_GetTransactionByStatusAndDate_Call struct {
	*mock.Call
}

// GetTransactionByStatusAndDate is a helper method to define mock.On call
//   - req FilterByStatusDate
//   - input FilterLimit
func (_e *MockRepositoryInterface_Expecter) GetTransactionByStatusAndDate(req interface{}, input interface{}) *MockRepositoryInterface_GetTransactionByStatusAndDate_Call {
	return &MockRepositoryInterface_GetTransactionByStatusAndDate_Call{Call: _e.mock.On("GetTransactionByStatusAndDate", req, input)}
}

func (_c *MockRepositoryInterface_GetTransactionByStatusAndDate_Call) Run(run func(req FilterByStatusDate, input FilterLimit)) *MockRepositoryInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(FilterByStatusDate), args[1].(FilterLimit))
	})
	return _c
}

func (_c *MockRepositoryInterface_GetTransactionByStatusAndDate_Call) Return(_a0 []Transaction, _a1 error) *MockRepositoryInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_GetTransactionByStatusAndDate_Call) RunAndReturn(run func(FilterByStatusDate, FilterLimit) ([]Transaction, error)) *MockRepositoryInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRepositoryInterface creates a new instance of MockRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
