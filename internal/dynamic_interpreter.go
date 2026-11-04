package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"maps"

	"golang.org/x/tools/go/ssa"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

var binaryOperators = map[token.Token]symbolic.BinaryOperator{
	token.ADD: symbolic.ADD,
	token.SUB: symbolic.SUB,
	token.MUL: symbolic.MUL,
	token.QUO: symbolic.DIV,
	token.REM: symbolic.MOD,
	token.EQL: symbolic.EQ,
	token.NEQ: symbolic.NE,
	token.LSS: symbolic.LT,
	token.LEQ: symbolic.LE,
	token.GTR: symbolic.GT,
	token.GEQ: symbolic.GE,
	token.AND: symbolic.BIT_AND,
	token.OR:  symbolic.BIT_OR,
	token.XOR: symbolic.BIT_XOR,
	token.SHL: symbolic.SHL,
	token.SHR: symbolic.SHR,
}

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

func (interpreter *Interpreter) clone() Interpreter {
	cloned := *interpreter

	cloned.CallStack = make([]CallStackFrame, len(interpreter.CallStack))
	copy(cloned.CallStack, interpreter.CallStack)

	for i := range cloned.CallStack {
		cloned.CallStack[i].LocalMemory =
			maps.Clone(interpreter.CallStack[i].LocalMemory)
	}

	return cloned
}

func (interpreter *Interpreter) interpretDynamically(element ssa.Instruction) []Interpreter {
	frame := &interpreter.CallStack[len(interpreter.CallStack)-1]

	switch instruction := element.(type) {
	case *ssa.BinOp:
		operator, ok := binaryOperators[instruction.Op]
		if !ok {
			panic(fmt.Sprintf(
				"unsupported binary operator: %s", instruction.Op,
			))
		}
		frame.LocalMemory[instruction.Name()] = symbolic.NewBinaryOperation(
			interpreter.resolveExpression(instruction.X),
			interpreter.resolveExpression(instruction.Y),
			operator,
		)

		frame.InstrIdx++
		return []Interpreter{*interpreter}
	case *ssa.UnOp:
		operand := interpreter.resolveExpression(instruction.X)

		var result symbolic.SymbolicExpression
		switch instruction.Op {
		case token.SUB:
			result = symbolic.NewBinaryOperation(symbolic.NewIntConstant(0), operand, symbolic.SUB)
		case token.NOT:
			result = symbolic.NewLogicalOperation(
				[]symbolic.SymbolicExpression{operand},
				symbolic.NOT,
			)
		case token.XOR:
			result = symbolic.NewBinaryOperation(
				symbolic.NewIntConstant(-1),
				operand,
				symbolic.SUB,
			)

		default:
			panic(fmt.Sprintf(
				"unsupported unary operator: %s", instruction.Op,
			))
		}

		frame.LocalMemory[instruction.Name()] = result
		frame.InstrIdx++
		return []Interpreter{*interpreter}
	case *ssa.If:
		condition := interpreter.resolveExpression(instruction.Cond)
		notCondition := symbolic.NewLogicalOperation(
			[]symbolic.SymbolicExpression{condition},
			symbolic.NOT,
		)

		trueState := interpreter.clone()
		falseState := interpreter.clone()
		trueState.PathCondition = symbolic.NewLogicalOperation(
			[]symbolic.SymbolicExpression{
				interpreter.PathCondition, condition,
			},
			symbolic.AND,
		)
		falseState.PathCondition = symbolic.NewLogicalOperation(
			[]symbolic.SymbolicExpression{
				interpreter.PathCondition, notCondition,
			},
			symbolic.AND,
		)

		trueFrame := &trueState.CallStack[len(trueState.CallStack)-1]
		trueFrame.PreviousBlock = frame.CurrentBlock
		trueFrame.CurrentBlock = frame.CurrentBlock.Succs[0]
		trueFrame.InstrIdx = 0

		falseFrame := &falseState.CallStack[len(falseState.CallStack)-1]
		falseFrame.PreviousBlock = frame.CurrentBlock
		falseFrame.CurrentBlock = frame.CurrentBlock.Succs[1]
		falseFrame.InstrIdx = 0

		return []Interpreter{trueState, falseState}
	case *ssa.Phi:
		predIndex := -1
		for i, predecessor := range frame.CurrentBlock.Preds {
			if predecessor == frame.PreviousBlock {
				predIndex = i
				break
			}
		}

		if predIndex == -1 {
			panic("previous block not found among predecessors")
		}

		values := make(map[string]symbolic.SymbolicExpression)

		for frame.InstrIdx < len(frame.CurrentBlock.Instrs) {
			phi, ok := frame.CurrentBlock.Instrs[frame.InstrIdx].(*ssa.Phi)
			if !ok {
				break
			}

			values[phi.Name()] =
				interpreter.resolveExpression(phi.Edges[predIndex])
			frame.InstrIdx++
		}

		maps.Copy(frame.LocalMemory, values)
		return []Interpreter{*interpreter}
	case *ssa.Jump:
		frame.PreviousBlock = frame.CurrentBlock
		frame.CurrentBlock = frame.CurrentBlock.Succs[0]
		frame.InstrIdx = 0

		return []Interpreter{*interpreter}
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
	case *ssa.Parameter, *ssa.BinOp, *ssa.UnOp, *ssa.Phi:
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

		basic, ok := value.Type().Underlying().(*types.Basic)
		if !ok {
			panic(fmt.Sprintf("unsupported constant type: %s", value.Type()))
		}

		switch {
		case basic.Info()&types.IsInteger != 0:
			number, exact := constant.Int64Val(
				constant.ToInt(value.Value),
			)
			if !exact {
				panic(fmt.Sprintf(
					"integer constant does not fit int64: %s", value,
				))
			}
			return symbolic.NewIntConstant(number)

		case basic.Info()&types.IsBoolean != 0:
			return symbolic.NewBoolConstant(
				constant.BoolVal(value.Value),
			)

		case basic.Kind() == types.Float64:
			number, _ := constant.Float64Val(
				constant.ToFloat(value.Value),
			)
			return symbolic.NewFloatConstant(number)

		default:
			panic(fmt.Sprintf("unsupported constant: %s", value))
		}
	default:
		panic(fmt.Sprintf("unsupported SSA value: %T", value))
	}
}
