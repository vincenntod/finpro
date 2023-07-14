package transaction

import (
	"golang/helper"
	"golang/module/transaction/entities"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_GetAllTransaction(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` LIMIT 10"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	type args struct {
		page int
		size int
	}

	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				page: 1,
				size: 10,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransaction(tt.args.page, tt.args.size)
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
	query := "SELECT * FROM `transactions` WHERE created_at BETWEEN ? AND ? LIMIT 10"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	type args struct {
		start string
		end   string
		page  int
		size  int
	}

	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Date",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				start: "2023-04-18",
				end:   "2023-04-19",
				page:  1,
				size:  10,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByDate(tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestRepository_GetAllTransactionByDateNoLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE created_at BETWEEN ? AND ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		args    args
		r       Repository
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Date",
			args: args{
				start: "2023-04-18",
				end:   "2023-04-19",
			},
			r: Repository{
				DB: mockDB,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByDateNoLimit(tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByDateNoLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByDateNoLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionByStatus(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status = ? LIMIT 10"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	type args struct {
		status string
		page   int
		size   int
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				status: "Success",
				page:   1,
				size:   10,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByStatus(tt.args.status, tt.args.page, tt.args.size)
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
	query := "SELECT * FROM `transactions` WHERE status =? AND(created_at BETWEEN ? AND ?) LIMIT 10"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	type args struct {
		status string
		start  string
		end    string
		page   int
		size   int
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status Date",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				status: "Success",
				start:  "2023-04-18",
				end:    "2023-04-19",
				page:   1,
				size:   10,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestRepository_GetAllTransactionByStatusDateNoLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status =? AND(created_at BETWEEN ? AND ?)"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	type args struct {
		status string
		start  string
		end    string
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status Date",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				status: "Success",
				start:  "2023-04-18",
				end:    "2023-04-19",
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByStatusDateNoLimit(tt.args.status, tt.args.start, tt.args.end)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByStatusDateNoLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByStatusDateNoLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionByStatusNoLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status = ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	type args struct {
		status string
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status",
			r: Repository{
				DB: mockDB,
			},
			args: args{
				status: "Success",
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionByStatusNoLimit(tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByStatusNoLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByStatusNoLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllTransactionNoLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions`"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	tests := []struct {
		name    string
		r       Repository
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction",
			r: Repository{
				DB: mockDB,
			},
			want: []entities.Transaction{
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
			got, err := tt.r.GetAllTransactionNoLimit()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionNoLimit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionNoLimit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
