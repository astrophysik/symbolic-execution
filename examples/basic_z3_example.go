// Пример 1: Основы работы с Z3 - решение простых уравнений
package main

import (
	"fmt"
	"log"

	"symbolic-execution-course/pkg/z3wrapper"
)

func main() {
	fmt.Println("=== Пример 1: Основы Z3 - решение уравнений ===")

	// Создаём solver
	solver := z3wrapper.NewSolver()
	defer solver.Close()

	// Пример 1: Найти x, y такие что x + y = 10 и x - y = 2
	fmt.Println("\nЗадача: найти x, y такие что x + y = 10 и x - y = 2")

	x := solver.CreateIntVar("x")
	y := solver.CreateIntVar("y")

	ten := solver.CreateIntLit(10)
	two := solver.CreateIntLit(2)

	// x + y = 10
	constraint1 := x.Add(y).Eq(ten)
	// x - y = 2
	constraint2 := x.Sub(y).Eq(two)

	solver.Assert(constraint1)
	solver.Assert(constraint2)

	sat, err := solver.IsSatisfiable()
	if err != nil {
		log.Fatal(err)
	}

	if sat {
		model := solver.Model()

		xVal, _ := solver.GetIntValue(model, x)
		yVal, _ := solver.GetIntValue(model, y)

		fmt.Printf("Решение найдено: x = %d, y = %d\n", xVal, yVal)
		fmt.Printf("Проверка: %d + %d = %d, %d - %d = %d\n",
			xVal, yVal, xVal+yVal, xVal, yVal, xVal-yVal)
	} else {
		fmt.Println("Решение не найдено")
	}

	fmt.Println("\n=== Пример 2: Символьное исполнение условного оператора ===")

	// Новый solver для второй задачи
	solver2 := z3wrapper.NewSolver()
	defer solver2.Close()

	// Моделируем код:
	// if (a > 5) {
	//     b = a * 2;
	// } else {
	//     b = a + 10;
	// }
	// Найти значения a, при которых b = 20

	a := solver2.CreateIntVar("a")
	b := solver2.CreateIntVar("b")

	five := solver2.CreateIntLit(5)
	twenty := solver2.CreateIntLit(20)
	ten2 := solver2.CreateIntLit(10)
	two2 := solver2.CreateIntLit(2)

	// Условие: b = 20
	solver2.Assert(b.Eq(twenty))

	fmt.Println("Анализируем условие: if (a > 5) { b = a * 2 } else { b = a + 10 }")
	fmt.Println("Ищем значения a, при которых b = 20")

	// Путь 1: a > 5 && b = a * 2
	solver2.Push()
	solver2.Assert(a.GT(five))        // a > 5
	solver2.Assert(b.Eq(a.Mul(two2))) // b = a * 2

	sat1, _ := solver2.IsSatisfiable()
	if sat1 {
		model1 := solver2.Model()
		aVal1, _ := solver2.GetIntValue(model1, a)
		fmt.Printf("Путь 1 (a > 5): a = %d, b = %d * 2 = %d\n", aVal1, aVal1, aVal1*2)
	} else {
		fmt.Println("Путь 1 недостижим")
	}
	solver2.Pop()

	// Путь 2: a <= 5 && b = a + 10
	solver2.Push()
	solver2.Assert(a.LE(five))        // a <= 5
	solver2.Assert(b.Eq(a.Add(ten2))) // b = a + 10

	sat2, _ := solver2.IsSatisfiable()
	if sat2 {
		model2 := solver2.Model()
		aVal2, _ := solver2.GetIntValue(model2, a)
		fmt.Printf("Путь 2 (a <= 5): a = %d, b = %d + 10 = %d\n", aVal2, aVal2, aVal2+10)
	} else {
		fmt.Println("Путь 2 недостижим")
	}
	solver2.Pop()
}
