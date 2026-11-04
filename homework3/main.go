package main

import (
	"fmt"
	"os"

	"symbolic-execution-course/internal"
	"symbolic-execution-course/internal/translator"
)

func main() {
	source, err := os.ReadFile("homework3/examples/test_functions.go")
	if err != nil {
		panic(err)
	}

	results := internal.Analyse(string(source), "testMultipleReturns")

	zt := translator.NewZ3Translator()
	defer zt.Close()

	for i, state := range results {
		frame := state.CallStack[len(state.CallStack)-1]

		condition, err := zt.TranslateExpression(state.PathCondition)
		if err != nil {
			panic(err)
		}

		value, err := zt.TranslateExpression(frame.ReturnValue)
		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"Path %d\nConstraints: %s\nResult: %s\nZ3 conditions: %v\nZ3 result: %v\n\n",
			i+1,
			state.PathCondition,
			frame.ReturnValue,
			condition,
			value,
		)
	}
}