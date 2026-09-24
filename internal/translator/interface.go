// Package translator предоставляет интерфейс для трансляции символьных выражений
package translator

import (
	"symbolic-execution-course/internal/symbolic"
)

// Translator интерфейс для конвертации символьных выражений в различные форматы
type Translator interface {
	// TranslateExpression конвертирует символьное выражение в целевой формат
	TranslateExpression(expr symbolic.SymbolicExpression) (interface{}, error)

	// GetContext возвращает контекст транслятора (например, Z3 Context)
	GetContext() interface{}

	// Reset сбрасывает состояние транслятора
	Reset()
}

// ExpressionTranslator базовый интерфейс для трансляции выражений
type ExpressionTranslator interface {
	// Visit методы для различных типов выражений
	VisitVariable(expr *symbolic.SymbolicVariable) (interface{}, error)
	VisitIntConstant(expr *symbolic.IntConstant) (interface{}, error)
	VisitBoolConstant(expr *symbolic.BoolConstant) (interface{}, error)
	VisitBinaryOperation(expr *symbolic.BinaryOperation) (interface{}, error)
	VisitLogicalOperation(expr *symbolic.LogicalOperation) (interface{}, error)
}

// TranslationError представляет ошибку трансляции
type TranslationError struct {
	Message    string
	Expression symbolic.SymbolicExpression
}

func (te *TranslationError) Error() string {
	return te.Message
}

// NewTranslationError создаёт новую ошибку трансляции
func NewTranslationError(message string, expr symbolic.SymbolicExpression) *TranslationError {
	return &TranslationError{
		Message:    message,
		Expression: expr,
	}
}
