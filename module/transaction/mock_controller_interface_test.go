// Code generated by mockery v2.30.1. DO NOT EDIT.

package transaction

import mock "github.com/stretchr/testify/mock"

// MockControllerInterface is an autogenerated mock type for the ControllerInterface type
type MockControllerInterface struct {
	mock.Mock
}

type MockControllerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockControllerInterface) EXPECT() *MockControllerInterface_Expecter {
	return &MockControllerInterface_Expecter{mock: &_m.Mock}
}

// GetAllTransaction provides a mock function with given fields:
func (_m *MockControllerInterface) GetAllTransaction() (*GetAllResponseDataTransaction, error) {
	ret := _m.Called()

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func() (*GetAllResponseDataTransaction, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *GetAllResponseDataTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetAllTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransaction'
type MockControllerInterface_GetAllTransaction_Call struct {
	*mock.Call
}

// GetAllTransaction is a helper method to define mock.On call
func (_e *MockControllerInterface_Expecter) GetAllTransaction() *MockControllerInterface_GetAllTransaction_Call {
	return &MockControllerInterface_GetAllTransaction_Call{Call: _e.mock.On("GetAllTransaction")}
}

func (_c *MockControllerInterface_GetAllTransaction_Call) Run(run func()) *MockControllerInterface_GetAllTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockControllerInterface_GetAllTransaction_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetAllTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetAllTransaction_Call) RunAndReturn(run func() (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetAllTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByDate provides a mock function with given fields: start, end
func (_m *MockControllerInterface) GetAllTransactionByDate(start string, end string) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(start, end)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*GetAllResponseDataTransaction, error)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string) *GetAllResponseDataTransaction); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetAllTransactionByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByDate'
type MockControllerInterface_GetAllTransactionByDate_Call struct {
	*mock.Call
}

// GetAllTransactionByDate is a helper method to define mock.On call
//   - start string
//   - end string
func (_e *MockControllerInterface_Expecter) GetAllTransactionByDate(start interface{}, end interface{}) *MockControllerInterface_GetAllTransactionByDate_Call {
	return &MockControllerInterface_GetAllTransactionByDate_Call{Call: _e.mock.On("GetAllTransactionByDate", start, end)}
}

func (_c *MockControllerInterface_GetAllTransactionByDate_Call) Run(run func(start string, end string)) *MockControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByDate_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByDate_Call) RunAndReturn(run func(string, string) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatus provides a mock function with given fields: status
func (_m *MockControllerInterface) GetAllTransactionByStatus(status string) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(status)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*GetAllResponseDataTransaction, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) *GetAllResponseDataTransaction); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetAllTransactionByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatus'
type MockControllerInterface_GetAllTransactionByStatus_Call struct {
	*mock.Call
}

// GetAllTransactionByStatus is a helper method to define mock.On call
//   - status string
func (_e *MockControllerInterface_Expecter) GetAllTransactionByStatus(status interface{}) *MockControllerInterface_GetAllTransactionByStatus_Call {
	return &MockControllerInterface_GetAllTransactionByStatus_Call{Call: _e.mock.On("GetAllTransactionByStatus", status)}
}

func (_c *MockControllerInterface_GetAllTransactionByStatus_Call) Run(run func(status string)) *MockControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByStatus_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByStatus_Call) RunAndReturn(run func(string) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatusDate provides a mock function with given fields: status, start, end
func (_m *MockControllerInterface) GetAllTransactionByStatusDate(status string, start string, end string) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(status, start, end)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*GetAllResponseDataTransaction, error)); ok {
		return rf(status, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *GetAllResponseDataTransaction); ok {
		r0 = rf(status, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(status, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetAllTransactionByStatusDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatusDate'
type MockControllerInterface_GetAllTransactionByStatusDate_Call struct {
	*mock.Call
}

// GetAllTransactionByStatusDate is a helper method to define mock.On call
//   - status string
//   - start string
//   - end string
func (_e *MockControllerInterface_Expecter) GetAllTransactionByStatusDate(status interface{}, start interface{}, end interface{}) *MockControllerInterface_GetAllTransactionByStatusDate_Call {
	return &MockControllerInterface_GetAllTransactionByStatusDate_Call{Call: _e.mock.On("GetAllTransactionByStatusDate", status, start, end)}
}

func (_c *MockControllerInterface_GetAllTransactionByStatusDate_Call) Run(run func(status string, start string, end string)) *MockControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByStatusDate_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetAllTransactionByStatusDate_Call) RunAndReturn(run func(string, string, string) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransaction provides a mock function with given fields: req
func (_m *MockControllerInterface) GetTransaction(req *FilterByStatusDate) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(req)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*FilterByStatusDate) (*GetAllResponseDataTransaction, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*FilterByStatusDate) *GetAllResponseDataTransaction); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*FilterByStatusDate) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransaction'
type MockControllerInterface_GetTransaction_Call struct {
	*mock.Call
}

// GetTransaction is a helper method to define mock.On call
//   - req *FilterByStatusDate
func (_e *MockControllerInterface_Expecter) GetTransaction(req interface{}) *MockControllerInterface_GetTransaction_Call {
	return &MockControllerInterface_GetTransaction_Call{Call: _e.mock.On("GetTransaction", req)}
}

func (_c *MockControllerInterface_GetTransaction_Call) Run(run func(req *FilterByStatusDate)) *MockControllerInterface_GetTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*FilterByStatusDate))
	})
	return _c
}

func (_c *MockControllerInterface_GetTransaction_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetTransaction_Call) RunAndReturn(run func(*FilterByStatusDate) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionByDate provides a mock function with given fields: req, input
func (_m *MockControllerInterface) GetTransactionByDate(req FilterByDate, input FilterLimit) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(req, input)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(FilterByDate, FilterLimit) (*GetAllResponseDataTransaction, error)); ok {
		return rf(req, input)
	}
	if rf, ok := ret.Get(0).(func(FilterByDate, FilterLimit) *GetAllResponseDataTransaction); ok {
		r0 = rf(req, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(FilterByDate, FilterLimit) error); ok {
		r1 = rf(req, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetTransactionByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionByDate'
type MockControllerInterface_GetTransactionByDate_Call struct {
	*mock.Call
}

// GetTransactionByDate is a helper method to define mock.On call
//   - req FilterByDate
//   - input FilterLimit
func (_e *MockControllerInterface_Expecter) GetTransactionByDate(req interface{}, input interface{}) *MockControllerInterface_GetTransactionByDate_Call {
	return &MockControllerInterface_GetTransactionByDate_Call{Call: _e.mock.On("GetTransactionByDate", req, input)}
}

func (_c *MockControllerInterface_GetTransactionByDate_Call) Run(run func(req FilterByDate, input FilterLimit)) *MockControllerInterface_GetTransactionByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(FilterByDate), args[1].(FilterLimit))
	})
	return _c
}

func (_c *MockControllerInterface_GetTransactionByDate_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetTransactionByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetTransactionByDate_Call) RunAndReturn(run func(FilterByDate, FilterLimit) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetTransactionByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTransactionByStatusAndDate provides a mock function with given fields: req, input
func (_m *MockControllerInterface) GetTransactionByStatusAndDate(req FilterByStatusDate, input FilterLimit) (*GetAllResponseDataTransaction, error) {
	ret := _m.Called(req, input)

	var r0 *GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(FilterByStatusDate, FilterLimit) (*GetAllResponseDataTransaction, error)); ok {
		return rf(req, input)
	}
	if rf, ok := ret.Get(0).(func(FilterByStatusDate, FilterLimit) *GetAllResponseDataTransaction); ok {
		r0 = rf(req, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(FilterByStatusDate, FilterLimit) error); ok {
		r1 = rf(req, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockControllerInterface_GetTransactionByStatusAndDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTransactionByStatusAndDate'
type MockControllerInterface_GetTransactionByStatusAndDate_Call struct {
	*mock.Call
}

// GetTransactionByStatusAndDate is a helper method to define mock.On call
//   - req FilterByStatusDate
//   - input FilterLimit
func (_e *MockControllerInterface_Expecter) GetTransactionByStatusAndDate(req interface{}, input interface{}) *MockControllerInterface_GetTransactionByStatusAndDate_Call {
	return &MockControllerInterface_GetTransactionByStatusAndDate_Call{Call: _e.mock.On("GetTransactionByStatusAndDate", req, input)}
}

func (_c *MockControllerInterface_GetTransactionByStatusAndDate_Call) Run(run func(req FilterByStatusDate, input FilterLimit)) *MockControllerInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(FilterByStatusDate), args[1].(FilterLimit))
	})
	return _c
}

func (_c *MockControllerInterface_GetTransactionByStatusAndDate_Call) Return(_a0 *GetAllResponseDataTransaction, _a1 error) *MockControllerInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockControllerInterface_GetTransactionByStatusAndDate_Call) RunAndReturn(run func(FilterByStatusDate, FilterLimit) (*GetAllResponseDataTransaction, error)) *MockControllerInterface_GetTransactionByStatusAndDate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockControllerInterface creates a new instance of MockControllerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockControllerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockControllerInterface {
	mock := &MockControllerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
