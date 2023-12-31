package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
	"golang/module/account/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_GetDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().GetDataUser().Return([]entities.Account{{
		Name:  "Vincen",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		want    []entities.Account
		wantErr bool
	}{
		{
			name: "Success Get Data User",
			u: UseCase{
				Repo: mockery,
			},
			want: []entities.Account{{
				Name:  "Vincen",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetDataUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.GetDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetDataUserById(t *testing.T) {
	type args struct {
		id string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().GetDataUserById("1").Return(entities.Account{
		Name:  "Vincen",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "Success Get Data User By Id",
			u:    UseCase{Repo: mockery},
			args: args{id: "1"},
			want: entities.Account{
				Name:  "Vincen",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetDataUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetDataUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.GetDataUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_EditDataUser(t *testing.T) {
	type args struct {
		id  string
		req *dto.EditDataUserRequest
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().EditDataUser("1", gomock.Any()).Return(entities.Account{
		Name:  "Vincencius Maxwell",
		Email: "vincen@gmail.com",
		Role:  "admin",
		Phone: "123",
	}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "Success Edit Data User",
			u:    UseCase{Repo: mockery},
			args: args{
				id: "1",
				req: &dto.EditDataUserRequest{
					Name:  "Vincencius Maxwell",
					Role:  "admin",
					Phone: "123",
				},
			},
			want: entities.Account{
				Name:  "Vincencius Maxwell",
				Email: "vincen@gmail.com",
				Role:  "admin",
				Phone: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.EditDataUser(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.EditDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.EditDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CreateAccount(t *testing.T) {
	type args struct {
		req *dto.CreateRequest
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().CreateAccount(gomock.Any()).Return(entities.Account{
		Id:       1,
		Name:     "Vincencius Maxwell",
		Email:    "vincen@gmail.com",
		Password: "123456",
		Role:     "admin",
		Phone:    "123",
	}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "Success Create Account",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				req: &dto.CreateRequest{
					Name:     "Vincencius Maxwell",
					Email:    "vincen@gmail.com",
					Password: "123456",
					Role:     "admin",
					Phone:    "123",
				},
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincencius Maxwell",
				Email:    "vincen@gmail.com",
				Password: "123456",
				Role:     "admin",
				Phone:    "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.CreateAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_DeleteDataUser(t *testing.T) {
	type args struct {
		id string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().DeleteDataUser("1").Return(entities.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "Success Delete Data User",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				id: "1",
			},
			want:    entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.DeleteDataUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.DeleteDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.DeleteDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_Login(t *testing.T) {
	type args struct {
		req *dto.LoginResponseRequest
	}

	tests := []struct {
		name    string
		u       UseCase
		args    args
		want1   entities.Account
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := tt.u.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UseCase.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUseCase_SendEmailForgotPassword(t *testing.T) {
	type args struct {
		email string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().SendEmail("maxwelvincen@gmail.com").Return(entities.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want:    entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.SendEmailForgotPassword(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendEmailForgotPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.SendEmailForgotPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_SendEmailRegistration(t *testing.T) {
	type args struct {
		email string
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().SendEmail("maxwelvincen@gmail.com").Return(entities.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want:    entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.SendEmailRegistration(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendEmailRegistration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.SendEmailRegistration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CompareVerificationCode(t *testing.T) {
	type args struct {
		verificationCode *dto.VerificationCodeRequest
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().CompareVerificationCode(&dto.VerificationCodeRequest{
		Email: "maxwelvincen@gmail.com",
		Code:  "1234",
	}).Return(entities.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "Success",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				&dto.VerificationCodeRequest{
					Email: "maxwelvincen@gmail.com",
					Code:  "1234",
				},
			},
			want:    entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.CompareVerificationCode(tt.args.verificationCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.CompareVerificationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.CompareVerificationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_EditPassword(t *testing.T) {
	type args struct {
		id  string
		req *dto.EditDataUserRequest
	}
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockRepositoryInterface(ctrl)

	mockery.EXPECT().EditPassword("1", gomock.Any()).Return(entities.Account{}, nil).Times(1)
	tests := []struct {
		name    string
		u       UseCase
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			u: UseCase{
				Repo: mockery,
			},
			args: args{
				"1", &dto.EditDataUserRequest{
					Password: "123456",
				},
			},
			want:    entities.Account{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.EditPassword(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.EditPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.EditPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
