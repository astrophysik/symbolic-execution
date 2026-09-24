# Домашнее задание 2: Символьные выражения и трансляция в SMT

## Цель
Создать систему символьных выражений для представления вычислений и реализовать их трансляцию в SMT-формулы для Z3 solver. 
Фокус на выражениях (expressions), а не на инструкциях (statements).

## Важно
Весь код должен быть реализован в общей кодовой базе: `../internal/`
Это задание позволяет реализовать модули symbolic и translator, которые работают совместно с SSA из ДЗ_1.

## Теоретические основы

### Символьные выражения
Символьные выражения представляют вычисления в абстрактном виде, где переменные могут иметь символьные (неконкретные) значения.

### SMT (Satisfiability Modulo Theories)
SMT - расширение SAT для работы с различными теориями (арифметика, массивы, строки и т.д.).

## Задания

### Задание 2.1: Проектирование иерархии символьных выражений

Создайте интерфейс и базовые типы символьных выражений:

```go
// Базовый интерфейс для всех символьных выражений
type SymbolicExpression interface {
    // Возвращает тип выражения
    Type() ExpressionType
    
    // Возвращает строковое представление
    String() string
    
    // Принимает visitor для обхода
    Accept(visitor Visitor) interface{}
}

// Типы выражений
type ExpressionType int

const (
    IntType ExpressionType = iota
    BoolType
    ArrayType
    // Добавьте другие типы по необходимости
)
```

### Задание 2.2: Visitor Pattern для обхода выражений

Реализуйте паттерн Visitor для работы с символьными выражениями:

```go
type Visitor interface {
    VisitVariable(expr *SymbolicVariable) interface{}
    VisitIntConstant(expr *IntConstant) interface{}
    VisitBoolConstant(expr *BoolConstant) interface{}
    VisitBinaryOperation(expr *BinaryOperation) interface{}
    VisitLogicalOperation(expr *LogicalOperation) interface{}
    // Добавьте методы для других типов выражений
}

// Пример реализации - вывод выражения в строку
type StringPrinter struct {
    result strings.Builder
}

func (sp *StringPrinter) VisitVariable(expr *SymbolicVariable) interface{} {
    // Реализация вывода переменной
}

// И т.д. для других типов
```

### Задание 2.3: Трансляция в Z3/SMT

Создайте translator для конвертации символьных выражений в Z3 формулы:

```go
type Z3Translator struct {
    ctx     *z3.Context
    solver  *z3.Solver
    vars    map[string]z3.Value  // Кэш переменных
}

func NewZ3Translator(ctx *z3.Context, solver *z3.Solver) *Z3Translator {
    return &Z3Translator{
        ctx:    ctx,
        solver: solver,
        vars:   make(map[string]z3.Value),
    }
}

func (zt *Z3Translator) VisitVariable(expr *SymbolicVariable) interface{} {
    // Если переменная уже создана, вернуть её
    if v, exists := zt.vars[expr.Name]; exists {
        return v
    }
    
    // Создать новую переменную в Z3
    switch expr.Type {
    case IntType:
        v := zt.ctx.IntConst(expr.Name)
        zt.vars[expr.Name] = v
        return v
    case BoolType:
        v := zt.ctx.BoolConst(expr.Name)
        zt.vars[expr.Name] = v
        return v
    }
    return nil
}

func (zt *Z3Translator) VisitBinaryOperation(expr *BinaryOperation) interface{} {
    left := expr.Left.Accept(zt)
    right := expr.Right.Accept(zt)
    
    // Приведение типов и выполнение операции
    switch expr.Operator {
    case ADD:
        return left.(z3.Int).Add(right.(z3.Int))
    case EQ:
        return left.(z3.Int).Eq(right.(z3.Int))
    // И т.д. для других операций
    }
    return nil
}
```

### Задание 2.4: Интеграция и тестирование

Создайте тесты для демонстрации работы с выражениями и их трансляцией в SMT.

## Что сдавать

**Реализация в общей кодовой базе:**
- `../internal/symbolic/expressions.go` - дополните методы для символьных выражений
- `../internal/translator/z3_translator.go` - реализуйте трансляцию выражений в Z3

**В папке homework2:**

```  
homework2/
├── examples/
│   └── test_functions.go    // Функции для тестирования
├── main.go                  // Демонстрация (использует internal/*)
└── README.md
```
