package math

import (
	"testing"
)

func TestDefaultNormalize(t *testing.T) {
	tests := []struct {
		num  float64
		want int
	}{
		{
			num:  0,
			want: 0,
		},
		{
			num:  1,
			want: 1000000,
		},
		{
			num:  1.1,
			want: 1100000,
		},
		{
			num:  9,
			want: 9000000,
		},
		{
			num:  12,
			want: 12000000,
		},
		{
			num:  12.34,
			want: 12340000,
		},
		{
			num:  12.345678,
			want: 12345678,
		},
		{
			num:  12.3456781,
			want: 12345678,
		},
		{
			num:  12.3456789,
			want: 12345678,
		},
		{
			num:  -12.345678,
			want: -12345678,
		},
		{
			num:  789.12345678,
			want: 789123456,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DefaultNormalize(tt.num); got != tt.want {
				t.Errorf("DefaultNormalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultUnnormalize(t *testing.T) {
	tests := []struct {
		num  int
		want float64
	}{
		{
			num:  0,
			want: 0,
		},
		{
			num:  1,
			want: 0.000001,
		},
		{
			num:  9,
			want: 0.000009,
		},
		{
			num: 10,
			// 誤差
			want: 0.000009999999999999999,
		},
		{
			num:  1000000,
			want: 1,
		},
		{
			num: 1100000,
			// 誤差
			want: 1.0999999999999999,
		},
		{
			num:  9000000,
			want: 9,
		},
		{
			num:  12000000,
			want: 12,
		},
		{
			num:  12340000,
			want: 12.34,
		},
		{
			num:  12345678,
			want: 12.345678,
		},
		{
			num:  -12345678,
			want: -12.345678,
		},
		{
			num: 789123456,
			// 誤差
			want: 789.1234559999999,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DefaultUnnormalize(tt.num); got != tt.want {
				t.Errorf("DefaultUnnormalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
