package unitTest

import (
	"golang/module/account"
	"reflect"
	"testing"
)

func TestRepository_GetDataUser(t *testing.T) {
	tests := []struct {
		name    string
		r       account.Repository
		want    []account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetDataUser()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetDataUserById(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetDataUserById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetDataUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetDataUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_EditDataUser(t *testing.T) {
	type args struct {
		id  string
		req *account.Account
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.EditDataUser(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.EditDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.EditDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_DeleteDataUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.DeleteDataUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.DeleteDataUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.DeleteDataUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CreateAccount(t *testing.T) {
	type args struct {
		req *account.Account
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CreateAccount(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Login(t *testing.T) {
	type args struct {
		req *account.Account
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    string
		want1   account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.r.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repository.Login() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Repository.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGenerateVerificationCode(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := account.GenerateVerificationCode(tt.args.email); got != tt.want {
				t.Errorf("GenerateVerificationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SendEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.SendEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.SendEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_CompareVerificationCode(t *testing.T) {
	type args struct {
		verificationCode *account.VerificationCodeRequest
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.CompareVerificationCode(tt.args.verificationCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.CompareVerificationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.CompareVerificationCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_EditPassword(t *testing.T) {
	type args struct {
		id  string
		req *account.Account
	}
	tests := []struct {
		name    string
		r       account.Repository
		args    args
		want    account.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.EditPassword(tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.EditPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.EditPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
