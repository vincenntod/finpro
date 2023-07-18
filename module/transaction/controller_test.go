package transaction

import (
	"golang/module/transaction/dto"
	"golang/module/transaction/entities"
	"golang/module/transaction/mocks"
	"reflect"
	"testing"
)

func TestController_GetAllTransaction(t *testing.T) {

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransaction(1, 1).Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
	}
	type args struct {
		page int
		size int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				page: 1,
				size: 1,
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByDate("2021-08-01 00:00:00", "2021-08-01 00:00:00", 1, 1).Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
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
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				start: "2021-08-01 00:00:00",
				end:   "2021-08-01 00:00:00",
				page:  1,
				size:  1,
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByDateNoLimit("2021-08-01 00:00:00", "2021-08-01 00:00:00").Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
	}
	type args struct {
		start string
		end   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				start: "2021-08-01 00:00:00",
				end:   "2021-08-01 00:00:00",
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
				UseCase: tt.fields.UseCase,
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

func TestController_GetAllTransactionByStatus(t *testing.T) {

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByStatus("Success", 1, 1).Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
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
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				status: "Success",
				page:   1,
				size:   1,
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByStatusDate("Success", "2021-08-01 00:00:00", "2021-08-01 00:00:00", 1, 1).Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
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
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				status: "Success",
				start:  "2021-08-01 00:00:00",
				end:    "2021-08-01 00:00:00",
				page:   1,
				size:   1,
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByStatusDateNoLimit("Success", "2021-08-01 00:00:00", "2021-08-01 00:00:00").Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
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
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
			},
			args: args{
				status: "Success",
				start:  "2021-08-01 00:00:00",
				end:    "2021-08-01 00:00:00",
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionByStatusNoLimit("Success").Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
	}
	type args struct {
		status string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
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
				UseCase: tt.fields.UseCase,
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

	mockUsecase := mocks.NewUseCaseInterface(t)
	mockUsecase.EXPECT().GetAllTransactionNoLimit().Return([]entities.Transaction{
		{
			Id:         1,
			OdaNumber:  "12345678",
			BillAmount: 100000,
			Status:     "Success",
			CreatedAt:  "2021-08-01 00:00:00",
		},
	}, nil).Once()

	type fields struct {
		UseCase UseCaseInterface
	}
	tests := []struct {
		name    string
		fields  fields
		want    *dto.GetAllResponseDataTransaction
		wantErr bool
	}{
		{
			name: "Test Case 1 | Success",
			fields: fields{
				UseCase: mockUsecase,
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
				UseCase: tt.fields.UseCase,
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
