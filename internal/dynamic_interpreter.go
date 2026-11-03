package internal

import (
	"fmt"
	"go/constant"
	
	"golang.org/x/tools/go/ssa"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	Heap          memory.Memory
	Finished      bool
}

type CallStackFrame struct {
	Function      *ssa.Function
	LocalMemory   map[string]symbolic.SymbolicExpression
	ReturnValue   symbolic.SymbolicExpression
	CurrentBlock  *ssa.BasicBlock
	PreviousBlock *ssa.BasicBlock
	InstrIdx      int
}

func (interpreter *Interpreter) interpretDynamically(element ssa.Instruction) []Interpreter {
	frame := &interpreter.CallStack[len(interpreter.CallStack)-1]

	switch instruction := element.(type) {
	case *ssa.Return:
		if len(instruction.Results) > 1 {
			panic("multiple return values are not supported yet")
		}

		if len(instruction.Results) == 1 {
			frame.ReturnValue =
				interpreter.resolveExpression(instruction.Results[0])
		}

		frame.InstrIdx++
		interpreter.Finished = true
		return []Interpreter{*interpreter}
	default:
		panic(fmt.Sprintf("unsupported SSA instruction: %T", element))
	}
}

func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {
	switch value := value.(type) {
	case *ssa.Parameter:
		frame := &interpreter.CallStack[len(interpreter.CallStack)-1]

		expression, ok := frame.LocalMemory[value.Name()]
		if !ok {
			panic(fmt.Sprintf("name %q not found", value.Name()))
		}
		return expression
	case *ssa.Const:
		if value.Value == nil {
			panic("nil constants are not supported yet")
		}

		switch value.Value.Kind() {
		case constant.Int:
			number, exact := constant.Int64Val(value.Value)
			if !exact {
				panic(fmt.Sprintf("integer constant does not fit int64: %s", value))
			}
			return symbolic.NewIntConstant(number)
		case constant.Bool:
			return symbolic.NewBoolConstant(constant.BoolVal(value.Value))
		default:
			panic(fmt.Sprintf("unsupported constant: %s", value))
		}
	default:
		panic(fmt.Sprintf("unsupported SSA value: %T", value))
	}
}
