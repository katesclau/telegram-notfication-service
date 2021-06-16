package utils

import "testing"

func TestKeyFromPath(t *testing.T) {
	type args struct {
		path     string
		position int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Get Existing Key from Path",
			args{
				"/topics/key",
				2,
			},
			"key",
		},
		{
			"Get Nothing, when non existing Key in Path",
			args{
				"/topics/",
				2,
			},
			"",
		},
		{
			"Get Existing Key from Path even with query params",
			args{
				"/topics/key?query=params",
				2,
			},
			"key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyFromPath(tt.args.path, tt.args.position); got != tt.want {
				t.Errorf("KeyFromPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
