// Демонстрационная программа для тестирования символьных выражений
package main

import (
	"fmt"
	"log"

	"github.com/ebukreev/go-z3/z3"

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
	fmt.Println("\nВсе проверки пройдены.")
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
	result, ok := z3Expr.(z3.Int)
	check("арифметика имеет тип int", expression.Type() == symbolic.IntType && ok)
	ctx := zt.GetContext().(*z3.Context)
	xIsFour := ctx.IntConst("x").Eq(ctx.FromInt(4, ctx.IntSort()).(z3.Int))
	checkSatisfiable(ctx, "x = 4 выполнимо", true, xIsFour)
	checkSatisfiable(ctx, "при x = 4 выражение (x + 2) * 3 равно 18", false,
		xIsFour, result.NE(ctx.FromInt(18, ctx.IntSort()).(z3.Int)))
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
	result, ok := z3Expr.(z3.Int)
	check("разность имеет тип int", expression.Type() == symbolic.IntType && ok)
	ctx := zt.GetContext().(*z3.Context)
	checkSatisfiable(ctx, "разность двух узлов с именем x всегда равна нулю", false,
		result.NE(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
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
	result, ok := z3Expr.(z3.Bool)
	check("логическое выражение имеет тип bool", expression.Type() == symbolic.BoolType && ok)
	ctx := zt.GetContext().(*z3.Context)
	cases := []struct {
		p, q, want bool
	}{
		{false, false, false},
		{false, true, true},
		{true, false, false},
		{true, true, false},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("p = %t, q = %t", tc.p, tc.q)
		pValue := ctx.BoolConst("p").Eq(ctx.FromBool(tc.p))
		qValue := ctx.BoolConst("q").Eq(ctx.FromBool(tc.q))
		checkSatisfiable(ctx, name+": значения выполнимы", true, pValue, qValue)
		checkSatisfiable(ctx, name+fmt.Sprintf(": результат равен %t", tc.want), false,
			pValue, qValue, result.NE(ctx.FromBool(tc.want)))
	}
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
	result, ok := z3Expr.(z3.Array)
	check("выражение имеет тип array", expression.Type() == symbolic.ArrayType && ok)
	domain, element := result.Sort().DomainAndRange()
	check("индексы массива имеют тип int", domain.Kind() == z3.KindInt)
	check("элементы массива имеют тип int", element.Kind() == z3.KindInt)
}

func check(name string, passed bool) {
	if !passed {
		log.Fatalf("FAIL: %s", name)
	}
	fmt.Printf("PASS: %s\n", name)
}

func checkSatisfiable(ctx *z3.Context, name string, want bool, constraints ...z3.Bool) {
	solver := z3.NewSolver(ctx)
	for _, constraint := range constraints {
		solver.Assert(constraint)
	}
	sat, err := solver.Check()
	if err != nil {
		log.Fatalf("FAIL: %s: ошибка solver: %v", name, err)
	}
	if sat != want {
		log.Fatalf("FAIL: %s: ожидалась выполнимость %t, получена %t", name, want, sat)
	}
	check(name, true)
}
