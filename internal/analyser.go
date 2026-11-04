package internal

import (
	"container/heap"
	"fmt"
	"go/types"
	ssabuilder "symbolic-execution-course/internal/ssa"
	"symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"

	"golang.org/x/tools/go/ssa"
)

type Analyser struct {
	Package      *ssa.Package
	StatesQueue  PriorityQueue
	PathSelector PathSelector
	Results      []Interpreter
	Z3Translator *translator.Z3Translator
	MaxSteps     int
	StepCount    int
}

func toSymbolicType(t types.Type) symbolic.NestedExpressionType {
	switch t := t.Underlying().(type) {
	case *types.Basic:
		switch {
		case t.Info()&types.IsBoolean != 0:
			return symbolic.NestedExpressionType{
				ExprType: symbolic.BoolType,
			}
		case t.Info()&types.IsInteger != 0:
			return symbolic.NestedExpressionType{
				ExprType: symbolic.IntType,
			}
		case t.Kind() == types.Float64:
			return symbolic.NestedExpressionType{
				ExprType: symbolic.FloatType,
			}
		default:
			panic(fmt.Sprintf("unsupported basic type: %s", t))
		}
	case *types.Slice:
		elementType := toSymbolicType(t.Elem())
		return symbolic.NestedExpressionType{
			ExprType:  symbolic.ArrayType,
			InnerType: &elementType,
		}
	default:
		panic(fmt.Sprintf("unsupported type: %s", t))
	}
}

func Analyse(source string, functionName string) []Interpreter {
	builder := ssabuilder.NewBuilder()
	functionSsa, err := builder.ParseAndBuildSSA(source, functionName)
	if err != nil {
		panic("error inside Analyse")
	}
	if functionSsa == nil {
		panic(fmt.Sprintf("function %q not found", functionName))
	}
	if len(functionSsa.Blocks) == 0 {
		panic(fmt.Sprintf("function %q has no body", functionName))
	}

	translator := translator.NewZ3Translator()
	defer translator.Close()
	analyser := Analyser{
		Package:      functionSsa.Pkg,
		StatesQueue:  PriorityQueue{},
		PathSelector: &BfsPathSelector{},
		Results:      nil,
		Z3Translator: translator,
		MaxSteps:     30,
	}

	heap.Init(&analyser.StatesQueue)
	initialState := Interpreter{
		Analyser:      &analyser,
		PathCondition: symbolic.NewBoolConstant(true),
		Heap:          nil,
		CallStack: []CallStackFrame{
			{Function: functionSsa,
				LocalMemory:  make(map[string]symbolic.SymbolicExpression),
				CurrentBlock: functionSsa.Blocks[0]},
		},
	}

	frame := &initialState.CallStack[0]
	for _, param := range frame.Function.Params {
		frame.LocalMemory[param.Name()] = &symbolic.SymbolicVariable{
			Name:     param.Name(),
			TypeInfo: toSymbolicType(param.Type()),
		}
	}

	heap.Push(&analyser.StatesQueue, &Item{
		value:    initialState,
		priority: analyser.PathSelector.CalculatePriority(initialState),
	})

	analyser.runAnalysis()
	return analyser.Results
}

func (a *Analyser) runAnalysis() {
	for a.StatesQueue.Len() > 0 && a.StepCount < a.MaxSteps {
		item := heap.Pop(&a.StatesQueue).(*Item)
		state := item.value

		if len(state.CallStack) == 0 {
			panic("active state has an empty call stack")
		}

		frame := &state.CallStack[len(state.CallStack)-1]

		if frame.CurrentBlock == nil {
			panic("active state has no current block")
		}

		if frame.InstrIdx < 0 || frame.InstrIdx >= len(frame.CurrentBlock.Instrs) {
			panic("instruction index out of bounds")
		}

		instruction := frame.CurrentBlock.Instrs[frame.InstrIdx]
		a.StepCount++
		nextStates := state.interpretDynamically(instruction)

		for _, nextState := range nextStates {
			if nextState.Finished {
				a.Results = append(a.Results, nextState)
				continue
			}

			heap.Push(&a.StatesQueue, &Item{
				value:    nextState,
				priority: a.PathSelector.CalculatePriority(nextState),
			})

		}
	}

	if a.StatesQueue.Len() > 0 {
		fmt.Printf(
			"Analys was stopped after %d steps. Left states %d\n",
			a.StepCount,
			a.StatesQueue.Len(),
		)
	}

}
