package exporttransaction

import (
	"golang/module/exporttransaction/dto"
	"golang/module/exporttransaction/entity"
	"golang/module/exporttransaction/mock/helper"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repository_GetAllTransaction(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")
	mockQuery.ExpectQuery("SELECT * FROM `transaction`").WillReturnRows(rows)

	tests := []struct {
		name    string
		r       Repository
		want    []entity.Transaction
		wantErr bool
	}{
		{name: "succes",
			r: NewRepository(mockDb),
			want: []entity.Transaction{
				{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransaction()
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetTransactionByStatusFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE status = ?").WillReturnRows(rows)

	type args struct {
		req *dto.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entity.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    NewRepository(mockDb),
			args: args{req: &dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "", EndDate: ""}}},
			want: []entity.Transaction{
				{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTransactionByStatusFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetTransactionByStatusFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetTransactionByStatusFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetAllTransactionByRangeDateFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(241, "5655429080696649", "9018506264648524", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(242, "1262435180570130", "1798502071284385", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "WAITING_FOR_DEBITTED", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE created_at BETWEEN ? AND ?").WillReturnRows(rows)

	type args struct {
		req *dto.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entity.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    NewRepository(mockDb),
			args: args{req: &dto.ExportCSVRequest{Filter: dto.Filter{Status: "", StartDate: "2023-04-17", EndDate: "2023-04-18"}}},
			want: []entity.Transaction{
				{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 241, Oda_number: "5655429080696649", Bank_account_no: "9018506264648524", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 242, Oda_number: "1262435180570130", Bank_account_no: "1798502071284385", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "WAITING_FOR_DEBITTED", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransactionByRangeDateFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetAllTransactionByRangeDateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetAllTransactionByRangeDateFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetTransactionByStatusAndRangeDateFilter(t *testing.T) {
	mockQuery, mockDb := helper.NewMockQueryDb(t)
	rows := sqlmock.NewRows([]string{"id", "oda_number", "bank_account_no", "billing_cycle_date", "payment_due_date",
		"overflow_amount", "bill_amount", "principal_amount", "interest_amount", "total_fee_amount", "status", "payment_method",
		"auto_debet_counter", "created_at", "updated_at", "payment_amount",
		"billing_gen_date"}).AddRow(2241, "9860709914616572", "2137645127682149", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17").
		AddRow(2242, "0969598565950745", "8298344194001672", "2023-04-17 00:00:00", "2023-04-17", 1000.000000, 1000.000000, 1000.000000, 1000.000000, 1000.000000, "SUCCESS", "", 0, "2023-04-18", "", 0.000000, "2023-04-17")

	mockQuery.ExpectQuery("SELECT * FROM `transaction` WHERE status =? AND(created_at BETWEEN ? AND ?)").WillReturnRows(rows)

	type args struct {
		req *dto.ExportCSVRequest
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entity.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "succes",
			r:    NewRepository(mockDb),
			args: args{req: &dto.ExportCSVRequest{Filter: dto.Filter{Status: "SUCCESS", StartDate: "2023-04-17", EndDate: "2023-04-18"}}},
			want: []entity.Transaction{
				{Id: 2241, Oda_number: "9860709914616572", Bank_account_no: "2137645127682149", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
				{Id: 2242, Oda_number: "0969598565950745", Bank_account_no: "8298344194001672", Billing_cycle_date: "2023-04-17 00:00:00", Payment_due_date: "2023-04-17", Overflow_amount: 1000.000000, Bill_amount: 1000.000000, Principal_amount: 1000.000000, Interest_amount: 1000.000000, Total_fee_amount: 1000.000000, Status: "SUCCESS", Payment_method: "", Auto_debet_counter: 0, Created_at: "2023-04-18", Updated_at: "", Payment_amount: 0.000000, Billing_gen_date: "2023-04-17"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetTransactionByStatusAndRangeDateFilter(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetTransactionByStatusAndRangeDateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetTransactionByStatusAndRangeDateFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
