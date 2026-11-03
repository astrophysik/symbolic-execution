package main

import (
	"fmt"
	"symbolic-execution-course/internal"
)

func main() {
	source := `
package main

func testFunction(x int) int {
	return x
}
`

	result := internal.Analyse(source, "testFunction")
	for _, interpreter := range result {
		fmt.Println(interpreter)
	}
}
