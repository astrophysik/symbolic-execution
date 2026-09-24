package internal

import (
	"golang.org/x/tools/go/ssa"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	Heap          memory.Memory
}

type CallStackFrame struct {
	Function    *ssa.Function
	LocalMemory map[string]symbolic.SymbolicExpression
	ReturnValue symbolic.SymbolicExpression
}

func (interpreter *Interpreter) interpretDynamically(element ssa.Instruction) []Interpreter {
	switch element.(type) {
	// TODO implement me
	}
	panic("implement me")
}

func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {
	switch value.(type) {
	// TODO implement me
	}
	panic("implement me")
}
