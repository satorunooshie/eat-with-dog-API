package math

import (
	"testing"
)

func TestAbsInt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{
			x:    100,
			want: 100,
		},
		{
			x:    0,
			want: 0,
		},
		{
			x:    -100,
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := AbsInt(tt.x); got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntBetween(t *testing.T) {
	type args struct {
		n   int
		min int
		max int
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				n:   1,
				min: 0,
				max: 3,
			},
			want: true,
		},
		{
			args: args{
				n:   1,
				min: 3,
				max: 6,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IntBetween(tt.args.n, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("IntBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	tests := []struct {
		num  interface{}
		want bool
	}{
		{
			num:  1,
			want: true,
		},
		{
			num:  3,
			want: true,
		},
		{
			num:  5,
			want: true,
		},
		{
			num:  7,
			want: true,
		},
		{
			num:  9,
			want: true,
		},
		{
			num:  11,
			want: true,
		},
		{
			num:  0,
			want: false,
		},
		{
			num:  2,
			want: false,
		},
		{
			num:  4,
			want: false,
		},
		{
			num:  6,
			want: false,
		},
		{
			num:  8,
			want: false,
		},
		{
			num:  10,
			want: false,
		},
		{
			num:  12,
			want: false,
		},
		{
			num:  int64(1),
			want: true,
		},
		{
			num:  int64(2),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsOdd(tt.num); got != tt.want {
				t.Errorf("IsOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		x float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "single x is bigger",
			args: args{
				x: 0.2,
				y: []float64{
					0.1,
				},
			},
			want: 0.2,
		},
		{
			name: "single y is bigger",
			args: args{
				x: 0.1,
				y: []float64{
					0.2,
				},
			},
			want: 0.2,
		},
		{
			name: "single minus value",
			args: args{
				x: -9.9,
				y: []float64{
					-10.2,
				},
			},
			want: -9.9,
		},
		{
			name: "single same value",
			args: args{
				x: 1.1,
				y: []float64{
					1.1,
				},
			},
			want: 1.1,
		},
		{
			name: "multiple x is bigger",
			args: args{
				x: 0.2,
				y: []float64{
					0.0,
					0.1,
				},
			},
			want: 0.2,
		},
		{
			name: "multiple y is bigger",
			args: args{
				x: 0.1,
				y: []float64{
					0.0,
					0.2,
					0.3,
				},
			},
			want: 0.3,
		},
		{
			name: "multiple minus value",
			args: args{
				x: -9.9,
				y: []float64{
					-100.2,
					-22.5,
					-828.828,
					-3.0,
				},
			},
			want: -3.0,
		},
		{
			name: "multiple same value",
			args: args{
				x: 2.2,
				y: []float64{
					2.2,
					2.2,
					2.2,
				},
			},
			want: 2.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.x, tt.args.y...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args struct {
		x float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "single x is smaller",
			args: args{
				x: 0.1,
				y: []float64{
					0.2,
				},
			},
			want: 0.1,
		},
		{
			name: "single y is smaller",
			args: args{
				x: 0.2,
				y: []float64{
					0.1,
				},
			},
			want: 0.1,
		},
		{
			name: "single minus value",
			args: args{
				x: -9.9,
				y: []float64{
					-10.2,
				},
			},
			want: -10.2,
		},
		{
			name: "single same value",
			args: args{
				x: 1.1,
				y: []float64{
					1.1,
				},
			},
			want: 1.1,
		},
		{
			name: "multiple x is smaller",
			args: args{
				x: 0.0,
				y: []float64{
					0.2,
					0.1,
				},
			},
			want: 0.0,
		},
		{
			name: "multiple y is smaller",
			args: args{
				x: 0.1,
				y: []float64{
					0.0,
					0.2,
					0.3,
				},
			},
			want: 0.0,
		},
		{
			name: "multiple minus value",
			args: args{
				x: -9.9,
				y: []float64{
					-100.2,
					-22.5,
					-828.828,
					-3.0,
				},
			},
			want: -828.828,
		},
		{
			name: "multiple same value",
			args: args{
				x: 0.1,
				y: []float64{
					0.1,
					0.1,
					0.1,
				},
			},
			want: 0.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.x, tt.args.y...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		x     float64
		digit float64
	}
	tests := []struct {
		args args
		want float64
	}{
		{
			args: args{
				x:     0.0,
				digit: 0,
			},
			want: 0.0,
		},
		{
			args: args{
				x:     0.4,
				digit: 0,
			},
			want: 0.0,
		},
		{
			args: args{
				x:     0.5,
				digit: 0,
			},
			want: 1.0,
		},
		{
			args: args{
				x:     5.4,
				digit: 0,
			},
			want: 5.0,
		},
		{
			args: args{
				x:     5.5,
				digit: 0,
			},
			want: 6.0,
		},
		{
			args: args{
				x:     -5.4,
				digit: 0,
			},
			want: -5.0,
		},
		{
			args: args{
				x:     -5.5,
				digit: 0,
			},
			want: -6.0,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Round(tt.args.x, tt.args.digit); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
