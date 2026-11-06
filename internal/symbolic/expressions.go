// Package symbolic содержит конкретные реализации символьных выражений
package symbolic

import (
	"fmt"
	"strings"
)

// SymbolicExpression - базовый интерфейс для всех символьных выражений
type SymbolicExpression interface {
	// Type возвращает тип выражения
	Type() ExpressionType

	// String возвращает строковое представление выражения
	String() string

	// Accept принимает visitor для обхода дерева выражений
	Accept(visitor Visitor) interface{}
}

// SymbolicVariable представляет символьную переменную
type SymbolicVariable struct {
	Name     string
	TypeInfo NestedExpressionType
}

// NewSymbolicVariable создаёт новую символьную переменную
func NewSymbolicVariable(name string, exprType ExpressionType) *SymbolicVariable {
	return &SymbolicVariable{
		Name: name,
		TypeInfo: NestedExpressionType{
			ExprType:  exprType,
			InnerType: nil,
		},
	}
}

// Type возвращает тип переменной
func (sv *SymbolicVariable) Type() ExpressionType {
	return sv.TypeInfo.ExprType
}

// String возвращает строковое представление переменной
func (sv *SymbolicVariable) String() string {
	return sv.Name
}

// Accept реализует Visitor pattern
func (sv *SymbolicVariable) Accept(visitor Visitor) interface{} {
	return visitor.VisitVariable(sv)
}

// IntConstant представляет целочисленную константу
type IntConstant struct {
	Value int64
}

// NewIntConstant создаёт новую целочисленную константу
func NewIntConstant(value int64) *IntConstant {
	return &IntConstant{Value: value}
}

// Type возвращает тип константы
func (ic *IntConstant) Type() ExpressionType {
	return IntType
}

// String возвращает строковое представление константы
func (ic *IntConstant) String() string {
	return fmt.Sprintf("%d", ic.Value)
}

// Accept реализует Visitor pattern
func (ic *IntConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitIntConstant(ic)
}

// BoolConstant представляет булеву константу
type BoolConstant struct {
	Value bool
}

// NewBoolConstant создаёт новую булеву константу
func NewBoolConstant(value bool) *BoolConstant {
	return &BoolConstant{Value: value}
}

// Type возвращает тип константы
func (bc *BoolConstant) Type() ExpressionType {
	return BoolType
}

// String возвращает строковое представление константы
func (bc *BoolConstant) String() string {
	return fmt.Sprintf("%t", bc.Value)
}

// Accept реализует Visitor pattern
func (bc *BoolConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitBoolConstant(bc)
}

type FloatConstant struct {
	Value float64
}

func NewFloatConstant(value float64) *FloatConstant {
	return &FloatConstant{Value: value}
}

func (fc *FloatConstant) Type() ExpressionType {
	return FloatType
}

func (fc *FloatConstant) String() string {
	return fmt.Sprintf("%g", fc.Value)
}

func (fc *FloatConstant) Accept(visitor Visitor) interface{} {
	return visitor.VisitFloatConstant(fc)
}

// BinaryOperation представляет бинарную операцию
type BinaryOperation struct {
	Left     SymbolicExpression
	Right    SymbolicExpression
	Operator BinaryOperator
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// NewBinaryOperation создаёт новую бинарную операцию
func NewBinaryOperation(left, right SymbolicExpression, op BinaryOperator) *BinaryOperation {
	if left.Type() != right.Type() {
		panic("Types mismatched")
	}

	return &BinaryOperation{
		Left:     left,
		Right:    right,
		Operator: op,
	}
}

// Type возвращает результирующий тип операции
func (bo *BinaryOperation) Type() ExpressionType {
	switch bo.Operator {
	case ADD, SUB, MUL, DIV:
		operandType := bo.Left.Type()
		if operandType != IntType && operandType != FloatType {
			panic("arithmetic operations require numeric operands")
		}
		return operandType
	case MOD:
		if bo.Left.Type() == IntType {
			return IntType
		} else {
			errorMessage := fmt.Sprintf("Unexpected type %s in integer binary operation", bo.Left.Type().String())
			panic(errorMessage)
		}
	case EQ, NE:
		return BoolType
	case LT, LE, GT, GE:
		operandType := bo.Left.Type()
		if operandType != IntType && operandType != FloatType {
			panic("ordered comparisons require numeric operands")
		}
		return BoolType
	case BIT_AND, BIT_OR, BIT_XOR, SHL, SHR:
		if bo.Left.Type() != IntType || bo.Right.Type() != IntType {
			panic("bitwise operations require integer operands")
		}
		return IntType
	default:
		errorMessage := fmt.Sprintf("Unknown binary type %s", bo.Operator.String())
		panic(errorMessage)
	}
}

// String возвращает строковое представление операции
func (bo *BinaryOperation) String() string {
	return "(" + bo.Left.String() + " " + bo.Operator.String() + " " + bo.Right.String() + ")"
}

// Accept реализует Visitor pattern
func (bo *BinaryOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitBinaryOperation(bo)
}

// LogicalOperation представляет логическую операцию
type LogicalOperation struct {
	Operands []SymbolicExpression
	Operator LogicalOperator
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// NewLogicalOperation создаёт новую логическую операцию
func NewLogicalOperation(operands []SymbolicExpression, op LogicalOperator) *LogicalOperation {
	for _, operand := range operands {
		if operand.Type() != BoolType {
			panic("non bool type operand in logic operation")
		}
	}
	return &LogicalOperation{
		Operands: operands,
		Operator: op,
	}
}

// Type возвращает тип логической операции (всегда bool)
func (lo *LogicalOperation) Type() ExpressionType {
	return BoolType
}

// String возвращает строковое представление логической операции
func (lo *LogicalOperation) String() string {
	switch lo.Operator {
	case AND:
		stringOperands := []string{}
		for _, v := range lo.Operands {
			stringOperands = append(stringOperands, v.String())
		}
		view := strings.Join(stringOperands, " && ")
		return "(" + view + ")"
	case OR:
		stringOperands := []string{}
		for _, v := range lo.Operands {
			stringOperands = append(stringOperands, v.String())
		}
		view := strings.Join(stringOperands, " || ")
		return "(" + view + ")"
	case NOT:
		return "!" + lo.Operands[0].String()
	case IMPLIES:
		return "(" + lo.Operands[0].String() + " => " + lo.Operands[1].String() + ")"
	default:
		panic("Unknown logic operation")
	}
}

// Accept реализует Visitor pattern
func (lo *LogicalOperation) Accept(visitor Visitor) interface{} {
	return visitor.VisitLogicalOperation(lo)
}

// Операторы для бинарных выражений
type BinaryOperator int

const (
	// Арифметические операторы
	ADD BinaryOperator = iota
	SUB
	MUL
	DIV
	MOD

	// Операторы сравнения
	EQ // равно
	NE // не равно
	LT // меньше
	LE // меньше или равно
	GT // больше
	GE // больше или равно

	// Bit operations
	BIT_AND
	BIT_OR
	BIT_XOR
	SHL
	SHR
)

// String возвращает строковое представление оператора
func (op BinaryOperator) String() string {
	switch op {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case MOD:
		return "%"
	case EQ:
		return "=="
	case NE:
		return "!="
	case LT:
		return "<"
	case LE:
		return "<="
	case GT:
		return ">"
	case GE:
		return ">="
	case BIT_AND:
		return "&"
	case BIT_OR:
		return "|"
	case BIT_XOR:
		return "^"
	case SHL:
		return "<<"
	case SHR:
		return ">>"
	default:
		return "unknown"
	}
}

// Логические операторы
type LogicalOperator int

const (
	AND LogicalOperator = iota
	OR
	NOT
	IMPLIES
)

// String возвращает строковое представление логического оператора
func (op LogicalOperator) String() string {
	switch op {
	case AND:
		return "&&"
	case OR:
		return "||"
	case NOT:
		return "!"
	case IMPLIES:
		return "=>"
	default:
		return "unknown"
	}
}

type Ref struct {
	// TODO: Выбрать и написать внутреннее представление символьной ссылки
}

func (ref *Ref) Type() ExpressionType {
	panic("не реализовано")
}

func (ref *Ref) String() string {
	panic("не реализовано")
}

func (ref *Ref) Accept(visitor Visitor) interface{} {
	panic("не реализовано")
}

// TODO: Добавьте дополнительные типы выражений по необходимости:
// - UnaryOperation (унарные операции: -x, !x)
// - ArrayAccess (доступ к элементам массива: arr[index])
// - FunctionCall (вызовы функций: f(x, y))
// - ConditionalExpression (тернарный оператор: condition ? true_expr : false_expr)
