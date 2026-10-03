// Демонстрационная программа для тестирования символьных выражений
package main

import (
	"fmt"
	"log"

	"symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"
)

func main() {
	fmt.Println("=== Symbolic Expressions Demo ===")

	// Создаём простые символьные выражения
	x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
	y := symbolic.NewSymbolicVariable("y", symbolic.IntType)
	five := symbolic.NewIntConstant(5)

	// Создаём выражение: x + y > 5
	sum := symbolic.NewBinaryOperation(x, y, symbolic.ADD)
	condition := symbolic.NewBinaryOperation(sum, five, symbolic.GT)

	fmt.Printf("Выражение: %s\n", condition.String())
	fmt.Printf("Тип выражения: %s\n", condition.Type().String())

	// Создаём Z3 транслятор
	translator := translator.NewZ3Translator()
	defer translator.Close()

	// Транслируем в Z3
	z3Expr, err := translator.TranslateExpression(condition)
	if err != nil {
		log.Fatalf("Ошибка трансляции: %v", err)
	}

	fmt.Printf("Z3 выражение создано: %T\n", z3Expr)

	// Создаём более сложное выражение: (x > 0) && (y < 10)
	zero := symbolic.NewIntConstant(0)
	ten := symbolic.NewIntConstant(10)

	cond1 := symbolic.NewBinaryOperation(x, zero, symbolic.GT)
	cond2 := symbolic.NewBinaryOperation(y, ten, symbolic.LT)

	andExpr := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{cond1, cond2}, symbolic.AND)

	fmt.Printf("Сложное выражение: %s\n", andExpr.String())

	// Транслируем сложное выражение
	z3AndExpr, err := translator.TranslateExpression(andExpr)
	if err != nil {
		log.Fatalf("Ошибка трансляции сложного выражения: %v", err)
	}

	fmt.Printf("Сложное Z3 выражение создано: %T\n", z3AndExpr)

	testArithmetic(translator)
	testRepeatedVariable(translator)
	testLogicalOperations(translator)
	testArrayOperations(translator)
}

func testArithmetic(zt *translator.Z3Translator) {
	fmt.Println("\n=== Арифметика ===")

	x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
	two := symbolic.NewIntConstant(2)
	three := symbolic.NewIntConstant(3)
	sum := symbolic.NewBinaryOperation(x, two, symbolic.ADD)
	expression := symbolic.NewBinaryOperation(sum, three, symbolic.MUL)

	fmt.Printf("Выражение: %s\n", expression.String())
	fmt.Printf("Тип выражения: %s\n", expression.Type().String())

	z3Expr, err := zt.TranslateExpression(expression)
	if err != nil {
		log.Fatalf("Ошибка трансляции арифметики: %v", err)
	}
	fmt.Printf("Z3 выражение (%T): %v\n", z3Expr, z3Expr)
}

func testRepeatedVariable(zt *translator.Z3Translator) {
	fmt.Println("\n=== Повторная переменная ===")

	x := symbolic.NewSymbolicVariable("x", symbolic.IntType)
	xAgain := symbolic.NewSymbolicVariable("x", symbolic.IntType)
	expression := symbolic.NewBinaryOperation(x, xAgain, symbolic.SUB)

	fmt.Printf("Выражение: %s\n", expression.String())
	fmt.Printf("Тип выражения: %s\n", expression.Type().String())

	z3Expr, err := zt.TranslateExpression(expression)
	if err != nil {
		log.Fatalf("Ошибка трансляции повторной переменной: %v", err)
	}
	fmt.Printf("Z3 выражение (%T): %v\n", z3Expr, z3Expr)
}

func testLogicalOperations(zt *translator.Z3Translator) {
	fmt.Println("\n=== Логические операции ===")

	p := symbolic.NewSymbolicVariable("p", symbolic.BoolType)
	q := symbolic.NewSymbolicVariable("q", symbolic.BoolType)
	pOrQ := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{p, q}, symbolic.OR)
	notP := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{p}, symbolic.NOT)
	expression := symbolic.NewLogicalOperation([]symbolic.SymbolicExpression{pOrQ, notP}, symbolic.AND)

	fmt.Printf("Выражение: %s\n", expression.String())
	fmt.Printf("Тип выражения: %s\n", expression.Type().String())

	z3Expr, err := zt.TranslateExpression(expression)
	if err != nil {
		log.Fatalf("Ошибка трансляции логических операций: %v", err)
	}
	fmt.Printf("Z3 выражение (%T): %v\n", z3Expr, z3Expr)
}

func testArrayOperations(zt *translator.Z3Translator) {
	fmt.Println("\n=== Массив целых чисел ===")

	expression := symbolic.NewSymbolicVariable("A", symbolic.ArrayType)
	expression.TypeInfo.InnerType = &symbolic.NestedExpressionType{ExprType: symbolic.IntType}

	fmt.Printf("Выражение: %s\n", expression.String())
	fmt.Printf("Тип выражения: %s\n", expression.Type().String())
	fmt.Printf("Тип элементов: %s\n", expression.TypeInfo.InnerType.ExprType.String())

	z3Expr, err := zt.TranslateExpression(expression)
	if err != nil {
		log.Fatalf("Ошибка трансляции массива: %v", err)
	}
	fmt.Printf("Z3 выражение (%T): %v\n", z3Expr, z3Expr)
}
