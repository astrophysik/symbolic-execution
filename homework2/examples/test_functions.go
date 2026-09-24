// Примеры функций для тестирования конвертации символьных выражений в SMT запрос
package main

// Простая арифметическая функция
func add(a, b int) int {
	return a + b
}

// Функция с условием
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Функция с несколькими операциями
func calculate(x, y int) int {
	sum := x + y
	diff := x - y
	product := sum * diff
	return product
}

// Функция с логическими операциями
func isValid(x, y int) bool {
	return x > 0 && y > 0 && x < 100
}

// Функция с вложенными условиями
func classify(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else {
		return "F"
	}
}

// Функция с циклом (для продвинутого анализа)
func sum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Функция с комплексными булевыми выражениями
func complexCondition(x, y, z int) bool {
	return (x > 0 || y > 0) && (z > x+y)
}

// Функция с арифметическими операциями
func polynomial(x int) int {
	return 3*x*x + 2*x + 1
}

// Функция с модульными операциями
func modOperations(x, y int) bool {
	return x%2 == 0 && y%3 == 1
}

// Функция с сравнениями
func compareAll(a, b, c int) bool {
	return a == b || b != c || a < c || c >= b
}

// Функция для тестирования деления
func divide(x, y int) int {
	if y != 0 {
		return x / y
	}
	return 0
}

// Функция с битовыми операциями
func bitwiseOps(x, y int) int {
	return (x & y) | (x ^ y)
}

// Функция с унарными операциями
func unaryOps(x int, flag bool) int {
	result := -x
	if !flag {
		result = -result
	}
	return result
}

// Функция с тернарной логикой (через if-else)
func ternary(condition bool, a, b int) int {
	if condition {
		return a
	}
	return b
}

// Функция с вычислением абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Функция с несколькими return'ами
func signFunction(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

// Функция для проверки диапазона
func inRange(x, min, max int) bool {
	return x >= min && x <= max
}

func main() {
	// Функции для тестирования символьных выражений
	// Используются для создания CFG и символьных выражений
}
