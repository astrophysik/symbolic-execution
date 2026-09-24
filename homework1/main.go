// Демонстрационная программа для тестирования SSA построения
package main

import (
	"fmt"
	"log"

	"symbolic-execution-course/internal/ssa"
)

func main() {
	fmt.Println("=== SSA Builder Demo ===")

	// Пример исходного кода для анализа
	source := `
package main

func testFunction(x int) int {
	if x > 0 {
		return x * 2
	} else {
		return x * -1
	}
}
`

	// Создаём builder для SSA
	builder := ssa.NewBuilder()

	// Строим SSA из исходного кода
	graph, err := builder.ParseAndBuildSSA(source, "testFunction")
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}
	fmt.Printf("CFG построен для функции с %d блоками\n", len(graph.Blocks))
}
