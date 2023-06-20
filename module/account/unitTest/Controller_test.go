package unitTest

import (
	"golang/module/account"
	"reflect"
	"testing"
)

func TestController_GetDataUser(t *testing.T) {
	tests := []struct {
		name    string
		c       account.Controller
		want    *account.ReadResponse
		wantErr bool
	}{}
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
