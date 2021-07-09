package math

func DefaultNormalize(num float64) int {
	return normalize6(num)
}

func DefaultUnnormalize(num int) float64 {
	return unnormalize6(num)
}

func normalize6(num float64) int {
	return int(num * 1.0e06)
}

func unnormalize6(num int) float64 {
	return float64(num) * 1e-06
}
