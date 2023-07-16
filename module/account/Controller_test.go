package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
	"golang/module/account/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestController_GetDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().GetDataUser().Return([]entities.Account{{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "admin@example.com",
		Phone: "123",
	}}, nil).Times(1)

	tests := []struct {
		name    string
		c       Controller
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			want: &dto.MessageResponse{
				Message: "Success Get Data",
				Code:    200,
				Status:  "OK",
				Data: []dto.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "admin@example.com",
					Phone: "123",
				}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetDataUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.GetDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.GetDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_GetDataUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().GetDataUserById("1").Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "admin@example.com",
		Phone: "123",
	}, nil).Times(1)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				id: "1",
			},
			want: &dto.MessageResponse{
				Message: "Success Get Data",
				Code:    200,
				Status:  "OK",
				Data: []dto.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "admin@example.com",
					Phone: "123",
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GetDataUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.GetDataUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.GetDataUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_CreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().CreateAccount(gomock.Any()).Return(entities.Account{
		Id:    0,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwel@example.com",
		Phone: "123",
	}, nil).Times(1)

	type args struct {
		req *dto.CreateRequest
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				req: &dto.CreateRequest{
					Name:     "Vincen",
					Phone:    "123",
					Role:     "admin",
					Password: "123456",
					Email:    "maxwel@gmail.com",
				},
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Save Data",
				Data: []dto.AccountItemResponse{{
					Id:    0,
					Name:  "Vincen",
					Role:  "admin",
					Email: "maxwel@gmail.com",
					Phone: "123",
				},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CreateAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_EditDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().EditDataUser("1", gomock.Any()).Return(entities.Account{
		Name:     "Vincen",
		Password: "123456",
		Role:     "Admin",
		Email:    "maxwel@gmail.com",
		Phone:    "123",
	}, nil).Times(1)
	type args struct {
		id  string
		req *dto.EditDataUserRequest
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				id: "1",
				req: &dto.EditDataUserRequest{
					Name:     "Vincen",
					Phone:    "123",
					Role:     "Admin",
					Password: "123456",
				},
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Edit Data",
				Data: []dto.AccountItemResponse{{
					Name:  "Vincen",
					Role:  "Admin",
					Email: "maxwel@gmail.com",
					Phone: "123",
				},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditDataUser(tt.args.id, tt.args.req)
			tt.args.req.Password = "123456"
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.EditDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.EditDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_DeleteDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().DeleteDataUser("1").Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwel@example.com",
		Phone: "123",
	}, nil).Times(1)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				id: "1",
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Delete Data",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.DeleteDataUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.DeleteDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.DeleteDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_Login(t *testing.T) {

	type args struct {
		req *dto.LoginResponseRequest
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.LoginResponse
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.c.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_SendEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().SendEmail("maxwelvincen@gmail.com").Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwelvincen@gmail.com",
		Phone: "123",
	}, nil).Times(1)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Send Email",
				Data: []dto.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "maxwelvincen@gmail.com",
					Phone: "123",
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.SendEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.SendEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_SendEmailRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().SendEmail("maxwelvincen@gmail.com").Return(entities.Account{}, nil).Times(1)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Send Email",
				Data:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.SendEmailRegister(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.SendEmailRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.SendEmailRegister() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_CompareVerificationCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	entities.VerificationCodes["maxwelvincen@gmail.com"] = "123456"
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	mockery.EXPECT().CompareVerificationCode(&dto.VerificationCodeRequest{
		Email: "maxwelvincen@gmail.com",
		Code:  "123456",
	}).Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwelvincen@gmail.com",
		Phone: "123",
	}, nil).Times(1)
	type args struct {
		verificationCode *dto.VerificationCodeRequest
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{
			name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				verificationCode: &dto.VerificationCodeRequest{
					Email: "maxwelvincen@gmail.com",
					Code:  "123456",
				},
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Verification Code",
				Data: []dto.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "maxwelvincen@gmail.com",
					Phone: "123",
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.CompareVerificationCode(tt.args.verificationCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.CompareVerificationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.CompareVerificationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestController_EditPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)
	entities.VerificationCodes["maxwelvincen@gmail.com"] = "123456"
	mockery.EXPECT().EditPassword("1", gomock.Any()).Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwelvincen@gmail.com",
		Phone: "123",
	}, nil).Times(1)
	mockery.EXPECT().GetDataUserById("1").Return(entities.Account{
		Id:    1,
		Name:  "Vincen",
		Role:  "Admin",
		Email: "maxwelvincen@gmail.com",
		Phone: "123",
	}, nil).Times(1)
	type args struct {
		id   string
		code string
		req  *dto.EditDataUserRequest
	}
	tests := []struct {
		name    string
		c       Controller
		args    args
		want    *dto.MessageResponse
		wantErr bool
	}{
		{name: "success",
			c: Controller{
				UseCase: mockery,
			},
			args: args{
				id:   "1",
				code: "123456",
				req: &dto.EditDataUserRequest{
					Id:       1,
					Password: "123456",
					Name:     "Vincen",
					Role:     "Admin",
					Phone:    "123",
				},
			},
			want: &dto.MessageResponse{
				Code:    200,
				Status:  "OK",
				Message: "Success Update Password",
				Data: []dto.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "maxwelvincen@gmail.com",
					Phone: "123",
				},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditPassword(tt.args.id, tt.args.code, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.EditPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Controller.EditPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
