package math

import (
	"reflect"
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

func TestRoundToFloat(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		decimals int
		expected float64
	}{
		{
			name:     "Round float64",
			input:    123.456789,
			decimals: 2,
			expected: 123.46,
		},
		{
			name:     "Round int",
			input:    123,
			decimals: 2,
			expected: 123.00,
		},
		{
			name:     "Round negative float64",
			input:    -123.456789,
			decimals: 2,
			expected: -123.46,
		},
		{
			name:     "Round float64 with zero decimals",
			input:    123.456789,
			decimals: 0,
			expected: 123.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result float64
			switch v := tt.input.(type) {
			case float64:
				result = RoundToFloat(v, tt.decimals)
			case int:
				result = RoundToFloat(v, tt.decimals)
			default:
				t.Fatalf("unsupported type %T", v)
			}

			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPercent(t *testing.T) {
	type args struct {
		val   float64
		total float64
		n     int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Both values are zero", args{0, 0, 2}, 0},
		{"Value is zero", args{0, 100, 2}, 0},
		{"Total is zero", args{50, 0, 2}, 0},
		{"Normal case", args{50, 100, 2}, 50.00},
		{"Decimal case", args{1, 3, 3}, 33.333},
		{"Negative values", args{-50, 100, 2}, -50.00},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Percent(tt.args.val, tt.args.total, tt.args.n); got != tt.want {
				t.Errorf("Percent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundToString(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, "123.46"},
		{"Round int to 2 decimals", args{123, 2}, "123.00"},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, "-123.46"},
		{"Round float64 to 0 decimals", args{123.456789, 0}, "123"},
		{"Round float64 to 5 decimals", args{123.456789, 5}, "123.45679"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			switch v := tt.args.x.(type) {
			case float64:
				got = RoundToString(v, tt.args.n)
			case int:
				got = RoundToString(v, tt.args.n)
			}

			if got != tt.want {
				t.Errorf("RoundToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncRound(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, 123.45},
		{"Round int to 2 decimals", args{123, 2}, 123},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, -123.45},
		{"Round float64 to 0 decimals", args{123.456789, 0}, 123.0},
		{"Round float64 to 5 decimals", args{123.456789, 5}, 123.45678},
		{"Round with n = 0 for integer", args{123, 0}, 123},
		{"Round with n < 0", args{123.456789, -1}, 123.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.args.x.(type) {
			case float64:
				if got := TruncRound(v, tt.args.n); got != tt.want {
					t.Errorf("TruncRound() = %v, want %v", got, tt.want)
				}
			case int:
				if got := TruncRound(v, tt.args.n); got != tt.want {
					t.Errorf("TruncRound() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFloorToFloat(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, 123.45},
		{"Round int to 2 decimals", args{123, 2}, 123.00},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, -123.46},
		{"Round float64 to 0 decimals", args{123.456789, 0}, 123.0},
		{"Round float64 to 5 decimals", args{123.456789, 5}, 123.45678},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch x := tt.args.x.(type) {
			case float64:
				if got := FloorToFloat(x, tt.args.n); got != tt.want {
					t.Errorf("FloorToFloat() = %v, want %v", got, tt.want)
				}
			case int:
				if got := FloorToFloat(x, tt.args.n); got != tt.want {
					t.Errorf("FloorToFloat() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFloorToStr(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, "123.45"},
		{"Round int to 2 decimals", args{123, 2}, "123.00"},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, "-123.46"},
		{"Round float64 to 0 decimals", args{123.456789, 0}, "123"},
		{"Round float64 to 5 decimals", args{123.456789, 5}, "123.45678"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch x := tt.args.x.(type) {
			case float64:
				if got := FloorToStr(x, tt.args.n); got != tt.want {
					t.Errorf("FloorToStr() = %v, want %v", got, tt.want)
				}
			case int:
				if got := FloorToStr(x, tt.args.n); got != tt.want {
					t.Errorf("FloorToStr() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCeilToFloat(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, 123.46},
		{"Round int to 2 decimals", args{123, 2}, 123.00},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, -123.45},
		{"Round float64 to 0 decimals", args{123.456789, 0}, 124.0},
		{"Round float64 to 5 decimals", args{123.456789, 5}, 123.45679},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch x := tt.args.x.(type) {
			case float64:
				if got := CeilToFloat(x, tt.args.n); got != tt.want {
					t.Errorf("CeilToFloat() = %v, want %v", got, tt.want)
				}
			case int:
				if got := CeilToFloat(x, tt.args.n); got != tt.want {
					t.Errorf("CeilToFloat() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestCeilToStr(t *testing.T) {
	type args struct {
		x any
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Round float64 to 2 decimals", args{123.456789, 2}, "123.46"},
		{"Round int to 2 decimals", args{123, 2}, "123.00"},
		{"Round negative float64 to 2 decimals", args{-123.456789, 2}, "-123.45"},
		{"Round float64 to 0 decimals", args{123.456789, 0}, "124"},
		{"Round float64 to 5 decimals", args{123.456789, 5}, "123.45679"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch x := tt.args.x.(type) {
			case float64:
				if got := CeilToStr(x, tt.args.n); got != tt.want {
					t.Errorf("CeilToStr() = %v, want %v", got, tt.want)
				}
			case int:
				if got := CeilToStr(x, tt.args.n); got != tt.want {
					t.Errorf("CeilToStr() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestSum(t *testing.T) {
	t.Run("Integer", func(t *testing.T) {
		if got := Sum([]int{1, 2, 3, 4, 5}...); !reflect.DeepEqual(got, 15) {
			t.Errorf("Sum() = %v, want %v", got, 15)
		}
	})

	t.Run("nil", func(t *testing.T) {
		if got := Sum([]int{}...); !reflect.DeepEqual(got, 0) {
			t.Errorf("Sum() = %v, want %v", got, 0)
		}
	})

	t.Run("Float64", func(t *testing.T) {
		if got := Sum([]float64{
			1.1, 2.2, 3.3, 4.4, 5.5,
		}...); !reflect.DeepEqual(got, 16.5) {
			t.Errorf("Sum() = %v, want %v", got, 16.5)
		}
	})
}
