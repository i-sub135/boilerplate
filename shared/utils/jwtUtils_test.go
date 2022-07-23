package utils

import (
	"fmt"
	"testing"
)

func TestCreateJwt(t *testing.T) {
	type args struct {
		param map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "be able create jwt",
			args:    args{param: map[string]interface{}{"name": "customer_name"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateJwt(tt.args.param)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateJwt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			fmt.Println(got)
		})
	}
}
