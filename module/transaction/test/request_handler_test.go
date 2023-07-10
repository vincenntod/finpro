package test

import (
	"encoding/json"
	"errors"
	"golang/helper"
	"golang/module/transaction"
	mocks "golang/module/transaction/mocks/module/transaction"
	"net/http"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_GetAllLimit(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:               "Test Case Get All Limit Error",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-transactions-limit/1", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewControllerInterface(t)
				err := errors.New("Error")
				mockController.EXPECT().GetAllLimit(transaction.FilterLimit{1, 100}).Return(&transaction.GetAllResponseDataTransaction{}, err, 6736).Once()
				return fields{
					ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {

				res := gin.H{"message": "Error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "Error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := transaction.RequestHandler{
				Ctrl: f.ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-transactions-limit/:id", h.GetAllLimit)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_GetAllTransaction(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Test Case Get All Transaction Error",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest("GET", "/get-transactions/", nil)
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewControllerInterface(t)
				err := errors.New("Error")
				mockController.EXPECT().GetAllTransaction().Return(&transaction.GetAllResponseDataTransaction{}, err).Once()
				return fields{
					ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {

				res := gin.H{"message": "Error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "Error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := transaction.RequestHandler{
				Ctrl: f.ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-transactions/", h.GetAllTransaction)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_GetAllTransactionByDate(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:               "Test Case Get All Transaction By Date",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-TransactionDate/18-04-2023/19-04-2023/", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewControllerInterface(t)
				err := errors.New("Error")
				mockController.EXPECT().GetAllTransactionByDate("2023-04-18", "2023-04-19").Return(&transaction.GetAllResponseDataTransaction{}, err).Once()
				return fields{
					ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {

				res := gin.H{"message": "Error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "Error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := transaction.RequestHandler{
				Ctrl: f.ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-TransactionDate/:start/:end/", h.GetAllTransactionByDate)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:               "Test Case Get All Transaction By Status Error",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-transaction-status/SUCCESS/", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewControllerInterface(t)
				err := errors.New("Error")
				mockController.EXPECT().GetAllTransactionByStatus("SUCCESS").Return(&transaction.GetAllResponseDataTransaction{}, err).Once()
				return fields{
					ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {

				res := gin.H{"message": "Error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "Error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := transaction.RequestHandler{
				Ctrl: f.ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-transaction-status/:status/", h.GetAllTransactionByStatus)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}

func TestRequestHandler_GetAllTransactionByStatusDate(t *testing.T) {
	type fields struct {
		ctrl transaction.ControllerInterface
	}
	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:               "Test Case Get All Transaction By Status Error",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-TransactionStatusDate/SUCCESS/18-04-2023/19-04-2023/", nil)
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewControllerInterface(t)
				err := errors.New("Error")
				mockController.EXPECT().GetAllTransactionByStatusDate("SUCCESS", "2023-04-18", "2023-04-19").Return(&transaction.GetAllResponseDataTransaction{}, err).Once()
				return fields{
					ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {

				res := gin.H{"message": "Error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "Error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := transaction.RequestHandler{
				Ctrl: f.ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-TransactionStatusDate/:status/:start/:end/", h.GetAllTransactionByStatusDate)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}

		})
	}
}
