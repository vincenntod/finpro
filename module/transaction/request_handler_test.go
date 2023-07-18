package transaction

import (
	"encoding/json"
	"errors"
	"golang/module/transaction/dto"
	"golang/module/transaction/helper"
	mocks "golang/module/transaction/mocks"

	"net/http"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_GetAllTransactionByRequestLimit(t *testing.T) {

	type fields struct {
		Ctrl ControllerInterface
	}

	tests := []struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}{
		{
			name:               "Error",
			expectedStatusCode: 500,
			makeRequest: func() *http.Request {
				req, _ := http.NewRequest(http.MethodGet, "/get-transaction-by-status", nil)
				return req
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				err := errors.New("error")
				mockController.EXPECT().GetAllTransactionByRequest(&dto.Request{}).Return(&dto.GetAllResponseDataTransaction{}, err).Once()
				return fields{
					Ctrl: mockController,
				}
			},
			assertValue: func(t assert.TestingT, data any, i ...interface{}) bool {
				res := gin.H{"message": "error"}
				_ = json.Unmarshal(data.([]byte), &res)
				return assert.Equal(t, gin.H{"message": "error"}, res, i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := tt.makeFields()
			h := RequestHandler{
				Ctrl: f.Ctrl,
			}

			statusCode, body := helper.CreateTestServer(tt.makeRequest(), func(router *gin.Engine) {
				router.GET("/get-transaction-by-status", h.GetAllTransactionByRequestLimit)
			})
			assert.Equal(t, tt.expectedStatusCode, statusCode)
			if !tt.assertValue(t, body) {
				t.Errorf("assert value %v", body)
			}
		})
	}
}
