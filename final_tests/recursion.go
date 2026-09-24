package final_tests

func Factorial(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	result := Factorial(n - 1)
	return n * result
}

func InfiniteRecursion(i int) {
	InfiniteRecursion(i + 1)
}
