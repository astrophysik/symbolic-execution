// Package symbolic определяет базовые типы символьных выражений
package symbolic

// ExpressionType представляет тип символьного выражения
type ExpressionType int

const (
	IntType ExpressionType = iota
	BoolType
	ArrayType
	FloatType
	RefType
	StringType
	// Добавьте другие типы по необходимости
)

func (a ExpressionType) Zero() SymbolicExpression {
	switch a {
	case IntType:
		return NewIntConstant(0)
	case BoolType:
		return NewBoolConstant(false)
	case FloatType:
		return NewFloatConstant(0)
	case StringType:
		return NewStringConstant("")
	case RefType:
		return NewStructRef(NewIntConstant(0))
	default:
		panic("unsupported ExpressionType")
	}
}

type NestedExpressionType struct {
	ExprType  ExpressionType
	InnerType *NestedExpressionType
}

func (a NestedExpressionType) Equals(b NestedExpressionType) bool {
	if a.ExprType != b.ExprType {
		return false
	}
	if a.InnerType == nil || b.InnerType == nil {
		return a.InnerType == nil && b.InnerType == nil
	}
	return a.InnerType.Equals(*b.InnerType)
}

// String возвращает строковое представление типа
func (et ExpressionType) String() string {
	switch et {
	case IntType:
		return "int"
	case BoolType:
		return "bool"
	case ArrayType:
		return "array"
	case FloatType:
		return "float"
	case RefType:
		return "ref"
	case StringType:
		return "string"
	default:
		return "unknown"
	}
}
