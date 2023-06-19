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
	}{}
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
