package test

import (
	"golang/module/exportcsv"
	"golang/module/exportcsv/test/helper"
	mocks "golang/module/exportcsv/test/mock"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandler_ExportCSVHandler(t *testing.T) {
	type fields struct {
		controller exportcsv.Controller
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
					ExportCSV(&exportcsv.ExportCSVRequest{"", "", ""}, &exportcsv.UrlRequest{"/export-transaction/csv", "/export-transaction/csv"}).
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
					ExportCSV(&exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"}, &exportcsv.UrlRequest{"/export-transaction/csv?start_date=2023-04-17&end_date=2023-04-18", "/export-transaction/csv"}).
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
					ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "", ""}, &exportcsv.UrlRequest{"/export-transaction/csv?status=SUCCESS", "/export-transaction/csv"}).
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
					ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"}, &exportcsv.UrlRequest{"/export-transaction/csv?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", "/export-transaction/csv"}).
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
					ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""}, &exportcsv.UrlRequest{"/export-transaction/csv?status=WAITING_FOR_DEBITTED", "/export-transaction/csv"}).
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
					ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"}, &exportcsv.UrlRequest{"/export-transaction/csv?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", "/export-transaction/csv"}).
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
			handler := exportcsv.NewRequestHandler(ctrlMock().controller)

			statusCode, body := helper.CreateTestServer(forTest.makeRequest(), func(router *gin.Engine) {
				router.GET("/export-transaction/:type", handler.ExportCSVHandler)
			})
			assert.Equal(t, 200, statusCode)
			assert.NotEmpty(t, body)
		})
	}
}
