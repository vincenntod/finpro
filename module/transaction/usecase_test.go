package transaction

import (
	"reflect"
	"testing"
	"time"
)

func TestNewUseCase(t *testing.T) {
	type args struct {
		repo RepositoryInterface
	}
	tests := []struct {
		name string
		args args
		want UseCaseInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUseCase(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetAllTransaction(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}

	tests := []struct {
		name    string
		fields  fields
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test GetAll",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			want: []Transaction{
				{
					Id:               1,
					OdaNumber:        123456789,
					BankAccountNo:    123456789,
					BillingCycleDate: "2021-01-01",
					PaymentDueDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					OverflowAmount:   100000,
					BillAmount:       100000,
					PrincipalAmount:  100000,
					InterestAmount:   100000,
					TotalFeeAmount:   100000,
					Status:           "Success",
					PaymentMethod:    "BCA",
					AutoDebetCounter: 1,
					CreatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					UpdatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsHold:           false,
					IsFstlPending:    false,
					IsHstlPending:    false,
					IsLaaPositif:     false,
					PaymentAmount:    100000,
					BillingGenDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsOdaPositif:     false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().GetAllTransaction().Return(tt.want, nil)
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
		repo RepositoryInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetAll By Date",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				start: "2021-01-01",
				end:   "2021-01-01",
			},
			want: []Transaction{
				{
					Id:               1,
					OdaNumber:        123456789,
					BankAccountNo:    123456789,
					BillingCycleDate: "2021-01-01",
					PaymentDueDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					OverflowAmount:   100000,
					BillAmount:       100000,
					PrincipalAmount:  100000,
					InterestAmount:   100000,
					TotalFeeAmount:   100000,
					Status:           "Success",
					PaymentMethod:    "BCA",
					AutoDebetCounter: 1,
					CreatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					UpdatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsHold:           false,
					IsFstlPending:    false,
					IsHstlPending:    false,
					IsLaaPositif:     false,
					PaymentAmount:    100000,
					BillingGenDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsOdaPositif:     false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().GetAllTransactionByDate(tt.args.start, tt.args.end).Return(tt.want, nil)
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
		repo RepositoryInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GetAll By Status",
			fields: fields{
				repo: NewMockRepositoryInterface(t),
			},
			args: args{
				status: "Success",
			},
			want: []Transaction{
				{
					Id:               1,
					OdaNumber:        123456789,
					BankAccountNo:    123456789,
					BillingCycleDate: "2021-01-01",
					PaymentDueDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					OverflowAmount:   100000,
					BillAmount:       100000,
					PrincipalAmount:  100000,
					InterestAmount:   100000,
					TotalFeeAmount:   100000,
					Status:           "Success",
					PaymentMethod:    "BCA",
					AutoDebetCounter: 1,
					CreatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					UpdatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsHold:           false,
					IsFstlPending:    false,
					IsHstlPending:    false,
					IsLaaPositif:     false,
					PaymentAmount:    100000,
					BillingGenDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsOdaPositif:     false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().GetAllTransactionByStatus(tt.args.status).Return(tt.want, nil)
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
		repo RepositoryInterface
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
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "GetAll By Status and Date",
			fields:  fields{
				repo: NewMockRepositoryInterface(t),
			},
			args:    args{
				status: "Success",
				start:  "2021-01-01",
				end:    "2021-01-01",
			},
			want:    []Transaction{
				{
					Id:               1,
					OdaNumber:        123456789,
					BankAccountNo:    123456789,
					BillingCycleDate: "2021-01-01",
					PaymentDueDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					OverflowAmount:   100000,
					BillAmount:       100000,
					PrincipalAmount:  100000,
					InterestAmount:   100000,
					TotalFeeAmount:   100000,
					Status:           "Success",
					PaymentMethod:    "BCA",
					AutoDebetCounter: 1,
					CreatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					UpdatedAt:        time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsHold:           false,
					IsFstlPending:    false,
					IsHstlPending:    false,
					IsLaaPositif:     false,
					PaymentAmount:    100000,
					BillingGenDate:   time.Date(2023, 06, 01, 11, 59, 10, 11, time.UTC),
					IsOdaPositif:     false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			tt.fields.repo.(*MockRepositoryInterface).EXPECT().GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end).Return(tt.want, nil)
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

func TestUseCase_GetTransactionByDate(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		req   FilterByDate
		input FilterLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetTransactionByDate(tt.args.req, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionByDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransactionByDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetTransactionByStatusAndDate(t *testing.T) {
	type fields struct {
		repo RepositoryInterface
	}
	type args struct {
		req   FilterByStatusDate
		input FilterLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCase{
				repo: tt.fields.repo,
			}
			got, err := u.GetTransactionByStatusAndDate(tt.args.req, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionByStatusAndDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransactionByStatusAndDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
