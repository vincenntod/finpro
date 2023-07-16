package exporttransaction

import (
	"golang/module/exporttransaction/dto"
	"golang/module/exporttransaction/entity"
	mocks "golang/module/exporttransaction/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionStringConverter(t *testing.T) {
	request1 := []entity.Transaction{{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"}}
	request2 := []entity.Transaction{{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"}}
	request3 := []entity.Transaction{{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"}}
	request4 := []entity.Transaction{{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"}}
	expected1 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"}}
	expected2 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}}
	expected3 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}
	expected4 := [][]string{{"id", "oda_number", "status", "price", "created_at"}, {"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}}
	tests := []struct {
		name     string
		request  []entity.Transaction
		expected [][]string
	}{
		{"positif1", request1, expected1},
		{"positif2", request2, expected2},
		{"positif3", request3, expected3},
		{"positif4", request4, expected4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := TransactionStringConverter(test.request)
			assert.Equal(t, test.expected, result)
			assert.Equal(t, nil, err)

		})
	}

}

func TestExportCSV(t *testing.T) {

	mockRepo := mocks.NewRepository(t)
	mockRepo.EXPECT().GetAllTransaction().Return([]entity.Transaction{
		{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetAllTransactionByRangeDateFilter(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}}).Return([]entity.Transaction{
		{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusFilter(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}}).Return([]entity.Transaction{
		{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusAndRangeDateFilter(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}}).Return([]entity.Transaction{
		{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusFilter(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "", EndDate: ""}}).Return([]entity.Transaction{
		{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()
	mockRepo.EXPECT().GetTransactionByStatusAndRangeDateFilter(&dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "2023-04-17", EndDate: "2023-04-18"}}).Return([]entity.Transaction{
		{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
		{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
	}, nil).Once()

	type test struct {
		name    string
		request dto.ExportCSVRequest
		expect  [][]string
	}

	tests := []test{
		{
			name:    "PositifTest-exportAll",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "", EndDate: ""}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, {
			name:    "PositifTest-rangeDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, {
			name:    "PositifTest-SUCCESS status filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		}, {
			name:    "PositifTest-SUCCESS statusAndRageDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"2241", "9860709914616572", "SUCCESS", "1000.000000", "2023-04-18"},
				{"2242", "0969598565950745", "SUCCESS", "1000.000000", "2023-04-18"}},
		}, {
			name:    "PositifTest-WAITING_FOR_DEBITTED filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "", EndDate: ""}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		}, {
			name:    "Positif-WAITING_FOR_DEBITTTED statusAndRageDate filter",
			request: dto.ExportCSVRequest{Filter: dto.Filter{Status: "WAITING_FOR_DEBITTED", StartDate: "2023-04-17", EndDate: "2023-04-18"}},
			expect: [][]string{{"id", "oda_number", "status", "price", "created_at"},
				{"241", "5655429080696649", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"},
				{"242", "1262435180570130", "WAITING_FOR_DEBITTED", "1000.000000", "2023-04-18"}},
		},
	}

	usecase := NewUseCase(mockRepo)
	for _, forTest := range tests {
		t.Run(forTest.name, func(t *testing.T) {
			result, err := usecase.ExportCSV(&forTest.request)
			assert.Nil(t, err)
			assert.Equal(t, forTest.expect, result)

		})
	}

}
