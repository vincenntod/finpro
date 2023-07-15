package transaction

import (
	"golang/module/transaction/entities"
	"golang/module/transaction/mocks"
	"reflect"
	"testing"
)

func TestUseCase_GetAllTransaction(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransaction(tt.args.page, tt.args.size).Return(tt.want, nil).Once()
			got, err := u.GetAllTransaction(tt.args.page, tt.args.size)
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
		Repo RepositoryInterface
	}
	type args struct {
		start string
		end   string
		page  int
		size  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				start: "2021-08-01 00:00:00",
				end:   "2021-08-01 00:00:00",
				page:  1,
				size:  10,
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByDate(tt.args.start, tt.args.end, tt.args.page, tt.args.size).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByDate(tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestUseCase_GetAllTransactionByDateNoLimit(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				start: "2021-08-01 00:00:00",
				end:   "2021-08-01 00:00:00",
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByDateNoLimit(tt.args.start, tt.args.end).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByDateNoLimit(tt.args.start, tt.args.end)
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

func TestUseCase_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}
	type args struct {
		status string
		page   int
		size   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatus(tt.args.status, tt.args.page, tt.args.size).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByStatus(tt.args.status, tt.args.page, tt.args.size)
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
		Repo RepositoryInterface
	}
	type args struct {
		status string
		start  string
		end    string
		page   int
		size   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				status: "Success",
				start:  "2021-08-01 00:00:00",
				end:    "2021-08-01 00:00:00",
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end, tt.args.page, tt.args.size).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end, tt.args.page, tt.args.size)
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

func TestUseCase_GetAllTransactionByStatusDateNoLimit(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
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
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
			},
			args: args{
				status: "Success",
				start:  "2021-08-01 00:00:00",
				end:    "2021-08-01 00:00:00",
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatusDateNoLimit(tt.args.status, tt.args.start, tt.args.end).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByStatusDateNoLimit(tt.args.status, tt.args.start, tt.args.end)
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

func TestUseCase_GetAllTransactionByStatusNoLimit(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionByStatusNoLimit(tt.args.status).Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionByStatusNoLimit(tt.args.status)
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

func TestUseCase_GetAllTransactionNoLimit(t *testing.T) {
	type fields struct {
		Repo RepositoryInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entities.Transaction
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				Repo: mocks.NewRepositoryInterface(t),
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
			u := UseCase{
				Repo: tt.fields.Repo,
			}
			tt.fields.Repo.(*mocks.RepositoryInterface).EXPECT().GetAllTransactionNoLimit().Return(tt.want, nil).Once()
			got, err := u.GetAllTransactionNoLimit()
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
