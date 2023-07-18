// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	"golang/module/transaction/dto"

	mock "github.com/stretchr/testify/mock"
)

// ControllerInterface is an autogenerated mock type for the ControllerInterface type
type ControllerInterface struct {
	mock.Mock
}

type ControllerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *ControllerInterface) EXPECT() *ControllerInterface_Expecter {
	return &ControllerInterface_Expecter{mock: &_m.Mock}
}


// GetAllTransaction provides a mock function with given fields: page, size
func (_m *ControllerInterface) GetAllTransaction(page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(page, size)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(page, size)
	}
	if rf, ok := ret.Get(0).(func(int, int) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransaction'
type ControllerInterface_GetAllTransaction_Call struct {
	*mock.Call
}

// GetAllTransaction is a helper method to define mock.On call
//   - page int
//   - size int
func (_e *ControllerInterface_Expecter) GetAllTransaction(page interface{}, size interface{}) *ControllerInterface_GetAllTransaction_Call {
	return &ControllerInterface_GetAllTransaction_Call{Call: _e.mock.On("GetAllTransaction", page, size)}
}

func (_c *ControllerInterface_GetAllTransaction_Call) Run(run func(page int, size int)) *ControllerInterface_GetAllTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransaction_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransaction_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransaction_Call) RunAndReturn(run func(int, int) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByDate provides a mock function with given fields: start, end, page, size
func (_m *ControllerInterface) GetAllTransactionByDate(start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(start, end, page, size)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(start, end, page, size)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(start, end, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) error); ok {
		r1 = rf(start, end, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByDate'
type ControllerInterface_GetAllTransactionByDate_Call struct {
	*mock.Call
}

// GetAllTransactionByDate is a helper method to define mock.On call
//   - start string
//   - end string
//   - page int
//   - size int
func (_e *ControllerInterface_Expecter) GetAllTransactionByDate(start interface{}, end interface{}, page interface{}, size interface{}) *ControllerInterface_GetAllTransactionByDate_Call {
	return &ControllerInterface_GetAllTransactionByDate_Call{Call: _e.mock.On("GetAllTransactionByDate", start, end, page, size)}
}

func (_c *ControllerInterface_GetAllTransactionByDate_Call) Run(run func(start string, end string, page int, size int)) *ControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(int), args[3].(int))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByDate_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByDate_Call) RunAndReturn(run func(string, string, int, int) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByDateNoLimit provides a mock function with given fields: start, end
func (_m *ControllerInterface) GetAllTransactionByDateNoLimit(start string, end string) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(start, end)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByDateNoLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByDateNoLimit'
type ControllerInterface_GetAllTransactionByDateNoLimit_Call struct {
	*mock.Call
}

// GetAllTransactionByDateNoLimit is a helper method to define mock.On call
//   - start string
//   - end string
func (_e *ControllerInterface_Expecter) GetAllTransactionByDateNoLimit(start interface{}, end interface{}) *ControllerInterface_GetAllTransactionByDateNoLimit_Call {
	return &ControllerInterface_GetAllTransactionByDateNoLimit_Call{Call: _e.mock.On("GetAllTransactionByDateNoLimit", start, end)}
}

func (_c *ControllerInterface_GetAllTransactionByDateNoLimit_Call) Run(run func(start string, end string)) *ControllerInterface_GetAllTransactionByDateNoLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByDateNoLimit_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByDateNoLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByDateNoLimit_Call) RunAndReturn(run func(string, string) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByDateNoLimit_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByRequest provides a mock function with given fields: req
func (_m *ControllerInterface) GetAllTransactionByRequest(req *dto.Request) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(req)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*dto.Request) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*dto.Request) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*dto.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByRequest'
type ControllerInterface_GetAllTransactionByRequest_Call struct {
	*mock.Call
}

// GetAllTransactionByRequest is a helper method to define mock.On call
//   - req *transaction.Request
func (_e *ControllerInterface_Expecter) GetAllTransactionByRequest(req interface{}) *ControllerInterface_GetAllTransactionByRequest_Call {
	return &ControllerInterface_GetAllTransactionByRequest_Call{Call: _e.mock.On("GetAllTransactionByRequest", req)}
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) Run(run func(req *dto.Request)) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*dto.Request))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByRequest_Call) RunAndReturn(run func(*dto.Request) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatus provides a mock function with given fields: status, page, size
func (_m *ControllerInterface) GetAllTransactionByStatus(status string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(status, page, size)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int, int) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(status, page, size)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(status, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) error); ok {
		r1 = rf(status, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatus'
type ControllerInterface_GetAllTransactionByStatus_Call struct {
	*mock.Call
}

// GetAllTransactionByStatus is a helper method to define mock.On call
//   - status string
//   - page int
//   - size int
func (_e *ControllerInterface_Expecter) GetAllTransactionByStatus(status interface{}, page interface{}, size interface{}) *ControllerInterface_GetAllTransactionByStatus_Call {
	return &ControllerInterface_GetAllTransactionByStatus_Call{Call: _e.mock.On("GetAllTransactionByStatus", status, page, size)}
}

func (_c *ControllerInterface_GetAllTransactionByStatus_Call) Run(run func(status string, page int, size int)) *ControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(int), args[2].(int))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatus_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatus_Call) RunAndReturn(run func(string, int, int) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByStatus_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatusDate provides a mock function with given fields: status, start, end, page, size
func (_m *ControllerInterface) GetAllTransactionByStatusDate(status string, start string, end string, page int, size int) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(status, start, end, page, size)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(status, start, end, page, size)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, int, int) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(status, start, end, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, int, int) error); ok {
		r1 = rf(status, start, end, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByStatusDate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatusDate'
type ControllerInterface_GetAllTransactionByStatusDate_Call struct {
	*mock.Call
}

// GetAllTransactionByStatusDate is a helper method to define mock.On call
//   - status string
//   - start string
//   - end string
//   - page int
//   - size int
func (_e *ControllerInterface_Expecter) GetAllTransactionByStatusDate(status interface{}, start interface{}, end interface{}, page interface{}, size interface{}) *ControllerInterface_GetAllTransactionByStatusDate_Call {
	return &ControllerInterface_GetAllTransactionByStatusDate_Call{Call: _e.mock.On("GetAllTransactionByStatusDate", status, start, end, page, size)}
}

func (_c *ControllerInterface_GetAllTransactionByStatusDate_Call) Run(run func(status string, start string, end string, page int, size int)) *ControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusDate_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusDate_Call) RunAndReturn(run func(string, string, string, int, int) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByStatusDate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatusDateNoLimit provides a mock function with given fields: status, start, end
func (_m *ControllerInterface) GetAllTransactionByStatusDateNoLimit(status string, start string, end string) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(status, start, end)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(status, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(status, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(status, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatusDateNoLimit'
type ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call struct {
	*mock.Call
}

// GetAllTransactionByStatusDateNoLimit is a helper method to define mock.On call
//   - status string
//   - start string
//   - end string
func (_e *ControllerInterface_Expecter) GetAllTransactionByStatusDateNoLimit(status interface{}, start interface{}, end interface{}) *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call {
	return &ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call{Call: _e.mock.On("GetAllTransactionByStatusDateNoLimit", status, start, end)}
}

func (_c *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call) Run(run func(status string, start string, end string)) *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call) RunAndReturn(run func(string, string, string) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByStatusDateNoLimit_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionByStatusNoLimit provides a mock function with given fields: status
func (_m *ControllerInterface) GetAllTransactionByStatusNoLimit(status string) (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called(status)

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf(status)
	}
	if rf, ok := ret.Get(0).(func(string) *dto.GetAllResponseDataTransaction); ok {
		r0 = rf(status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionByStatusNoLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionByStatusNoLimit'
type ControllerInterface_GetAllTransactionByStatusNoLimit_Call struct {
	*mock.Call
}

// GetAllTransactionByStatusNoLimit is a helper method to define mock.On call
//   - status string
func (_e *ControllerInterface_Expecter) GetAllTransactionByStatusNoLimit(status interface{}) *ControllerInterface_GetAllTransactionByStatusNoLimit_Call {
	return &ControllerInterface_GetAllTransactionByStatusNoLimit_Call{Call: _e.mock.On("GetAllTransactionByStatusNoLimit", status)}
}

func (_c *ControllerInterface_GetAllTransactionByStatusNoLimit_Call) Run(run func(status string)) *ControllerInterface_GetAllTransactionByStatusNoLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusNoLimit_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionByStatusNoLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionByStatusNoLimit_Call) RunAndReturn(run func(string) (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionByStatusNoLimit_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllTransactionNoLimit provides a mock function with given fields:
func (_m *ControllerInterface) GetAllTransactionNoLimit() (*dto.GetAllResponseDataTransaction, error) {
	ret := _m.Called()

	var r0 *dto.GetAllResponseDataTransaction
	var r1 error
	if rf, ok := ret.Get(0).(func() (*dto.GetAllResponseDataTransaction, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *dto.GetAllResponseDataTransaction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.GetAllResponseDataTransaction)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ControllerInterface_GetAllTransactionNoLimit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllTransactionNoLimit'
type ControllerInterface_GetAllTransactionNoLimit_Call struct {
	*mock.Call
}

// GetAllTransactionNoLimit is a helper method to define mock.On call
func (_e *ControllerInterface_Expecter) GetAllTransactionNoLimit() *ControllerInterface_GetAllTransactionNoLimit_Call {
	return &ControllerInterface_GetAllTransactionNoLimit_Call{Call: _e.mock.On("GetAllTransactionNoLimit")}
}

func (_c *ControllerInterface_GetAllTransactionNoLimit_Call) Run(run func()) *ControllerInterface_GetAllTransactionNoLimit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ControllerInterface_GetAllTransactionNoLimit_Call) Return(_a0 *dto.GetAllResponseDataTransaction, _a1 error) *ControllerInterface_GetAllTransactionNoLimit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ControllerInterface_GetAllTransactionNoLimit_Call) RunAndReturn(run func() (*dto.GetAllResponseDataTransaction, error)) *ControllerInterface_GetAllTransactionNoLimit_Call {
	_c.Call.Return(run)
	return _c
}

// NewControllerInterface creates a new instance of ControllerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewControllerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ControllerInterface {
	mock := &ControllerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

type mockConstructorTestingTNewController interface {
	mock.TestingT
	Cleanup(func())
}

func NewController(t mockConstructorTestingTNewController) *ControllerInterface {
	mock := &ControllerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
