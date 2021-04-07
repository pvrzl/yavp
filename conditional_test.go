package yavp

import (
	"errors"
	"testing"
)

func TestWhen(t *testing.T) {
	type args struct {
		condition bool
		validator error
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "skip if condition false",
			args: args{
				condition: false,
				validator: nil,
			},
			wantErr: false,
		},
		{
			name: "check error if condition true",
			args: args{
				condition: true,
				validator: errors.New("something"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := When(tt.args.condition, tt.args.validator); (err != nil) != tt.wantErr {
				t.Errorf("When() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
