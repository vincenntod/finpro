package account

import (
	"golang/module/account/dto"
	"golang/module/account/entities"
	"golang/module/account/mocks"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestRepository_GetDataUser(t *testing.T) {

	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "SELECT * FROM `account`"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)

	tests := []struct {
		name    string
		r       Repository
		want    []entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			want: []entities.Account{{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			}},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "SELECT * FROM `account` WHERE `account`.`id` = ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				id: "1",
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	queryUpdate := "UPDATE account SET name=?,phone=?,role=?,password=? WHERE id = ?"
	queryFind := "SELECT * FROM account WHERE id = ?"

	mockQuery.ExpectQuery(queryUpdate).WillReturnRows(
		sqlmock.NewRows([]string{"name", "phone", "role", "password", "email", "id"}).AddRow("Vincen", "213", "Admin", "123", "admin@mail.com", 1),
	)
	mockQuery.ExpectQuery(queryFind).WillReturnRows(
		sqlmock.NewRows([]string{"name", "phone", "role", "password", "email", "id"}).AddRow("Vincen", "213", "Admin", "123", "admin@mail.com", 1),
	)

	type args struct {
		id  string
		req *entities.Account
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				id: "1",
				req: &entities.Account{
					Id:       1,
					Name:     "Vincen",
					Phone:    "213",
					Role:     "Admin",
					Password: "123",
					Email:    "admin@mail.com",
				},
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "DELETE FROM `account` WHERE id = ?"
	mockQuery.ExpectExec(query).WithArgs("1").WillReturnResult(sqlmock.NewResult(0, 1))

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "INSERT INTO account (name,phone,role,password,email) VALUES (?,?,?,?,?)"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"name", "phone", "role", "password", "email"}).AddRow("Vincen", "213", "Admin", "123", "admin@mail.com"),
	)

	type args struct {
		req *entities.Account
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				req: &entities.Account{
					Name:     "Vincen",
					Phone:    "213",
					Role:     "Admin",
					Password: "123",
					Email:    "admin@mail.com",
				},
			},
			want: entities.Account{
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "SELECT * FROM account WHERE email = ? ORDER BY account.id LIMIT 1"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "Role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)
	type args struct {
		req *entities.Account
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want1   entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				req: &entities.Account{
					Password: "123456",
					Email:    "admin@mail.com",
				},
			},
			want1: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got1, err := tt.r.Login(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Repository.Login() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRepository_SendEmail(t *testing.T) {
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "SELECT * FROM account WHERE email = ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				email: "maxwelvincen@gmail.com",
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	query := "SELECT * FROM account WHERE email = ?"
	mockQuery.ExpectQuery(query).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)
	type args struct {
		verificationCode *dto.VerificationCodeRequest
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				verificationCode: &dto.VerificationCodeRequest{
					Email: "maxwelvincen@gmail.com",
					Code:  "1234",
				},
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
	mockQuery, mockDb := mocks.NewMockQueryDb(t)
	queryUpdate := "UPDATE account SET password=? WHERE id = ?"
	queryFind := "SELECT * FROM account WHERE account.id = ?"
	mockQuery.ExpectQuery(queryUpdate).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)
	mockQuery.ExpectQuery(queryFind).WillReturnRows(
		sqlmock.NewRows([]string{"id", "name", "phone", "role", "password", "email"}).AddRow(1, "Vincen", "213", "Admin", "123", "admin@mail.com"),
	)
	type args struct {
		id  string
		req *entities.Account
	}
	tests := []struct {
		name    string
		r       Repository
		args    args
		want    entities.Account
		wantErr bool
	}{
		{
			name: "success",
			r: Repository{
				Db: mockDb,
			},
			args: args{
				id: "1",
				req: &entities.Account{
					Id:       1,
					Name:     "Vincen",
					Phone:    "213",
					Role:     "Admin",
					Password: "123",
					Email:    "admin@mail.com",
				},
			},
			want: entities.Account{
				Id:       1,
				Name:     "Vincen",
				Phone:    "213",
				Role:     "Admin",
				Password: "123",
				Email:    "admin@mail.com",
			},
			wantErr: false,
		},
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
