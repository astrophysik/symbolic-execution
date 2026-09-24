// Простые тестовые функции для построения SSA
package main

// Функция с простым условным оператором
func simpleIf(x int) int {
	if x > 0 {
		return x * 2
	}
	return 0
}

// Функция с if-else
func ifElse(x int) int {
	if x > 0 {
		return x * 2
	} else {
		return x * -1
	}
	// Недостижимый код
	return 999
}

// Функция с вложенными условиями
func nestedIf(x, y int) int {
	if x > 0 {
		if y > 0 {
			return x + y
		}
		return x
	}
	return 0
}

// Функция с простым циклом
func simpleLoop(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}

// Функция с while-циклом
func whileLoop(x int) int {
	for x > 0 {
		x = x - 1
	}
	return x
}

// Функция с вложенными циклами
func nestedLoops(n, m int) int {
	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += i * j
		}
	}
	return sum
}

// Функция с break и continue
func loopWithBreakContinue(arr []int) int {
	sum := 0
	for i, val := range arr {
		if i%2 == 0 {
			continue
		}
		if val < 0 {
			break
		}
		sum += val
	}
	return sum
}

// Функция с switch-case
func switchExample(x int) string {
	switch x {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "two"
	default:
		return "other"
	}
}

// Функция с множественными return'ами
func multipleReturns(x, y int) int {
	if x > y {
		return x
	}
	if x < y {
		return y
	}
	if x == 0 {
		return -1
	}
	return 0
}

// Функция с panic
func withPanic(x int) int {
	if x < 0 {
		panic("negative value")
	}
	if x == 0 {
		return 1
	}
	return x * x
}

// Рекурсивная функция
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Функция с goto (сложный CFG)
func withGoto(x int) int {
	if x > 10 {
		goto big
	}
	x *= 2
	goto end

big:
	x *= 10

end:
	return x
}

// Функция с defer (влияет на CFG)
func withDefer(x int) int {
	defer func() {
		// cleanup code
	}()

	if x < 0 {
		return -1
	}
	return x * 2
}

// Функция с комплексными булевыми выражениями
func complexConditions(x, y, z int) bool {
	return (x > 0 && y > 0) || (z > 0 && x+y > z)
}

func main() {
	// Функции для тестирования - не вызываются
	// Используются только для анализа CFG
}
