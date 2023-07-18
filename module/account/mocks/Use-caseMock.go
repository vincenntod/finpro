// Code generated by MockGen. DO NOT EDIT.
// Source: Use-case.go

// Package mock_account is a generated GoMock package.
package mocks

import (
        dto "golang/module/account/dto"
        entities "golang/module/account/entities"
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
func (m *MockUseCaseInterface) CompareVerificationCode(verificationCode *dto.VerificationCodeRequest) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CompareVerificationCode", verificationCode)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CompareVerificationCode indicates an expected call of CompareVerificationCode.
func (mr *MockUseCaseInterfaceMockRecorder) CompareVerificationCode(verificationCode interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareVerificationCode", reflect.TypeOf((*MockUseCaseInterface)(nil).CompareVerificationCode), verificationCode)
}

// CreateAccount mocks base method.
func (m *MockUseCaseInterface) CreateAccount(req *dto.CreateRequest) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAccount", req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockUseCaseInterfaceMockRecorder) CreateAccount(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockUseCaseInterface)(nil).CreateAccount), req)        
}

// DeleteDataUser mocks base method.
func (m *MockUseCaseInterface) DeleteDataUser(id string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteDataUser", id)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// DeleteDataUser indicates an expected call of DeleteDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) DeleteDataUser(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).DeleteDataUser), id)       
}

// EditDataUser mocks base method.
func (m *MockUseCaseInterface) EditDataUser(id string, req *dto.EditDataUserRequest) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditDataUser", id, req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditDataUser indicates an expected call of EditDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) EditDataUser(id, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).EditDataUser), id, req)      
}

// EditPassword mocks base method.
func (m *MockUseCaseInterface) EditPassword(code string, req *dto.EditDataUserRequest) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditPassword", code, req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditPassword indicates an expected call of EditPassword.
func (mr *MockUseCaseInterfaceMockRecorder) EditPassword(code, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPassword", reflect.TypeOf((*MockUseCaseInterface)(nil).EditPassword), code, req)    
}

// GetDataUser mocks base method.
func (m *MockUseCaseInterface) GetDataUser() ([]entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUser")
        ret0, _ := ret[0].([]entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUser indicates an expected call of GetDataUser.
func (mr *MockUseCaseInterfaceMockRecorder) GetDataUser() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUser", reflect.TypeOf((*MockUseCaseInterface)(nil).GetDataUser))
}

// GetDataUserById mocks base method.
func (m *MockUseCaseInterface) GetDataUserById(id string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUserById", id)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUserById indicates an expected call of GetDataUserById.
func (mr *MockUseCaseInterfaceMockRecorder) GetDataUserById(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUserById", reflect.TypeOf((*MockUseCaseInterface)(nil).GetDataUserById), id)     
}

// Login mocks base method.
func (m *MockUseCaseInterface) Login(req *dto.LoginResponseRequest) (string, entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Login", req)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(entities.Account)
        ret2, _ := ret[2].(error)
        return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockUseCaseInterfaceMockRecorder) Login(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUseCaseInterface)(nil).Login), req)
}

// SendEmailForgotPassword mocks base method.
func (m *MockUseCaseInterface) SendEmailForgotPassword(email string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendEmailForgotPassword", email)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SendEmailForgotPassword indicates an expected call of SendEmailForgotPassword.
func (mr *MockUseCaseInterfaceMockRecorder) SendEmailForgotPassword(email interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmailForgotPassword", reflect.TypeOf((*MockUseCaseInterface)(nil).SendEmailForgotPassword), email)
}

// SendEmailRegistration mocks base method.
func (m *MockUseCaseInterface) SendEmailRegistration(email string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendEmailRegistration", email)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SendEmailRegistration indicates an expected call of SendEmailRegistration.
func (mr *MockUseCaseInterfaceMockRecorder) SendEmailRegistration(email interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmailRegistration", reflect.TypeOf((*MockUseCaseInterface)(nil).SendEmailRegistration), email)
}