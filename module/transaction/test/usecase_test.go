package test

import (
	"golang/module/transaction"
	"golang/module/transaction/mocks/module/transaction"

	"reflect"
	"testing"
)

func TestUseCase_GetAllLimit(t *testing.T) {
	type fields struct {
		repo transaction.RepositoryInterface
	}
	type args struct {
		input transaction.FilterLimit
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []transaction.Transaction
		want1  error
		want2  int64
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Limit",
			fields: fields{
				repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				input: transaction.FilterLimit{
					Page:     1,
					PageSize: 100,
				},
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
			want2: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := transaction.UseCase{
				Repo: tt.fields.repo,
			}
			tt.fields.repo.(*mocks.RepositoryInterface).EXPECT().GetAllLimit(tt.args.input).Return(tt.want, tt.want1, tt.want2)
			got, got1, got2 := u.GetAllLimit(tt.args.input)
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

func TestUseCase_GetAllTransaction(t *testing.T) {
	type fields struct {
		repo transaction.RepositoryInterface
	}

	tests := []struct {
		name    string
		fields  fields
		want    []transaction.Transaction
		wantErr bool
	}{
		{
			name: "Test Case Get All Transaction",
			fields: fields{
				repo: mocks.NewRepositoryInterface(t),
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
			u := transaction.UseCase{
				Repo: tt.fields.repo,
			}
			tt.fields.repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransaction().Return(tt.want, nil)
			got, err := u.GetAllTransaction()
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

func TestUseCase_GetAllTransactionByDate(t *testing.T) {
	type fields struct {
		repo transaction.RepositoryInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Date",
			fields: fields{
				repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				start: "2021-08-01",
				end:   "2021-08-02",
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
			u := transaction.UseCase{
				Repo: tt.fields.repo,
			}
			tt.fields.repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByDate(tt.args.start, tt.args.end).Return(tt.want, nil)
			got, err := u.GetAllTransactionByDate(tt.args.start, tt.args.end)
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

func TestUseCase_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		repo transaction.RepositoryInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status",
			fields: fields{
				repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				status: "Success",
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
			u := transaction.UseCase{
				Repo: tt.fields.repo,
			}
			tt.fields.repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatus(tt.args.status).Return(tt.want, nil)
			got, err := u.GetAllTransactionByStatus(tt.args.status)
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

func TestUseCase_GetAllTransactionByStatusDate(t *testing.T) {
	type fields struct {
		repo transaction.RepositoryInterface
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
		want    []transaction.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Case Get All Transaction By Status Date",
			fields: fields{
				repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				status: "Success",
				start:  "2021-08-01",
				end:    "2021-08-02",
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
			u := transaction.UseCase{
				Repo: tt.fields.repo,
			}
			tt.fields.repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end).Return(tt.want, nil)
			got, err := u.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end)
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
