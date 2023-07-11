package test

import (
	"golang/helper"
	"golang/module/transaction"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_GetAllLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	queryGet := "SELECT count(*) FROM `transactions`"
	queryLimit := "SELECT * FROM `transactions` ORDER BY created_at desc LIMIT 100"
	mockQuery.ExpectQuery(queryGet).WillReturnRows(
		sqlmock.NewRows([]string{"total_data"}).AddRow(10),
	)
	mockQuery.ExpectQuery(queryLimit).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	tests := []struct {
		name  string
		r     transaction.Repository
		want  []transaction.Transaction
		want1 error
		want2 int64
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Limit",
			r: transaction.Repository{
				DB: mockDB,
			},
			want: []transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			},
			want1: nil,
			want2: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.r.GetAllLimit(transaction.FilterLimit{1, 100})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllLimit() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetAllLimit() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("GetAllLimit() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestRepository_GetAllTransaction(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions`"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	tests := []struct {
		name    string
		r       transaction.Repository
		want    []transaction.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction",
			r: transaction.Repository{
				DB: mockDB,
			},
			want: []transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransaction()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransaction() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionByDate(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE created_at BETWEEN ? AND ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	tests := []struct {
		name    string
		r       transaction.Repository
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Date",
			r: transaction.Repository{
				DB: mockDB,
			},
			want: []transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2023-04-18 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllTransactionByDate("18-04-2023", "19-04-2023")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionByStatus(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status = ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	tests := []struct {
		name    string
		r       transaction.Repository
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status",
			r: transaction.Repository{
				DB: mockDB,
			},
			want: []transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.r.GetAllTransactionByStatus("Success")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByStatus() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionByStatusDate(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status =? AND(created_at BETWEEN ? AND ?)"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	tests := []struct {
		name    string
		r       transaction.Repository
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status Date",
			r: transaction.Repository{
				DB: mockDB,
			},
			want: []transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2023-04-18 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.r.GetAllTransactionByStatusDate("Success", "18-04-2023", "19-04-2023")
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByStatusDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByStatusDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
