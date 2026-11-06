package memory

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"
)

type Memory interface {
	Allocate(tpe symbolic.ExpressionType) *symbolic.Ref
	AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression)
	GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression

	AllocateStruct(fieldTypes ...symbolic.ExpressionType) *symbolic.Ref
	SymbolicStruct(name string, fieldTypes ...symbolic.ExpressionType) *symbolic.Ref
	AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression)
	GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression

	AllocateArray(elementType symbolic.ExpressionType) *symbolic.Ref
	SymbolicArray(name string, elementType symbolic.ExpressionType) *symbolic.Ref
	AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression)
	GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression
}

type fieldKey struct {
	index     int
	valueType symbolic.ExpressionType
}

type SymbolicMemory struct {
	nextAddress int64
	fields      map[fieldKey]symbolic.SymbolicExpression
	arrays      map[symbolic.ExpressionType]symbolic.SymbolicExpression
}

func NewSymbolicMemory() SymbolicMemory {
	return SymbolicMemory{fields: make(map[fieldKey]symbolic.SymbolicExpression), arrays: make(map[symbolic.ExpressionType]symbolic.SymbolicExpression)}
}

func (mem *SymbolicMemory) Allocate(tpe symbolic.ExpressionType) *symbolic.Ref {
	// primitive type = struct {primitive type}
	return mem.AllocateStruct(tpe)
}

func (mem *SymbolicMemory) AssignValue(ref *symbolic.Ref, value symbolic.SymbolicExpression) {
	mem.AssignField(ref, 0, value)
}

func (mem *SymbolicMemory) GetValue(ref *symbolic.Ref) symbolic.SymbolicExpression {
	return mem.GetFieldValue(ref, 0)
}

func (mem *SymbolicMemory) freshAddress() symbolic.SymbolicExpression {
	mem.nextAddress++
	return symbolic.NewIntConstant(mem.nextAddress)
}

func (mem *SymbolicMemory) AllocateStruct(fieldTypes ...symbolic.ExpressionType) *symbolic.Ref {
	ref := symbolic.NewStructRef(mem.freshAddress(), fieldTypes...)
	for i, t := range fieldTypes {
		mem.AssignField(ref, i, t.Zero())
	}
	return ref
}

func (mem *SymbolicMemory) AllocateArray(elementType symbolic.ExpressionType) *symbolic.Ref {
	ref := symbolic.NewArrayRef(mem.freshAddress(), elementType)
	mem.arrays[elementType] = symbolic.NewArrayStore(mem.arrayHeap(elementType), ref.Address, symbolic.NewArrayConstant(elementType.Zero()))
	return ref
}

func (mem *SymbolicMemory) fieldHeap(index int, valueType symbolic.ExpressionType) symbolic.SymbolicExpression {
	key := fieldKey{index, valueType}
	if h, ok := mem.fields[key]; ok {
		return h
	}
	inner := symbolic.NestedExpressionType{ExprType: valueType}
	h := &symbolic.SymbolicVariable{Name: fmt.Sprintf("memory_field%d_%s", index, valueType), TypeInfo: symbolic.NestedExpressionType{ExprType: symbolic.ArrayType, InnerType: &inner}}
	mem.fields[key] = h
	return h
}

func (mem *SymbolicMemory) arrayHeap(valueType symbolic.ExpressionType) symbolic.SymbolicExpression {
	if h, ok := mem.arrays[valueType]; ok {
		return h
	}
	leaf := symbolic.NestedExpressionType{ExprType: valueType}
	inner := symbolic.NestedExpressionType{ExprType: symbolic.ArrayType, InnerType: &leaf}
	h := &symbolic.SymbolicVariable{Name: fmt.Sprintf("memory_elements_%s", valueType), TypeInfo: symbolic.NestedExpressionType{ExprType: symbolic.ArrayType, InnerType: &inner}}
	mem.arrays[valueType] = h
	return h
}

func (mem *SymbolicMemory) AssignField(ref *symbolic.Ref, fieldIdx int, value symbolic.SymbolicExpression) {
	t := ref.FieldTypes[fieldIdx]
	if value.Type() != t {
		panic("field value type mismatch")
	}
	mem.fields[fieldKey{fieldIdx, t}] = symbolic.NewArrayStore(mem.fieldHeap(fieldIdx, t), ref.Address, value)
}

func (mem *SymbolicMemory) GetFieldValue(ref *symbolic.Ref, fieldIdx int) symbolic.SymbolicExpression {
	return symbolic.NewArraySelect(mem.fieldHeap(fieldIdx, ref.FieldTypes[fieldIdx]), ref.Address)
}

func (mem *SymbolicMemory) AssignToArray(ref *symbolic.Ref, index int, value symbolic.SymbolicExpression) {
	if value.Type() != ref.ElementType {
		panic("array value type mismatch")
	}

	heap := mem.arrayHeap(ref.ElementType)
	array := symbolic.NewArraySelect(heap, ref.Address)
	updated := symbolic.NewArrayStore(array, symbolic.NewIntConstant(int64(index)), value)
	mem.arrays[ref.ElementType] = symbolic.NewArrayStore(heap, ref.Address, updated)
}

func (mem *SymbolicMemory) GetFromArray(ref *symbolic.Ref, index int) symbolic.SymbolicExpression {
	array := symbolic.NewArraySelect(mem.arrayHeap(ref.ElementType), ref.Address)
	return symbolic.NewArraySelect(array, symbolic.NewIntConstant(int64(index)))
}
