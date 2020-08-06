package root_dir

import "testing"

func Test_inferRootDir(t *testing.T) {
	type args struct {
		sub string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{sub: "/utils"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inferRootDir(tt.args.sub); got != tt.want {
				t.Errorf("inferRootDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
