package exporttransaction

import (
	"golang/module/exporttransaction/dto"
	mocks "golang/module/exporttransaction/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportCSVController(t *testing.T) {
	mockUsecase := mocks.NewUsecase(t)
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv", PathUrl: "/export-transaction/csv"}}).Return([][]string{{"id", "oda_number", "status", "price", "created_at"},
		{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
		{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
		{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
		{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}, nil).Once()
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS", PathUrl: "/export-transaction/csv"}}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
			{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED", PathUrl: "/export-transaction/csv"}}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()
	mockUsecase.EXPECT().ExportCSV(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}}).Return(
		[][]string{{"id", "oda_number", "status", "price", "created_at"},
			{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
			{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		nil).Once()

	type test struct {
		name    string
		request dto.ExportCSVRequest
		expect  [][]string
	}

	tests := []test{
		{
			name:    "PositifTest-exportAll",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		{
			name:    "PositifTest-rangeDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		{
			name:    "PositifTest-SUCCESS status filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		},
		{
			name:    "PositifTest-SUCCESS statusAndRageDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=SUCCESS&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		},
		{
			name:    "PositifTest-WAITING_FOR_DEBITTED filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "", EndDate: ""}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
		{
			name:    "Positif-WAITING_FOR_DEBITTTED statusAndRageDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "2023-04-17", EndDate: "2023-04-18"}, Url: dto.Url{ClientUrl: "/export-transaction/csv?status=WAITING_FOR_DEBITTED&start_date=2023-04-17&end_date=2023-04-18", PathUrl: "/export-transaction/csv"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
	}
	controller := NewController(mockUsecase)
	for _, forTest := range tests {
		t.Run(forTest.name, func(t *testing.T) {
			result, err := controller.ExportCSV(&forTest.request)
			assert.Nil(t, err)
			assert.Equal(t, forTest.expect, result)

		})
	}

}
