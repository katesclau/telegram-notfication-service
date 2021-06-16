package utils

import "testing"

func TestBuildString(t *testing.T) {
	type args struct {
		all []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Build simple string",
			args{
				[]interface{}{"my ", "name ", "is ", "Earl"},
			},
			"my name is Earl",
		},
		{
			"Build string with int",
			args{
				[]interface{}{"my ", "age ", "is ", 32},
			},
			"my age is 32",
		},
		{
			"Build string with float",
			args{
				[]interface{}{"my ", "height ", "is ", 5.2},
			},
			"my height is 5.20",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildString(tt.args.all...); got != tt.want {
				t.Errorf("BuildString() = %v, want %v", got, tt.want)
			}
		})
	}
}
