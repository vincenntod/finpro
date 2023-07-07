// Code generated by MockGen. DO NOT EDIT.
// Source: Use-case.go

// Package mock_account is a generated GoMock package.
package mocks

import (
        account "golang/module/account"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockUseCaseInterface is a mock of UseCaseInterface interface.
type MockUseCaseInterface struct {
        ctrl     *gomock.Controller
        recorder *MockUseCaseInterfaceMockRecorder
}

// MockUseCaseInterfaceMockRecorder is the mock recorder for MockUseCaseInterface.
type MockUseCaseInterfaceMockRecorder struct {
        mock *MockUseCaseInterface
}

// NewMockUseCaseInterface creates a new mock instance.
func NewMockUseCaseInterface(ctrl *gomock.Controller) *MockUseCaseInterface {
        mock := &MockUseCaseInterface{ctrl: ctrl}
        mock.recorder = &MockUseCaseInterfaceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCaseInterface) EXPECT() *MockUseCaseInterfaceMockRecorder {
        return m.recorder
}

// CompareVerificationCode mocks base method.
func (m *MockUseCaseInterface) CompareVerificationCode(verificationCode *account.VerificationCodeRequest) (account.Account, error) {   
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CompareVerificationCode", verificationCode)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CompareVerificationCode indicates an expected call of CompareVerificationCode.
func (mr *MockUseCaseInterfaceMockRecorder) CompareVerificationCode(verificationCode interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareVerificationCode", reflect.TypeOf((*MockUseCaseInterface)(nil).CompareVerificationCode), verificationCode)
}

// CreateAccount mocks base method.
func (m *MockUseCaseInterface) CreateAccount(req *account.Account) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAccount", req)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockUseCaseInterfaceMockRecorder) CreateAccount(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockUseCaseInterface)(nil).CreateAccount), req)
}

// DeleteDataUser mocks base method.
func (m *MockUseCaseInterface) DeleteDataUser(id string) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteDataUser", id)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// DeleteDataUser indicates an expected call of DeleteDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) DeleteDataUser(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).DeleteDataUser), id)
}

// EditDataUser mocks base method.
func (m *MockUseCaseInterface) EditDataUser(id string, req *account.Account) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditDataUser", id, req)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditDataUser indicates an expected call of EditDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) EditDataUser(id, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).EditDataUser), id, req)
}

// EditPassword mocks base method.
func (m *MockUseCaseInterface) EditPassword(code string, req *account.Account) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditPassword", code, req)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditPassword indicates an expected call of EditPassword.
func (mr *MockUseCaseInterfaceMockRecorder) EditPassword(code, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPassword", reflect.TypeOf((*MockUseCaseInterface)(nil).EditPassword), code, req)
}

// GetDataUser mocks base method.
func (m *MockUseCaseInterface) GetDataUser() ([]account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUser")
        ret0, _ := ret[0].([]account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUser indicates an expected call of GetDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) GetDataUser() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).GetDataUser)) 
}

// GetDataUserById mocks base method.
func (m *MockUseCaseInterface) GetDataUserById(id int) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUserById", id)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUserById indicates an expected call of GetDataUserById.
func (mr *MockUseCaseInterfaceMockRecorder) GetDataUserById(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUserById", reflect.TypeOf((*MockUseCaseInterface)(nil).GetDataUserById), id)
}

// Login mocks base method.
func (m *MockUseCaseInterface) Login(req *account.Account) (string, account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Login", req)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(account.Account)
        ret2, _ := ret[2].(error)
        return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockUseCaseInterfaceMockRecorder) Login(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCaseInterface)(nil).Login), req)        
}

// SendEmail mocks base method.
func (m *MockUseCaseInterface) SendEmail(email string) (account.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendEmail", email)
        ret0, _ := ret[0].(account.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockUseCaseInterfaceMockRecorder) SendEmail(email interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockUseCaseInterface)(nil).SendEmail), email)
}