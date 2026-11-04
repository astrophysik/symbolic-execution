// Демонстрационная программа для тестирования SSA построения
package main

import (
	"fmt"
	"log"
	

	ssainternal "symbolic-execution-course/internal/ssa"
	"golang.org/x/tools/go/ssa"
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
	builder := ssainternal.NewBuilder()

	// Строим SSA из исходного кода
	graph, err := builder.ParseAndBuildSSA(source, "testFunction")
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}

	for _, basicBlock := range graph.DomPreorder() {
		fmt.Printf("Basic Block : %d\n", basicBlock.Index)
		for _, instruction := range basicBlock.Instrs {
			fmt.Println(instruction.String())
			value, ok := instruction.(ssa.Value)
			if ok == true {
				fmt.Printf("	with value %s\n", value.Name())
				
			}
		}
	}


	graph.DomPreorder()
	fmt.Printf("CFG построен для функции с %d блоками\n", len(graph.Blocks))
}
