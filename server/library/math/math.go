package math

import (
	"math"
)

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func Max(x float64, y ...float64) float64 {
	max := x
	for _, v := range y {
		if v > max {
			max = v
		}
	}
	return max
}

func Min(x float64, y ...float64) float64 {
	min := x
	for _, v := range y {
		if v < min {
			min = v
		}
	}
	return min
}

func Round(x, digit float64) float64 {
	shift := math.Pow(10, digit)
	if x > 0.0 {
		return math.Trunc(x*shift+0.5) / shift
	}
	return math.Trunc(x*shift-0.5) / shift
}

func IsOdd(num interface{}) bool {
	switch num := num.(type) {
	case int:
		return num%2 == 1
	case int64:
		return num%2 == 1
	case int32:
		return num%2 == 1
	case int16:
		return num%2 == 1
	case int8:
		return num%2 == 1
	case uint:
		return num%2 == 1
	case uint64:
		return num%2 == 1
	case uint32:
		return num%2 == 1
	case uint16:
		return num%2 == 1
	case uint8:
		return num%2 == 1
	}
	return false
}

func IntBetween(n, min, max int) bool {
	return n >= min && max >= n
}
