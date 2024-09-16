package math

import (
	"testing"
)

func TestExponent(t *testing.T) {
	type args struct {
		x int64
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"Test case 1", args{2, 3}, 8},
		{"Test case 2", args{2, 10}, 1024},
		{"Test case 3", args{2, 0}, 1},
		{"Test case 4", args{0, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exponent(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("Exponent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	type args struct {
		x uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"Test case 1", args{0}, 1},
		{"Test case 2", args{1}, 1},
		{"Test case 3", args{5}, 120},
		{"Test case 4", args{10}, 3628800},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.x); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
