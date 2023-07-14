package transaction

import (
	"golang/helper"
	"golang/module/transaction/dto"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestController_GetAllTransaction(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				page: 1,
				size: 10,
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-01 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransaction(tt.args.page, tt.args.size)
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

func TestController_GetAllTransactionByDate(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Date",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				start: "18-04-2023",
				end:   "19-04-2023",
				page:  1,
				size:  10,
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2023-04-18 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByDate(tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestController_GetAllTransactionByDateNoLimit(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Date",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				start: "18-04-2023",
				end:   "19-04-2023",
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2023-04-18 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByDateNoLimit(tt.args.start, tt.args.end)
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

func TestController_GetAllTransactionByRequest(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions` WHERE status =? AND(created_at BETWEEN ? AND ?) LIMIT 10"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2023-04-18 00:00:00"),
	)

	type args struct {
		req *dto.Request
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Request",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				req: &dto.Request{
					Status:    "Success",
					StartDate: "18-04-2023",
					EndDate:   "19-04-2023",
					Page:      1,
					Size:      10,
				},
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2023-04-18 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByRequest(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTransactionByRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTransactionByRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_GetAllTransactionByStatus(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				status: "Success",
				page:   1,
				size:   10,
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-01 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByStatus(tt.args.status, tt.args.page, tt.args.size)
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

func TestController_GetAllTransactionByStatusDate(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status Date",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				status: "Success",
				start:  "18-04-2023",
				end:    "19-04-2023",
				page:   1,
				size:   10,
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2023-04-18 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestController_GetAllTransactionByStatusDateNoLimit(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status Date",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				status: "Success",
				start:  "18-04-2023",
				end:    "19-04-2023",
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2023-04-18 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByStatusDateNoLimit(tt.args.status, tt.args.start, tt.args.end)
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

func TestController_GetAllTransactionByStatusNoLimit(t *testing.T) {

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
		c       Controller
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction By Status",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			args: args{
				status: "Success",
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-01 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionByStatusNoLimit(tt.args.status)
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

func TestController_GetAllTransactionNoLimit(t *testing.T) {

	mockQuery, mockDB := helper.NewMockQueryDb(t)
	query := "SELECT * FROM `transactions`"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "oda_number", "bill_amount", "status", "created_at"}).AddRow(1, "12345678", 100000, "Success", "2021-08-01 00:00:00"),
	)

	tests := []struct {
		name    string
		c       Controller
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction No Limit",
			c: Controller{
				UseCase{
					Repo: Repository{
						DB: mockDB,
					},
				},
			},
			want: &dto.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data: []dto.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-01 00:00:00",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				UseCase: tt.c.UseCase,
			}
			got, err := c.GetAllTransactionNoLimit()
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
