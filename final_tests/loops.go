package final_tests

func LoopWithConcreteBound(n int) int {
	result := 0
	for i := 0; i < 10; i++ {
		result += i
	}
	return result
}

func LoopWithSymbolicBound(n int) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	return result
}

func LoopWithSymbolicBoundAndSymbolicBranching(n int, condition bool) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		if condition && i%2 == 0 {
			result += i
		}
	}
	return result
}

func LoopWithSymbolicBoundAndComplexControlFlow(n int, condition bool) int {
	if n > 10 {
		return -1
	}

	result := 0
	for i := 0; i < n; i++ {
		if condition && i == 3 {
			break
		}
		if i%2 != 0 {
			continue
		}
		result += i
	}
	return result
}

func WhileCycle(x int) int {
	i := 0
	sum := 0
	for i < x {
		sum += i
		i++
	}
	return sum
}

func LoopInsideLoop(x int) int {
	for i := x - 5; i < x; i++ {
		if i < 0 {
			return 2
		} else {
			for j := i; j < x+i; j++ {
				if j == 7 {
					return 1
				}
			}
		}
	}
	return -1
}
