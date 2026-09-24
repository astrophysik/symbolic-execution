package final_tests

func CompareWithDiv(a, b float64) float64 {
	z := a + 0.5
	if (a / z) > b {
		return 1.0
	} else {
		return 0.0
	}
}

func Mul(a, b float64) float64 {
	if a*b > 33.32 && a*b < 33.333 {
		return 1.1
	} else if a*b > 33.333 && a*b < 33.7592 {
		return 1.2
	} else {
		return 1.3
	}
}
