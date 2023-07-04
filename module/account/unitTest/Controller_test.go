package unitTest

import (
	"golang/module/account"
	"golang/module/account/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestController_GetDataUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockery := mocks.NewMockUseCaseInterface(ctrl)

	mockery.EXPECT().GetDataUser().Return(&account.ReadResponse{
		Message: "Success",
		Code:    200,
		Status:  "OK",
		Data: []account.AccountItemResponse{{
			Id:    1,
			Name:  "Vincen",
			Role:  "Admin",
			Email: "vincen@example.com",
			Phone: "08211",
		},
		},
	}).Times(1)
	tests := []struct {
		name    string
		c       account.Controller
		want    *account.ReadResponse
		wantErr bool
	}{
		{
			name: "Success",
			c: account.Controller{
				UseCase: mockery,
			},
			want: &account.ReadResponse{
				Message: "Success",
				Code:    200,
				Status:  "OK",
				Data: []account.AccountItemResponse{{
					Id:    1,
					Name:  "Vincen",
					Role:  "Admin",
					Email: "vincen@example.com",
					Phone: "08211",
				},
				},
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
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.ReadResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		req *account.CreateRequest
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		id  string
		req *account.EditDataUserRequest
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditDataUser(tt.args.id, tt.args.req)
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
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
		req *account.LoginResponseRequest
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    string
		want1   *account.LoginResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.c.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Controller.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Controller.Login() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Controller.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestController_SendEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		verificationCode *account.VerificationCodeRequest
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
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
	type args struct {
		code string
		req  *account.EditDataUserRequest
	}
	tests := []struct {
		name    string
		c       account.Controller
		args    args
		want    *account.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.EditPassword(tt.args.code, tt.args.req)
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
