package test

import (
	"errors"
	"golang/helper"
	"golang/module/transaction"
	"reflect"
	"regexp"
	"testing"

	"gorm.io/gorm"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want transaction.RepositoryInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transaction.NewRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetAllLimit(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := transaction.Repository{
				DB: tt.fields.db,
			}
			got, got1, got2 := r.GetAllLimit(tt.args.input)
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

func TestRepository_GetAllTransaction(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type testCase struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}

	var tests []testCase

	mockQuery, mockDb := helper.NewMockQueryDb(t)

	name := "Test Error"
	f := fields{db: mockDb}
	query := regexp.QuoteMeta("SELECT * FROM `transactions`")
	mockQuery.ExpectQuery(query).WillReturnError(errors.New("Error"))
	tests = append(tests, testCase{
		name:    name,
		fields:  f,
		want:    "",
		wantErr: true,
	})

	// name = "Test Success"
	// version := "8"
	// mockQuery.ExpectQuery(query).WillReturnRows(
	// 	sqlmock.NewRows([]string{"version()"}).AddRow(version),
	// )
	// tests = append(tests, testCase{
	// 	name:    name,
	// 	fields:  f,
	// 	want:    version,
	// 	wantErr: false,
	// })

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := transaction.Repository{
				DB: tt.fields.db,
			}
			got, err := r.GetAllTransaction()
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

func TestRepository_GetAllTransactionByDate(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := transaction.Repository{
				DB: tt.fields.db,
			}
			got, err := r.GetAllTransactionByDate(tt.args.start, tt.args.end)
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

func TestRepository_GetAllTransactionByStatus(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := transaction.Repository{
				DB: tt.fields.db,
			}
			got, err := r.GetAllTransactionByStatus(tt.args.status)
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

func TestRepository_GetAllTransactionByStatusDate(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := transaction.Repository{
				DB: tt.fields.db,
			}
			got, err := r.GetAllTransactionByStatusDate(tt.args.status, tt.args.start, tt.args.end)
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
