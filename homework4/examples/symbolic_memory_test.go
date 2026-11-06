package main

import (
	"github.com/ebukreev/go-z3/z3"
	"symbolic-execution-course/internal/memory"
	s "symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"
	"testing"
)

// Здесь тесты из test_functions.go  

func check(t *testing.T, actual, expected s.SymbolicExpression) {
	t.Helper()

	tr := translator.NewZ3Translator()
	defer tr.Close()

	v, err := tr.TranslateExpression(s.NewBinaryOperation(actual, expected, s.NE))
	if err != nil {
		t.Fatal(err)
	}

	solver := z3.NewSolver(tr.GetContext().(*z3.Context))
	solver.Assert(v.(z3.Bool))
	sat, err := solver.Check()
	if err != nil {
		t.Fatal(err)
	}
	if sat {
		t.Fatalf("got %s, expected %s", actual, expected)
	}
}

func to_z3_int(v int) s.SymbolicExpression {
	return s.NewIntConstant(int64(v))
}

func initPerson(m *memory.SymbolicMemory, p Person) *s.Ref {
	ref := m.AllocateStruct(s.StringType, s.IntType, s.IntType)
	m.AssignField(ref, 0, s.NewStringConstant(p.Name))
	m.AssignField(ref, 1, to_z3_int(p.Age))
	m.AssignField(ref, 2, to_z3_int(p.ID))
	return ref
}

func assertPerson(t *testing.T, m *memory.SymbolicMemory, ref *s.Ref, want Person) {
	t.Helper()
	check(t, m.GetFieldValue(ref, 0), s.NewStringConstant(want.Name))
	check(t, m.GetFieldValue(ref, 1), to_z3_int(want.Age))
	check(t, m.GetFieldValue(ref, 2), to_z3_int(want.ID))
}

func modifyPerson(m *memory.SymbolicMemory, p *s.Ref, age, id int) {
	m.AssignField(p, 1, s.NewBinaryOperation(m.GetFieldValue(p, 1), to_z3_int(age), s.ADD))
	m.AssignField(p, 2, s.NewBinaryOperation(m.GetFieldValue(p, 2), to_z3_int(id), s.ADD))
}

func TestStructExamples(t *testing.T) {
	t.Run("testStructBasic", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		p := initPerson(&m, Person{"Alice", 25, 1001})
		assertPerson(t, &m, p, testStructBasic())
	})

	t.Run("testStructModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		input := Person{"Anna", 19, 42}
		p := initPerson(&m, input)
		m.AssignField(p, 1, s.NewBinaryOperation(m.GetFieldValue(p, 1), to_z3_int(1), s.ADD))
		m.AssignField(p, 2, s.NewBinaryOperation(m.GetFieldValue(p, 2), to_z3_int(2), s.MUL))
		assertPerson(t, &m, p, testStructModification(input))
	})

	t.Run("testStructPointer", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		p := initPerson(&m, Person{"Bob", 30, 2002})
		m.AssignField(p, 1, s.NewBinaryOperation(m.GetFieldValue(p, 1), to_z3_int(5), s.ADD))
		want := testStructPointer()
		assertPerson(t, &m, p, *want)
	})

	t.Run("testStructPointerModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		want := Person{"Bob", 30, 2002}
		p := initPerson(&m, want)
		alias := s.NewStructRef(p.Address, s.StringType, s.IntType, s.IntType)
		m.AssignField(alias, 1, s.NewBinaryOperation(m.GetFieldValue(alias, 1), to_z3_int(10), s.ADD))
		m.AssignField(alias, 2, s.NewBinaryOperation(m.GetFieldValue(alias, 2), to_z3_int(1000), s.ADD))
		testStructPointerModification(&want)
		assertPerson(t, &m, p, want)
	})

	t.Run("testNestedStructPointer", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		p := initPerson(&m, Person{"Bob", 30, 2002})
		modifyPerson(&m, p, 5, 0)
		alias := s.NewStructRef(p.Address, s.StringType, s.IntType, s.IntType)
		modifyPerson(&m, alias, 10, 1000)
		testNestedStructPointer()
		assertPerson(t, &m, p, Person{"Bob", 45, 3002})
	})
}

func TestArrayAndSliceExamples(t *testing.T) {
	t.Run("testArrayFixed", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.IntType)
		for i := 0; i < 5; i++ {
			m.AssignToArray(arr, i, to_z3_int(i*i))
		}
		want := testArrayFixed()
		for i, v := range want {
			check(t, m.GetFromArray(arr, i), to_z3_int(v))
		}
	})

	t.Run("testArrayModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.IntType)
		input := [5]int{0, 1, 4, 9, 16}
		for i, v := range input {
			m.AssignToArray(arr, i, to_z3_int(v))
			m.AssignToArray(arr, i, s.NewBinaryOperation(m.GetFromArray(arr, i), to_z3_int(1), s.ADD))
		}
		want := testArrayModification(input)
		for i, v := range want {
			check(t, m.GetFromArray(arr, i), to_z3_int(v))
		}
	})

	t.Run("testSliceCreation", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.IntType)
		for i := 0; i < 5; i++ {
			m.AssignToArray(arr, i, to_z3_int(i*2))
		}
		want := testSliceCreation()
		for i, v := range want {
			check(t, m.GetFromArray(arr, i), to_z3_int(v))
		}
	})

	t.Run("testSliceAppend", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.IntType)
		for i := 0; i < 3; i++ {
			m.AssignToArray(arr, i, to_z3_int(i*10))
		}
		want := testSliceAppend()
		for i, v := range want {
			check(t, m.GetFromArray(arr, i), to_z3_int(v))
		}
	})

	t.Run("testSliceModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.IntType)
		input := []int{3, 7, 11}
		for i, v := range input {
			m.AssignToArray(arr, i, to_z3_int(v))
			m.AssignToArray(arr, i, s.NewBinaryOperation(m.GetFromArray(arr, i), to_z3_int(2), s.MUL))
		}
		m.AssignToArray(arr, len(input), to_z3_int(999))
		want := testSliceModification(input)
		for i, v := range want {
			check(t, m.GetFromArray(arr, i), to_z3_int(v))
		}
	})
}

func TestStructWithArrayExample(t *testing.T) {
	t.Run("testStructWithArray", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		student := m.AllocateStruct(s.StringType, s.RefType, s.FloatType)
		grades := m.AllocateArray(s.IntType)
		m.AssignField(student, 0, s.NewStringConstant("Charlie"))
		m.AssignField(student, 1, grades)
		var sum s.SymbolicExpression = to_z3_int(0)
		for i, g := range []int{85, 90, 78, 92, 88} {
			m.AssignToArray(grades, i, to_z3_int(g))
			sum = s.NewBinaryOperation(sum, m.GetFromArray(grades, i), s.ADD)
		}
		check(t, sum, to_z3_int(433))
		m.AssignField(student, 2, s.NewFloatConstant(433.0/5))
		want := testStructWithArray()
		check(t, m.GetFieldValue(student, 0), s.NewStringConstant(want.Name))
		check(t, m.GetFieldValue(student, 2), s.NewFloatConstant(want.Average))
		nested := s.NewArrayRef(m.GetFieldValue(student, 1), s.IntType)
		for i, g := range want.Grades {
			check(t, m.GetFromArray(nested, i), to_z3_int(g))
		}
	})
}

func TestNestedStructExamples(t *testing.T) {
	t.Run("testNestedStructs", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		p := initPerson(&m, Person{"David", 35, 3003})
		addr := m.AllocateStruct(s.StringType, s.StringType, s.IntType)
		emp := m.AllocateStruct(s.RefType, s.RefType, s.FloatType)
		m.AssignField(addr, 0, s.NewStringConstant("Main St"))
		m.AssignField(addr, 1, s.NewStringConstant("Boston"))
		m.AssignField(addr, 2, to_z3_int(12345))
		m.AssignField(emp, 0, p)
		m.AssignField(emp, 1, addr)
		m.AssignField(emp, 2, s.NewFloatConstant(75000))
		want := testNestedStructs()
		nestedP := s.NewStructRef(m.GetFieldValue(emp, 0), s.StringType, s.IntType, s.IntType)
		nestedA := s.NewStructRef(m.GetFieldValue(emp, 1), s.StringType, s.StringType, s.IntType)
		assertPerson(t, &m, nestedP, want.Person)
		check(t, m.GetFieldValue(nestedA, 0), s.NewStringConstant(want.Address.Street))
		check(t, m.GetFieldValue(nestedA, 1), s.NewStringConstant(want.Address.City))
		check(t, m.GetFieldValue(nestedA, 2), to_z3_int(want.Address.ZipCode))
		check(t, m.GetFieldValue(emp, 2), s.NewFloatConstant(want.Salary))
	})

	t.Run("testNestedStructModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		p := initPerson(&m, Person{"David", 35, 3003})
		addr := m.AllocateStruct(s.StringType, s.StringType, s.IntType)
		emp := m.AllocateStruct(s.RefType, s.RefType, s.FloatType)
		m.AssignField(addr, 0, s.NewStringConstant("Main St"))
		m.AssignField(addr, 1, s.NewStringConstant("Boston"))
		m.AssignField(addr, 2, to_z3_int(12345))
		m.AssignField(emp, 0, p)
		m.AssignField(emp, 1, addr)
		m.AssignField(emp, 2, s.NewFloatConstant(75000))
		want := Employee{Person: Person{"David", 35, 3003}, Address: Address{"Main St", "Boston", 12345}, Salary: 75000}
		nestedP := s.NewStructRef(m.GetFieldValue(emp, 0), s.StringType, s.IntType, s.IntType)
		nestedA := s.NewStructRef(m.GetFieldValue(emp, 1), s.StringType, s.StringType, s.IntType)
		modifyPerson(&m, nestedP, 1, 0)
		m.AssignField(emp, 2, s.NewBinaryOperation(m.GetFieldValue(emp, 2), s.NewFloatConstant(1.1), s.MUL))
		m.AssignField(nestedA, 2, to_z3_int(54321))
		testNestedStructModification(&want)
		assertPerson(t, &m, nestedP, want.Person)
		check(t, m.GetFieldValue(nestedA, 0), s.NewStringConstant(want.Address.Street))
		check(t, m.GetFieldValue(nestedA, 1), s.NewStringConstant(want.Address.City))
		check(t, m.GetFieldValue(nestedA, 2), to_z3_int(want.Address.ZipCode))
		check(t, m.GetFieldValue(emp, 2), s.NewFloatConstant(want.Salary))
	})
}

func TestArrayOfStructsExamples(t *testing.T) {
	t.Run("testArrayOfStructs", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.RefType)
		for i, p := range []Person{{"Alice", 25, 1}, {"Bob", 30, 2}, {"Charlie", 35, 3}} {
			m.AssignToArray(arr, i, initPerson(&m, p))
		}
		bob := s.NewStructRef(m.GetFromArray(arr, 1), s.StringType, s.IntType, s.IntType)
		modifyPerson(&m, bob, 5, 0)
		want := testArrayOfStructs()
		for i, p := range want {
			ref := s.NewStructRef(m.GetFromArray(arr, i), s.StringType, s.IntType, s.IntType)
			assertPerson(t, &m, ref, p)
		}
	})

	t.Run("testArrayOfStructsModification", func(t *testing.T) {
		m := memory.NewSymbolicMemory()
		arr := m.AllocateArray(s.RefType)
		for i, p := range []Person{{"Alice", 25, 1}, {"Bob", 30, 2}, {"Charlie", 35, 3}} {
			m.AssignToArray(arr, i, initPerson(&m, p))
		}
		bob := s.NewStructRef(m.GetFromArray(arr, 1), s.StringType, s.IntType, s.IntType)
		modifyPerson(&m, bob, 5, 0)
		want := [3]Person{{"Alice", 25, 1}, {"Bob", 35, 2}, {"Charlie", 35, 3}}
		for i := 0; i < 3; i++ {
			p := s.NewStructRef(m.GetFromArray(arr, i), s.StringType, s.IntType, s.IntType)
			m.AssignField(p, 2, s.NewBinaryOperation(m.GetFieldValue(p, 2), to_z3_int(10), s.MUL))
		}
		testArrayOfStructsModification(&want)
		for i, p := range want {
			ref := s.NewStructRef(m.GetFromArray(arr, i), s.StringType, s.IntType, s.IntType)
			assertPerson(t, &m, ref, p)
		}
	})
}
