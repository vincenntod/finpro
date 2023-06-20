package unitTest

import (
	"golang/module/account"
	"reflect"
	"testing"
)

func TestUseCase_GetDataUser(t *testing.T) {
	tests := []struct {
		name    string
		u       account.UseCase
		want    []account.Account
		wantErr bool
	}{}
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
