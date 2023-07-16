// Code generated by MockGen. DO NOT EDIT.
// Source: Repository.go

// Package mock_account is a generated GoMock package.
package mocks

import (
        account "golang/module/account/dto"
        "golang/module/account/entities"
        reflect "reflect"

        gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
        ctrl     *gomock.Controller
        recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
        mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
        mock := &MockRepositoryInterface{ctrl: ctrl}
        mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
        return m.recorder
}

// CompareVerificationCode mocks base method.
func (m *MockRepositoryInterface) CompareVerificationCode(verificationCode *account.VerificationCodeRequest) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CompareVerificationCode", verificationCode)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CompareVerificationCode indicates an expected call of CompareVerificationCode.
func (mr *MockRepositoryInterfaceMockRecorder) CompareVerificationCode(verificationCode interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompareVerificationCode", reflect.TypeOf((*MockRepositoryInterface)(nil).CompareVerificationCode), verificationCode)
}

// CreateAccount mocks base method.
func (m *MockRepositoryInterface) CreateAccount(req *entities.Account) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAccount", req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockRepositoryInterfaceMockRecorder) CreateAccount(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateAccount), req)
}

// DeleteDataUser mocks base method.
func (m *MockRepositoryInterface) DeleteDataUser(id string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteDataUser", id)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// DeleteDataUser indicates an expected call of DeleteDataUser.
func (mr *MockRepositoryInterfaceMockRecorder) DeleteDataUser(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDataUser", reflect.TypeOf((*MockRepositoryInterface)(nil).DeleteDataUser), id)
}

// EditDataUser mocks base method.
func (m *MockRepositoryInterface) EditDataUser(id string, req *entities.Account) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditDataUser", id, req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditDataUser indicates an expected call of EditDataUser.
func (mr *MockRepositoryInterfaceMockRecorder) EditDataUser(id, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDataUser", reflect.TypeOf((*MockRepositoryInterface)(nil).EditDataUser), id, req)
}

// EditPassword mocks base method.
func (m *MockRepositoryInterface) EditPassword(id string, req *entities.Account) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "EditPassword", id, req)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// EditPassword indicates an expected call of EditPassword.
func (mr *MockRepositoryInterfaceMockRecorder) EditPassword(id, req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditPassword", reflect.TypeOf((*MockRepositoryInterface)(nil).EditPassword), id, req)
}

// GetDataUser mocks base method.
func (m *MockRepositoryInterface) GetDataUser() ([]entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUser")
        ret0, _ := ret[0].([]entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUser indicates an expected call of GetDataUser.
func (mr *MockRepositoryInterfaceMockRecorder) GetDataUser() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUser", reflect.TypeOf((*MockRepositoryInterface)(nil).GetDataUser))
}

// GetDataUserById mocks base method.
func (m *MockRepositoryInterface) GetDataUserById(id string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetDataUserById", id)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetDataUserById indicates an expected call of GetDataUserById.
func (mr *MockRepositoryInterfaceMockRecorder) GetDataUserById(id interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataUserById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetDataUserById), id)
}

// Login mocks base method.
func (m *MockRepositoryInterface) Login(req *entities.Account) (string, entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Login", req)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(entities.Account)
        ret2, _ := ret[2].(error)
        return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockRepositoryInterfaceMockRecorder) Login(req interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRepositoryInterface)(nil).Login), req)
}

// SendEmail mocks base method.
func (m *MockRepositoryInterface) SendEmail(email string) (entities.Account, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SendEmail", email)
        ret0, _ := ret[0].(entities.Account)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SendEmail indicates an expected call of SendEmail.
func (mr *MockRepositoryInterfaceMockRecorder) SendEmail(email interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEmail", reflect.TypeOf((*MockRepositoryInterface)(nil).SendEmail), email)
}