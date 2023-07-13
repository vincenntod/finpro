package account

import (
	"bytes"
	"encoding/json"
	"golang/module/account/dto"
	"golang/module/account/mocks"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_GetDataUser(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/data-user", nil)
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					GetDataUser().
					Return(&dto.MessageResponse{
						Status:  "OK",
						Message: "Success Get Data",
						Code:    200,
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Name:  "Vincen",
							Role:  "admin",
							Email: "admin@example.com",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Status:  "OK",
					Message: "Success Get Data",
					Code:    200,
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Name:  "Vincen",
						Role:  "admin",
						Email: "admin@example.com",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/data-user", h.GetDataUser)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_GetDataUserById(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/data-user/1", nil)
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					GetDataUserById("").
					Return(&dto.MessageResponse{
						Status:  "OK",
						Message: "Success Get Data",
						Code:    200,
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Name:  "Vincen",
							Role:  "admin",
							Email: "admin@example.com",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Status:  "OK",
					Message: "Success Get Data",
					Code:    200,
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Name:  "Vincen",
						Role:  "admin",
						Email: "admin@example.com",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/data-user/1", h.GetDataUserById)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_CreateAccount(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(dto.CreateRequest{
					Name:     "Vincen",
					Role:     "admin",
					Email:    "admin@example.com",
					Phone:    "123",
					Password: "123456"})
				request, _ := http.NewRequest(http.MethodPost, "/create-user", bytes.NewReader(body))
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					CreateAccount(&dto.CreateRequest{
						Name:     "Vincen",
						Role:     "admin",
						Email:    "admin@example.com",
						Phone:    "123",
						Password: "123456",
					}).
					Return(&dto.MessageResponse{
						Status:  "OK",
						Message: "Success Create Data",
						Code:    200,
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Name:  "Vincen",
							Role:  "admin",
							Email: "admin@example.com",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Status:  "OK",
					Message: "Success Create Data",
					Code:    200,
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Name:  "Vincen",
						Role:  "admin",
						Email: "admin@example.com",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/create-user", h.CreateAccount)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_EditDataUser(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(dto.EditDataUserRequest{
					Id:       1,
					Name:     "Vincen",
					Role:     "admin",
					Phone:    "123",
					Password: "123456"})
				request, _ := http.NewRequest(http.MethodPost, "/data-user/1", bytes.NewReader(body))
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					EditDataUser("", &dto.EditDataUserRequest{
						Id:       1,
						Name:     "Vincen",
						Role:     "admin",
						Phone:    "123",
						Password: "123456",
					}).
					Return(&dto.MessageResponse{
						Status:  "OK",
						Message: "Success Edit Data",
						Code:    200,
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Name:  "Vincen",
							Role:  "admin",
							Email: "admin@example.com",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Status:  "OK",
					Message: "Success Edit Data",
					Code:    200,
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Name:  "Vincen",
						Role:  "admin",
						Email: "admin@example.com",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/data-user/1", h.EditDataUser)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_DeleteDataUser(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {

				request, _ := http.NewRequest(http.MethodDelete, "/data-user/1", nil)
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					DeleteDataUser("").
					Return(&dto.MessageResponse{
						Status:  "OK",
						Message: "Success Delete Data",
						Code:    200,
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Status:  "OK",
					Message: "Success Delete Data",
					Code:    200,
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.DELETE("/data-user/1", h.DeleteDataUser)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_Login(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(&dto.LoginResponseRequest{
					Email:    "maxwelvincen@gmail.com",
					Password: "123456"})
				request, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					Login(&dto.LoginResponseRequest{
						Email:    "maxwelvincen@gmail.com",
						Password: "123456"}).
					Return("123", &dto.LoginResponse{
						Code:    200,
						Status:  "OK",
						Message: "Login Success",
						Data: []dto.LoginResponseWithToken{{
							Id:    1,
							Name:  "Vincen",
							Email: "maxwelvincen@gmail.com",
							Role:  "admin",
							Token: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.LoginResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.LoginResponse{
					Code:    200,
					Status:  "OK",
					Message: "Login Success",
					Data: []dto.LoginResponseWithToken{{
						Id:    1,
						Name:  "Vincen",
						Email: "maxwelvincen@gmail.com",
						Role:  "admin",
						Token: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/login", h.Login)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_SendEmail(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/send-email/maxwelvincen@gmail.com", nil)
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					SendEmail("").
					Return(&dto.MessageResponse{
						Code:    200,
						Status:  "OK",
						Message: "Success Send Email",
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Name:  "Vincen",
							Email: "maxwelvincen@gmail.com",
							Role:  "admin",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Code:    200,
					Status:  "OK",
					Message: "Success Send Email",
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Name:  "Vincen",
						Email: "maxwelvincen@gmail.com",
						Role:  "admin",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/send-email/maxwelvincen@gmail.com", h.SendEmail)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_SendEmailRegister(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodPost, "/send-email-registration/maxwelvincen@gmail.com", nil)
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					SendEmailRegister("").
					Return(&dto.MessageResponse{
						Code:    200,
						Status:  "OK",
						Message: "Success Send Email",
						Data:    nil,
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Code:    200,
					Status:  "OK",
					Message: "Success Send Email",
					Data:    nil,
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/send-email-registration/maxwelvincen@gmail.com", h.SendEmailRegister)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_CompareVerificationCode(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(&dto.VerificationCodeRequest{
					Email: "maxwelvincen@gmail.com",
					Code:  "1234"})
				request, _ := http.NewRequest(http.MethodPost, "/compare-verification-code", bytes.NewReader(body))
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					CompareVerificationCode(&dto.VerificationCodeRequest{
						Email: "maxwelvincen@gmail.com",
						Code:  "1234"}).
					Return(&dto.MessageResponse{
						Code:    200,
						Status:  "OK",
						Message: "Success Verification Code",
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Email: "maxwelvincen@gmail.com",
							Name:  "Vincen",
							Role:  "Admin",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Code:    200,
					Status:  "OK",
					Message: "Success Verification Code",
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Email: "maxwelvincen@gmail.com",
						Name:  "Vincen",
						Role:  "Admin",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.POST("/compare-verification-code", h.CompareVerificationCode)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_EditPassword(t *testing.T) {

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() RequestHandler
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Success",
			expectedStatusCode: 200,
			makeRequest: func() *http.Request {
				body, _ := json.Marshal(&dto.EditDataUserRequest{
					Id:       1,
					Password: "123456",
				})
				request, _ := http.NewRequest(http.MethodPut, "/edit-password/6647/?id=1", bytes.NewReader(body))
				return request
			},
			makeFields: func() RequestHandler {
				ctrl := gomock.NewController(t)
				mockController := mocks.NewMockControllerInterface(ctrl)
				mockController.EXPECT().
					EditPassword("1", "6647", &dto.EditDataUserRequest{
						Id:       1,
						Password: "123456"}).
					Return(&dto.MessageResponse{
						Code:    200,
						Status:  "OK",
						Message: "Success Update Password",
						Data: []dto.AccountItemResponse{{
							Id:    1,
							Email: "maxwelvincen@gmail.com",
							Name:  "Vincen",
							Role:  "Admin",
							Phone: "123",
						},
						},
					}, nil).
					Times(1)
				return RequestHandler{Ctrl: mockController}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := &dto.MessageResponse{}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, &dto.MessageResponse{
					Code:    200,
					Status:  "OK",
					Message: "Success Update Password",
					Data: []dto.AccountItemResponse{{
						Id:    1,
						Email: "maxwelvincen@gmail.com",
						Name:  "Vincen",
						Role:  "Admin",
						Phone: "123",
					},
					},
				}, res, i...)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}
			statusCode, body := mocks.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.PUT("/edit-password/:verification-code/", h.EditPassword)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}
