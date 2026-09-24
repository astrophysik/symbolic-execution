package z3wrapper

import (
	"testing"
)

func TestSolverBasicOperations(t *testing.T) {
	solver := NewSolver()
	defer solver.Close()

	// Создаём переменные
	x := solver.CreateIntVar("x")
	y := solver.CreateIntVar("y")

	// Создаём константы
	five := solver.CreateIntLit(5)
	ten := solver.CreateIntLit(10)

	// Добавляем ограничения: x + y = 10, x = 5
	solver.Assert(x.Add(y).Eq(ten))
	solver.Assert(x.Eq(five))

	// Проверяем выполнимость
	sat, err := solver.IsSatisfiable()
	if err != nil {
		t.Fatalf("Error checking satisfiability: %v", err)
	}

	if !sat {
		t.Fatal("Expected satisfiable constraints")
	}

	// Получаем модель
	model := solver.Model()
	xVal, err := solver.GetIntValue(model, x)
	if err != nil {
		t.Fatalf("Error getting x value: %v", err)
	}

	yVal, err := solver.GetIntValue(model, y)
	if err != nil {
		t.Fatalf("Error getting y value: %v", err)
	}

	// Проверяем результаты
	if xVal != 5 {
		t.Errorf("Expected x = 5, got %d", xVal)
	}
	if yVal != 5 {
		t.Errorf("Expected y = 5, got %d", yVal)
	}
}

func TestSolverUnsatisfiable(t *testing.T) {
	solver := NewSolver()
	defer solver.Close()

	x := solver.CreateIntVar("x")
	zero := solver.CreateIntLit(0)
	one := solver.CreateIntLit(1)

	// Несовместимые ограничения: x = 0 и x = 1
	solver.Assert(x.Eq(zero))
	solver.Assert(x.Eq(one))

	sat, err := solver.IsSatisfiable()
	if err != nil {
		t.Fatalf("Error checking satisfiability: %v", err)
	}

	if sat {
		t.Fatal("Expected unsatisfiable constraints")
	}
}

func TestSolverPushPop(t *testing.T) {
	solver := NewSolver()
	defer solver.Close()

	x := solver.CreateIntVar("x")
	five := solver.CreateIntLit(5)
	ten := solver.CreateIntLit(10)

	// Базовое ограничение
	solver.Assert(x.Eq(five))

	// Проверяем, что x = 5 выполнимо
	sat1, _ := solver.IsSatisfiable()
	if !sat1 {
		t.Fatal("Expected x = 5 to be satisfiable")
	}

	// Добавляем противоречащее ограничение
	solver.Push()
	solver.Assert(x.Eq(ten))

	// Проверяем, что теперь невыполнимо
	sat2, _ := solver.IsSatisfiable()
	if sat2 {
		t.Fatal("Expected x = 5 AND x = 10 to be unsatisfiable")
	}

	// Возвращаемся к предыдущему состоянию
	solver.Pop()

	// Проверяем, что снова выполнимо
	sat3, _ := solver.IsSatisfiable()
	if !sat3 {
		t.Fatal("Expected x = 5 to be satisfiable after pop")
	}
}

func TestBooleanOperations(t *testing.T) {
	solver := NewSolver()
	defer solver.Close()

	a := solver.CreateBoolVar("a")
	b := solver.CreateBoolVar("b")

	// a AND (NOT b) должно быть выполнимо
	solver.Assert(a)
	solver.Assert(b.Not())

	sat, err := solver.IsSatisfiable()
	if err != nil {
		t.Fatalf("Error checking satisfiability: %v", err)
	}

	if !sat {
		t.Fatal("Expected a AND (NOT b) to be satisfiable")
	}

	model := solver.Model()
	aVal, err := solver.GetBoolValue(model, a)
	if err != nil {
		t.Fatalf("Error getting a value: %v", err)
	}

	bVal, err := solver.GetBoolValue(model, b)
	if err != nil {
		t.Fatalf("Error getting b value: %v", err)
	}

	if !aVal {
		t.Errorf("Expected a = true, got %v", aVal)
	}
	if bVal {
		t.Errorf("Expected b = false, got %v", bVal)
	}
}
