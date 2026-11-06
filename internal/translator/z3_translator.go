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
	case symbolic.IntType, symbolic.RefType:
		return zt.ctx.IntSort()
	case symbolic.StringType:
		return zt.ctx.ArraySort(zt.ctx.IntSort(), zt.ctx.IntSort())
	case symbolic.BoolType:
		return zt.ctx.BoolSort()
	case symbolic.ArrayType:
		return zt.ctx.ArraySort(zt.ctx.IntSort(), ExpressionToSort(zt, *typeInfo.InnerType))
	case symbolic.FloatType:
		return zt.ctx.FloatSort(11, 53)
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

func (zt *Z3Translator) VisitFloatConstant(expr *symbolic.FloatConstant) interface{} {
	return zt.ctx.FromFloat64(expr.Value, zt.ctx.FloatSort(11, 53))
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
		case symbolic.FloatType:
			return left.(z3.Float).Add(right.(z3.Float))
		default:
			panic("unsupported type for add")
		}
	case symbolic.SUB:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Sub(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).Sub(right.(z3.Float))
		default:
			panic("unsupported type for sub")
		}
	case symbolic.MUL:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Mul(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).Mul(right.(z3.Float))
		default:
			panic("unsupported type for mul")
		}
	case symbolic.DIV:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Div(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).Div(right.(z3.Float))
		default:
			panic("unsupported type for div")
		}
	case symbolic.EQ:
		switch expr.Left.Type() {
		case symbolic.IntType, symbolic.RefType:
			return left.(z3.Int).Eq(right.(z3.Int))
		case symbolic.BoolType:
			return left.(z3.Bool).Eq(right.(z3.Bool))
		case symbolic.ArrayType, symbolic.StringType:
			return left.(z3.Array).Eq(right.(z3.Array))
		case symbolic.FloatType:
			return left.(z3.Float).IEEEEq(right.(z3.Float))
		default:
			panic("unsupported type for eq")
		}
	case symbolic.NE:
		switch expr.Left.Type() {
		case symbolic.IntType, symbolic.RefType:
			return left.(z3.Int).NE(right.(z3.Int))
		case symbolic.BoolType:
			return left.(z3.Bool).NE(right.(z3.Bool))
		case symbolic.ArrayType, symbolic.StringType:
			return left.(z3.Array).NE(right.(z3.Array))
		case symbolic.FloatType:
			return left.(z3.Float).IEEEEq(right.(z3.Float)).Not()
		default:
			panic("unsupported type for ne")
		}
	case symbolic.LT:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).LT(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).LT(right.(z3.Float))
		default:
			panic("unsupported type for lt")
		}
	case symbolic.LE:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).LE(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).LE(right.(z3.Float))
		default:
			panic("unsupported type for le")
		}
	case symbolic.GT:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).GT(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).GT(right.(z3.Float))
		default:
			panic("unsupported type for gt")
		}
	case symbolic.GE:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).GE(right.(z3.Int))
		case symbolic.FloatType:
			return left.(z3.Float).GE(right.(z3.Float))
		default:
			panic("unsupported type for ge")
		}
	case symbolic.MOD:
		switch expr.Left.Type() {
		case symbolic.IntType:
			return left.(z3.Int).Mod(right.(z3.Int))
		default:
			panic("unsupported type for mod")
		}
	case symbolic.BIT_AND, symbolic.BIT_OR, symbolic.BIT_XOR,
		symbolic.SHL, symbolic.SHR:

		if expr.Left.Type() != symbolic.IntType ||
			expr.Right.Type() != symbolic.IntType {
			panic("bitwise operations require integer operands")
		}

		const bitWidth = 64
		l := left.(z3.Int).ToBV(bitWidth)
		r := right.(z3.Int).ToBV(bitWidth)

		var result z3.BV
		switch expr.Operator {
		case symbolic.BIT_AND:
			result = l.And(r)
		case symbolic.BIT_OR:
			result = l.Or(r)
		case symbolic.BIT_XOR:
			result = l.Xor(r)
		case symbolic.SHL:
			result = l.Lsh(r)
		case symbolic.SHR:
			result = l.SRsh(r)
		}

		return result.SToInt()
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

func (zt *Z3Translator) VisitRef(expr *symbolic.Ref) interface{} {
	return expr.Address.Accept(zt)
}

func (zt *Z3Translator) VisitStringConstant(expr *symbolic.StringConstant) interface{} {
	to_z3_integer := func(n int64) z3.Value { return zt.ctx.FromInt(n, zt.ctx.IntSort()) }

	result := zt.ctx.ConstArray(zt.ctx.IntSort(), to_z3_integer(0))

	// str[-1] = strlen
	result = result.Store(to_z3_integer(-1), to_z3_integer(int64(len(expr.Value))))
	for i := 0; i < len(expr.Value); i++ {
		result = result.Store(to_z3_integer(int64(i)), to_z3_integer(int64(expr.Value[i])))
	}
	return result
}

func (zt *Z3Translator) VisitArrayConstant(expr *symbolic.ArrayConstant) interface{} {
	return zt.ctx.ConstArray(zt.ctx.IntSort(), expr.Value.Accept(zt).(z3.Value))
}

func (zt *Z3Translator) VisitArraySelect(expr *symbolic.ArraySelect) interface{} {
	return expr.Array.Accept(zt).(z3.Array).Select(expr.Index.Accept(zt).(z3.Value))
}

func (zt *Z3Translator) VisitArrayStore(expr *symbolic.ArrayStore) interface{} {
	return expr.Array.Accept(zt).(z3.Array).Store(expr.Index.Accept(zt).(z3.Value), expr.Value.Accept(zt).(z3.Value))
}
