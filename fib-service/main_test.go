package main

import "testing"

func Test_fibonacci(t *testing.T) {
	type args struct {
		a uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
		{
			name: "zero",
			args: args{0},
			want: 0,
		},
		{
			name: "one",
			args: args{1},
			want: 1,
		},
		{
			name: "three",
			args: args{3},
			want: 2,
		},
		{
			name: "ten",
			args: args{10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.a); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
