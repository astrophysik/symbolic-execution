package symbolic

// Visitor интерфейс для обхода символьных выражений (Visitor Pattern)
type Visitor interface {
	VisitVariable(expr *SymbolicVariable) interface{}
	VisitIntConstant(expr *IntConstant) interface{}
	VisitBoolConstant(expr *BoolConstant) interface{}
	VisitFloatConstant(expr *FloatConstant) interface{}
	VisitBinaryOperation(expr *BinaryOperation) interface{}
	VisitLogicalOperation(expr *LogicalOperation) interface{}
	VisitRef(expr *Ref) interface{}
	VisitStringConstant(expr *StringConstant) interface{}
	VisitArrayConstant(expr *ArrayConstant) interface{}
	VisitArraySelect(expr *ArraySelect) interface{}
	VisitArrayStore(expr *ArrayStore) interface{}
	// TODO: Добавьте методы для других типов выражений по мере необходимости
}
