package test

import (
	"golang/module/transaction"
	"golang/module/transaction/mocks/module/transaction"

	"reflect"
	"testing"
)

func TestController_GetAllLimit(t *testing.T) {
	type fields struct {
		useCase transaction.UseCaseInterface
	}
	type args struct {
		input transaction.FilterLimit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *transaction.GetAllResponseDataTransaction
		want1  error
		want2  int64
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Limit",
			fields: fields{
				useCase: mocks.NewUseCaseInterface(t),
			},
			args: args{
				input: transaction.FilterLimit{
					Page:     1,
					PageSize: 100,
				},
			},
			want: &transaction.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data:      []transaction.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-01 00:00:00",
					},
			},
		},
			want1: nil,
			want2: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := transaction.Controller{
				UseCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*mocks.UseCaseInterface).EXPECT().GetAllLimit(tt.args.input).Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			}, nil, 1)
			got, got1, got2 := c.GetAllLimit(tt.args.input)
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

func TestController_GetAllTransaction(t *testing.T) {
	type fields struct {
		useCase transaction.UseCaseInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    *transaction.GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction",
			fields: fields{
				useCase: mocks.NewUseCaseInterface(t),
			},
			want: &transaction.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data:      []transaction.TransactionItemResponse{
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
			c := transaction.Controller{
				UseCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*mocks.UseCaseInterface).EXPECT().GetAllTransaction().Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			}, nil)
			got, err := c.GetAllTransaction()
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
	type fields struct {
		useCase transaction.UseCaseInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *transaction.GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Date",
			fields: fields{
				useCase: mocks.NewUseCaseInterface(t),
			},
			args: args{
				start: "2021-08-01 00:00:00",
				end:   "2021-08-03 00:00:00",
			},
			want: &transaction.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data:      []transaction.TransactionItemResponse{
					{
						Id:         1,
						OdaNumber:  "12345678",
						BillAmount: 100000,
						Status:     "Success",
						CreatedAt:  "2021-08-02 00:00:00",
					},
			},
		},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := transaction.Controller{
				UseCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*mocks.UseCaseInterface).EXPECT().GetAllTransactionByDate(tt.args.start, tt.args.end).Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-02 00:00:00",
				},
			}, nil)
			got, err := c.GetAllTransactionByDate(tt.args.start, tt.args.end)
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

func TestController_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		useCase transaction.UseCaseInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *transaction.GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status",
			fields: fields{
				useCase: mocks.NewUseCaseInterface(t),
			},
			args: args{
				status: "Success",
			},
			want: &transaction.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data:      []transaction.TransactionItemResponse{
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
			c := transaction.Controller{
				UseCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*mocks.UseCaseInterface).EXPECT().GetAllTransactionByStatus(tt.args.status).Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			}, nil)
			got, err := c.GetAllTransactionByStatus(tt.args.status)
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
	type fields struct {
		useCase transaction.UseCaseInterface
	}
	type args struct {
		status string
		start  string
		end    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *transaction.GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status Date",
			fields: fields{
				useCase: mocks.NewUseCaseInterface(t),
			},
			args: args{
				status: "Success",
				start:  "2021-08-01 00:00:00",
				end:    "2021-08-03 00:00:00",
			},
			want: &transaction.GetAllResponseDataTransaction{
				Code:      200,
				Message:   "Data Berhasil Diambil",
				Error:     "Success",
				TotalData: 1,
				Data:      []transaction.TransactionItemResponse{
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
			c := transaction.Controller{
				UseCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*mocks.UseCaseInterface).EXPECT().GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end).Return([]transaction.Transaction{
				{
					Id:         1,
					OdaNumber:  "12345678",
					BillAmount: 100000,
					Status:     "Success",
					CreatedAt:  "2021-08-01 00:00:00",
				},
			}, nil)
			got, err := c.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end)
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


