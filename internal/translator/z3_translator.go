// Package translator содержит реализацию транслятора в Z3
package translator

import (
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

// Z3Translator транслирует символьные выражения в Z3 формулы
type Z3Translator struct {
	ctx    *z3.Context
	config *z3.Config
	vars   map[string]z3.Value // Кэш переменных
}

func ExpressionToSort(zt *Z3Translator, typeInfo symbolic.NestedExpressionType) z3.Sort {
	switch typeInfo.ExprType {
	case symbolic.IntType:
		return zt.ctx.IntSort()
	case symbolic.BoolType:
		return zt.ctx.BoolSort()
	case symbolic.ArrayType:
		return zt.ctx.ArraySort(zt.ctx.IntSort(), ExpressionToSort(zt, *typeInfo.InnerType))
	default:
		panic("unknown symbolic expression type")
	}
}

// NewZ3Translator создаёт новый экземпляр Z3 транслятора
func NewZ3Translator() *Z3Translator {
	config := &z3.Config{}
	ctx := z3.NewContext(config)

	return &Z3Translator{
		ctx:    ctx,
		config: config,
		vars:   make(map[string]z3.Value),
	}
}

// GetContext возвращает Z3 контекст
func (zt *Z3Translator) GetContext() interface{} {
	return zt.ctx
}

// Reset сбрасывает состояние транслятора
func (zt *Z3Translator) Reset() {
	zt.vars = make(map[string]z3.Value)
}

// Close освобождает ресурсы
func (zt *Z3Translator) Close() {
	// Z3 контекст закрывается автоматически
}

// TranslateExpression транслирует символьное выражение в Z3
func (zt *Z3Translator) TranslateExpression(expr symbolic.SymbolicExpression) (interface{}, error) {
	return expr.Accept(zt), nil
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// VisitVariable транслирует символьную переменную в Z3
func (zt *Z3Translator) VisitVariable(expr *symbolic.SymbolicVariable) interface{} {
	if value, ok := zt.vars[expr.Name]; ok {
		return value
	}

	value := zt.ctx.Const(expr.Name, ExpressionToSort(zt, expr.TypeInfo))
	zt.vars[expr.Name] = value
	return value
}

// VisitIntConstant транслирует целочисленную константу в Z3
func (zt *Z3Translator) VisitIntConstant(expr *symbolic.IntConstant) interface{} {
	return zt.ctx.FromInt(expr.Value, zt.ctx.IntSort())
}

// VisitBoolConstant транслирует булеву константу в Z3
func (zt *Z3Translator) VisitBoolConstant(expr *symbolic.BoolConstant) interface{} {
	return zt.ctx.FromBool(expr.Value)
}

// VisitBinaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitBinaryOperation(expr *symbolic.BinaryOperation) interface{} {

	left := expr.Left.Accept(zt)
	right := expr.Right.Accept(zt)

	switch expr.Operator {
	case symbolic.ADD:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Add(right.(z3.Int))
		default:
			panic("unsupported type for add")
		}
	case symbolic.SUB:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Sub(right.(z3.Int))
		default:
			panic("unsupported type for sub")
		}
	case symbolic.MUL:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Mul(right.(z3.Int))
		default:
			panic("unsupported type for mul")
		}
	case symbolic.DIV:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Div(right.(z3.Int))
		default:
			panic("unsupported type for div")
		}
	case symbolic.EQ:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Eq(right.(z3.Int))
		case symbolic.BoolType:
			return left.(z3.Bool).Eq(right.(z3.Bool))
		case symbolic.ArrayType:
			return left.(z3.Array).Eq(right.(z3.Array))
		default:
			panic("unsupported type for eq")
		}
	case symbolic.NE:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).NE(right.(z3.Int))
		case symbolic.BoolType:
			return left.(z3.Bool).NE(right.(z3.Bool))
		case symbolic.ArrayType:
			return left.(z3.Array).NE(right.(z3.Array))
		default:
			panic("unsupported type for ne")
		}
	case symbolic.LT:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).LT(right.(z3.Int))
		default:
			panic("unsupported type for lt")
		}
	case symbolic.LE:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).LE(right.(z3.Int))
		default:
			panic("unsupported type for le")
		}
	case symbolic.GT:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).GT(right.(z3.Int))
		default:
			panic("unsupported type for gt")
		}
	case symbolic.GE:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).GE(right.(z3.Int))
		default:
			panic("unsupported type for ge")
		}
	default:
		panic("unsupported binary operation")
	}
}

// VisitLogicalOperation транслирует логическую операцию в Z3
func (zt *Z3Translator) VisitLogicalOperation(expr *symbolic.LogicalOperation) interface{} {
	z3Operands := make([]z3.Bool, len(expr.Operands))
	for i, value := range expr.Operands {
		z3Operands[i] = value.Accept(zt).(z3.Bool)
	}

	switch expr.Operator {
	case symbolic.AND:
		return z3Operands[0].And(z3Operands[1:]...)
	case symbolic.OR:
		return z3Operands[0].Or(z3Operands[1:]...)
	case symbolic.NOT:
		return z3Operands[0].Not()
	case symbolic.IMPLIES:
		return z3Operands[0].Implies(z3Operands[1])
	default:
		panic("unkonwn logical operation")
	}
}

// Вспомогательные методы

// createZ3Variable создаёт Z3 переменную соответствующего типа
func (zt *Z3Translator) createZ3Variable(name string, exprType symbolic.ExpressionType) z3.Value {
	exprSort := ExpressionToSort(zt, symbolic.NestedExpressionType{ExprType: exprType, InnerType: nil})
	return zt.ctx.Const(name, exprSort)
}

// castToZ3Type приводит значение к нужному Z3 типу
func (zt *Z3Translator) castToZ3Type(value interface{}, targetType symbolic.ExpressionType) (z3.Value, error) {
	switch targetType {
	case symbolic.IntType:
		return value.(z3.Int), nil
	case symbolic.BoolType:
		return value.(z3.Bool), nil
	case symbolic.ArrayType:
		return value.(z3.Array), nil
	default:
		panic("Unknown type in cast")
	}
}
