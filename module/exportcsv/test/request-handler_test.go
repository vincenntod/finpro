package test

import (
	"bytes"
	"encoding/json"
	"errors"
	mocks "golang/mocks/module/exportcsv"
	"golang/module/exportcsv"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestExportCSVRequestHandler_ExportCSVHandler(t *testing.T) {
	type fields struct {
		controller exportcsv.ExportCSVController
	}

	type test struct {
		name               string
		expectedStatusCode int
		makeRequest        func() *http.Request
		makeFields         func() fields
		assertValue        assert.ValueAssertionFunc
	}
	tests := []test{
		test{
			{"1"},
			{http.StatusFound},
			{
				
			},
		{},
		{},
		}
	}

	// mockController := mocks.NewExportCSVController(t)
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"", "", ""}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 	nil).Once()
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 	nil).Once()
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "", ""}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
	// 	nil).Once()
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
	// 	nil).Once()
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 	nil).Once()
	// mockController.EXPECT().ExportCSV(&exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"}).Return(
	// 	[][]string{{"id", "od_number", "status", "price", "created_at"},
	// 		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 	nil).Once()

	// 	type test struct {
	// 		name    string
	// 		request exportcsv.ExportCSVRequest
	// 		expect  [][]string
	// 	}

	// 	tests := []test{
	// 		test{
	// 			name:    "PositifTest-exportAll",
	// 			request: exportcsv.ExportCSVRequest{"", "", ""},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 		},
	// 		test{
	// 			name:    "PositifTest-rangeDate filter",
	// 			request: exportcsv.ExportCSVRequest{"", "2023-04-17", "2023-04-18"},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 		},
	// 		test{
	// 			name:    "PositifTest-SUCCESS status filter",
	// 			request: exportcsv.ExportCSVRequest{"SUCCESS", "", ""},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
	// 		},
	// 		test{
	// 			name:    "PositifTest-SUCCESS statusAndRageDate filter",
	// 			request: exportcsv.ExportCSVRequest{"SUCCESS", "2023-04-17", "2023-04-18"},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
	// 				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
	// 		},
	// 		test{
	// 			name:    "PositifTest-WAITING_FOR_DEBITTED filter",
	// 			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "", ""},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 		},
	// 		test{
	// 			name:    "Positif-WAITING_FOR_DEBITTTED statusAndRageDate filter",
	// 			request: exportcsv.ExportCSVRequest{"WAITING_FOR_DEBITTED", "2023-04-17", "2023-04-18"},
	// 			expect: [][]string{{"id", "od_number", "status", "price", "created_at"},
	// 				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
	// 				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
	// 		},
	// 	}
	// 	handler := exportcsv.NewRequestHandler(mockController)
	// 	for _, forTest := range tests {
	// 		t.Run(forTest.name, func(t *testing.T) {
	// 			result, err := handler.ExportCSVHandler()
	// 			assert.Nil(t, err)
	// 			assert.Equal(t, forTest.expect, result)

	// 		})
	// 	}

}
