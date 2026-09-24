package main

func test1(x int) int {
	if x > 10 {
		return x + 1
	} else {
		return x - 1
	}
}

func test2(a, b bool) bool {
	if a && b {
		return true
	}
	return false
}

func testArithmetic(x, y int) int {
	result := (x + y) * (x - y)

	if result < 0 {
		result = -result
	}

	return result % 10
}

func testUnary(x int) int {
	y := -x
	z := ^y

	if z > 0 {
		return z
	}
	return y
}

func testComparisons(a, b int) bool {
	if a == b {
		return true
	}

	if a != b && a > 0 {
		return b < 0
	}

	return a >= b || a <= -b
}

func testLogicalOps(x, y int) bool {
	cond1 := x > 0 && y > 0
	cond2 := x < 0 || y < 0

	return cond1 != cond2
}

func testWhileLoop(n int) int {
	i := 0
	sum := 0

	for i < n {
		sum += i
		i++
	}
	return sum
}

func testForLoop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result *= i
		} else {
			result += i
		}
	}
	return result
}

func testInfiniteLoopBreak(x int) int {
	i := 0
	for {
		if i >= x {
			break
		}
		i++
	}
	return i
}

func testLoopWithConcreteBoundAndSymbolicBranching(condition bool) int {
	result := 0
	for i := 0; i < 10; i++ {
		if condition && i%2 == 0 {
			result += i
		}
	}
	return result
}

func testLoopWithSymbolicBoundAndSymbolicBranching(n int, condition bool) int {
	if n > 10 {
		panic("Assumption violated: n should be less than or equal to 10")
	}

	result := 0
	for i := 0; i < n; i++ {
		if condition && i%2 == 0 {
			result += i
		}
	}
	return result
}

func testComplexConditions(a, b, c int) int {
	if (a > b && b < c) || (a < b && b > c) {
		return 1
	} else if a == b && b == c {
		return 2
	} else {
		return 3
	}
}

func testNestedIf(x, y int) int {
	if x > 0 {
		if y > 0 {
			return 1
		} else {
			return 2
		}
	} else {
		if y > 0 {
			return 3
		} else {
			return 4
		}
	}
}

func testBitwise(a, b int) int {
	and := a & b
	or := a | b
	xor := a ^ b

	if (and | or) > 0 {
		return xor << 1
	}
	return xor >> 1
}

func testCombined(x int) int {
	y := x * 2
	z := 0

	for i := 0; i < y; i++ {
		if i%3 == 0 {
			z += i
		} else if i%3 == 1 {
			z -= i
		} else {
			z *= 2
		}

		if z > 100 {
			break
		}
	}

	if z < 0 {
		return -z
	}
	return z
}

func testEdgeCases(a int) int {
	if a == 0 {
		return 0
	}

	if a < 0 {
		return -1
	}

	return 1
}

func testMultipleReturns(x int) int {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}

	return x * 2
}

func testSimpleSum(a, b float64) float64 {
	c := a + 1.1
	if b+c > 10.1 && b+c < 11.125 {
		return 1.1
	} else {
		return 1.2
	}
}
