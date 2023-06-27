package transaction

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestController_GetAllTransaction(t *testing.T) {
	type fields struct {
		useCase UseCaseInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get All Transaction | Success",
			fields: fields{
				useCase: NewMockUseCaseInterface(t),
			},
			want: &GetAllResponseDataTransaction{
				Code:    200,
				Message: "Success",
				Error:   "Not Found",
				Data: []TransactionItemResponse{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*MockUseCaseInterface).EXPECT().GetAllTransaction().Return([]Transaction{
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
		useCase UseCaseInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get All By Date| Success",
			fields: fields{
				useCase: NewMockUseCaseInterface(t),
			},
			args: args{
				start: "2023-01-01",
				end:   "2023-07-01",
			},
			want: &GetAllResponseDataTransaction{
				Code:    200,
				Message: "Success",
				Error:   " ",
				Data: []TransactionItemResponse{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*MockUseCaseInterface).EXPECT().GetAllTransactionByDate(tt.args.start, tt.args.end).Return([]Transaction{
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
		useCase UseCaseInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get All By Status| Success",
			fields: fields{
				useCase: NewMockUseCaseInterface(t),
			},
			args: args{
				status: "SUCCESS",
			},
			want: &GetAllResponseDataTransaction{
				Code:    http.StatusOK,
				Message: "Success",
				Error:   "Not Found",
				Data: []TransactionItemResponse{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*MockUseCaseInterface).EXPECT().GetAllTransactionByStatus(tt.args.status).Return([]Transaction{
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
		useCase UseCaseInterface
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
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get All By Status and Date| Success",
			fields: fields{
				useCase: NewMockUseCaseInterface(t),
			},
			args: args{
				status: "SUCCESS",
				start:  "2023-01-01",
				end:    "2023-07-01",
			},
			want: &GetAllResponseDataTransaction{
				Code:    http.StatusOK,
				Message: "Success",
				Error:   " ",
				Data: []TransactionItemResponse{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			tt.fields.useCase.(*MockUseCaseInterface).EXPECT().GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end).Return([]Transaction{
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

func TestController_GetTransactionByDate(t *testing.T) {
	type fields struct {
		useCase UseCaseInterface
	}
	type args struct {
		req   FilterByDate
		input FilterLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			got, err := c.GetTransactionByDate(tt.args.req, tt.args.input)
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

func TestController_GetTransactionByStatusAndDate(t *testing.T) {
	type fields struct {
		useCase UseCaseInterface
	}
	type args struct {
		req   FilterByStatusDate
		input FilterLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetAllResponseDataTransaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Controller{
				useCase: tt.fields.useCase,
			}
			got, err := c.GetTransactionByStatusAndDate(tt.args.req, tt.args.input)
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

func TestNewController(t *testing.T) {
	type args struct {
		useCase UseCaseInterface
	}
	tests := []struct {
		name string
		args args
		want ControllerInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewController(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewController() = %v, want %v", got, tt.want)
			}
		})
	}
}
