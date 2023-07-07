// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	exportcsv "golang/module/exportcsv"

	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

type Controller_Expecter struct {
	mock *mock.Mock
}

func (_m *Controller) EXPECT() *Controller_Expecter {
	return &Controller_Expecter{mock: &_m.Mock}
}

// ExportCSV provides a mock function with given fields: req, url
func (_m *Controller) ExportCSV(req *exportcsv.ExportCSVRequest, url *exportcsv.UrlRequest) ([][]string, error) {
	ret := _m.Called(req, url)

	var r0 [][]string
	var r1 error
	if rf, ok := ret.Get(0).(func(*exportcsv.ExportCSVRequest, *exportcsv.UrlRequest) ([][]string, error)); ok {
		return rf(req, url)
	}
	if rf, ok := ret.Get(0).(func(*exportcsv.ExportCSVRequest, *exportcsv.UrlRequest) [][]string); ok {
		r0 = rf(req, url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*exportcsv.ExportCSVRequest, *exportcsv.UrlRequest) error); ok {
		r1 = rf(req, url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Controller_ExportCSV_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExportCSV'
type Controller_ExportCSV_Call struct {
	*mock.Call
}

// ExportCSV is a helper method to define mock.On call
//   - req *exportcsv.ExportCSVRequest
//   - url *exportcsv.UrlRequest
func (_e *Controller_Expecter) ExportCSV(req interface{}, url interface{}) *Controller_ExportCSV_Call {
	return &Controller_ExportCSV_Call{Call: _e.mock.On("ExportCSV", req, url)}
}

func (_c *Controller_ExportCSV_Call) Run(run func(req *exportcsv.ExportCSVRequest, url *exportcsv.UrlRequest)) *Controller_ExportCSV_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*exportcsv.ExportCSVRequest), args[1].(*exportcsv.UrlRequest))
	})
	return _c
}

func (_c *Controller_ExportCSV_Call) Return(_a0 [][]string, _a1 error) *Controller_ExportCSV_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Controller_ExportCSV_Call) RunAndReturn(run func(*exportcsv.ExportCSVRequest, *exportcsv.UrlRequest) ([][]string, error)) *Controller_ExportCSV_Call {
	_c.Call.Return(run)
	return _c
}

// NewController creates a new instance of Controller. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewController(t interface {
	mock.TestingT
	Cleanup(func())
}) *Controller {
	mock := &Controller{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
