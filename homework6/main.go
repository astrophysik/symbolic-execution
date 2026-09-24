package main

import (
	"fmt"
	"symbolic-execution-course/internal"
)

func main() {
	source := `
package main

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
`

	result := internal.Analyse(source, "factorial")
	for _, interpreter := range result {
		fmt.Println(interpreter)
	}
}
