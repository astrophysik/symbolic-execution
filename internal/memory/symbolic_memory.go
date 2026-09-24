package memory

import "symbolic-execution-course/internal/symbolic"

type Memory interface {
	Allocate(tpe symbolic.ExpressionType) *symbolic.Ref

	AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression)

	GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression

	AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression)

	GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression
}

type SymbolicMemory struct {
	// TODO: Реализуйте внутреннее представление символьной памяти
}

func NewSymbolicMemory() SymbolicMemory {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) Allocate(tpe symbolic.ExpressionType) *symbolic.Ref {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	//TODO implement me
	panic("implement me")
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	//TODO implement me
	panic("implement me")
}
