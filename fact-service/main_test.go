package main

import "testing"

func Test_factorial(t *testing.T) {
	type args struct {
		a uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "zero",
			args: args{0},
			want: 1,
		},
		{
			name: "one",
			args: args{1},
			want: 1,
		},
		{
			name: "three",
			args: args{3},
			want: 6,
		},
		{
			name: "ten",
			args: args{7},
			want: 5040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorial(tt.args.a); got != tt.want {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
