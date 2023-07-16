package exporttransaction

import (
	"golang/module/exporttransaction/dto"
	mocks "golang/module/exporttransaction/mock"
	"golang/module/exporttransaction/mock/helper"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_ExportTransactionHandler(t *testing.T) {
	type fields struct {
		controller Controller
	}

	tests := []struct {
		name        string
		makeRequest func() *http.Request
		makeFields  func() fields
	}{
		{
			name: "positif 1 exportAllTransaction",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
						{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
						{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
						{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
		{
			name: "positif 2 exportAllTransactionByRangeDateFilter",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv?start_date=2023-04-17&end_date=2023-04-18", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
						{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
						{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
						{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
		{
			name: "positif 3 exportTransactionBy-SUCCESS-StatusFilter",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv?status=SUCCESS", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
						{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
		{
			name: "positif 4 exportTransactionBy-SUCCESS-StatusAndRaneDateFilter",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
						{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
		{
			name: "positif 5 exportTransactionBy-WAITING_FOR_DEBITTED-StatusFilter",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv?status=WAITING_FOR_DEBITTED", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
						{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
		{
			name: "positif 6 exportTransactionBy-WAITING_FOR_DEBITTED-StatusFilter",
			makeRequest: func() *http.Request {
				request, _ := http.NewRequest(http.MethodGet, "/export-transaction/csv?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", nil)
				return request
			},
			makeFields: func() fields {
				mockController := mocks.NewController(t)
				mockController.EXPECT().
					ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).
					Return([][]string{{"id", "od_number", "status", "price", "created_at"},
						{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
						{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
						nil).Once()
				return fields{controller: mockController}
			},
		},
	}
	for _, forTest := range tests {
		t.Run(forTest.name, func(t *testing.T) {
			ctrlMock := forTest.makeFields
			handler := NewRequestHandler(ctrlMock().controller)

			statusCode, body := helper.CreateTestServer(forTest.makeRequest(), func(router *gin.Engine) {
				router.GET("/export-transaction/csv", handler.ExportTransactionHandler)
			})
			assert.Equal(t, 200, statusCode)
			assert.NotEmpty(t, body)
		})
	}
}
