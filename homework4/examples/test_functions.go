package main

type Person struct {
	Name string
	Age  int
	ID   int
}

func testStructBasic() Person {
	var p Person
	p.Name = "Alice"
	p.Age = 25
	p.ID = 1001
	return p
}

func testStructModification(p Person) Person {
	p.Age = p.Age + 1
	p.ID = p.ID * 2
	return p
}

func testStructPointer() *Person {
	p := &Person{Name: "Bob", Age: 30, ID: 2002}
	p.Age = p.Age + 5
	return p
}

func testStructPointerModification(p *Person) {
	if p != nil {
		p.Age = p.Age + 10
		p.ID = p.ID + 1000
	}
}

func testNestedStructPointer() {
	p := testStructPointer()
	testStructPointerModification(p)
}

func testArrayFixed() [5]int {
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i * i
	}
	return arr
}

func testArrayModification(arr [5]int) [5]int {
	for i := range arr {
		arr[i] = arr[i] + 1
	}
	return arr
}

func testSliceCreation() []int {
	slice := make([]int, 5)
	for i := range slice {
		slice[i] = i * 2
	}
	return slice
}

func testSliceAppend() []int {
	var slice []int
	for i := 0; i < 3; i++ {
		slice = append(slice, i*10)
	}
	return slice
}

func testSliceModification(slice []int) []int {
	for i := range slice {
		slice[i] = slice[i] * 2
	}
	slice = append(slice, 999)
	return slice
}

type Student struct {
	Name    string
	Grades  [5]int
	Average float64
}

func testStructWithArray() Student {
	var s Student
	s.Name = "Charlie"
	s.Grades = [5]int{85, 90, 78, 92, 88}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	s.Average = float64(sum) / float64(len(s.Grades))

	return s
}

type Address struct {
	Street  string
	City    string
	ZipCode int
}

type Employee struct {
	Person  Person
	Address Address
	Salary  float64
}

func testNestedStructs() Employee {
	emp := Employee{
		Person: Person{
			Name: "David",
			Age:  35,
			ID:   3003,
		},
		Address: Address{
			Street:  "Main St",
			City:    "Boston",
			ZipCode: 12345,
		},
		Salary: 75000.0,
	}
	return emp
}

func testNestedStructModification(emp *Employee) {
	emp.Person.Age += 1
	emp.Salary *= 1.1
	emp.Address.ZipCode = 54321
}

func testArrayOfStructs() [3]Person {
	var people [3]Person
	people[0] = Person{Name: "Alice", Age: 25, ID: 1}
	people[1] = Person{Name: "Bob", Age: 30, ID: 2}
	people[2] = Person{Name: "Charlie", Age: 35, ID: 3}

	people[1].Age += 5

	return people
}

func testArrayOfStructsModification(people *[3]Person) {
	for i := range people {
		people[i].ID = people[i].ID * 10
	}
}

type Foo struct {
	a int
}

func Aliasing(foo1 *Foo, foo2 *Foo) int {
	foo2.a = 5
	foo1.a = 2
	if foo2.a == 2 {
		return 4
	}
	return 5
}

func testPathConstraintMutability(p Person) {
	if p.Age != 18 {
		p.Age = 18
		if p.Age != 18 {
			panic("Seems impossible")
		} else {
			println("Seems ok")
		}
	}
}
