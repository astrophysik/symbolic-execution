# Домашнее задание 1: Построение Static Single Assignment (SSA)

## Цель
Научиться извлекать промежуточное представление (IR) из Go компилятора для дальнейшего анализа.

## Важно
Весь код должен быть реализован в общей кодовой базе: `../internal/ssa/`
Это задание позволяет реализовать модуль построения SSA, который будет использоваться в других частях курса.

### Go SSA IR
Go компилятор предоставляет SSA (Static Single Assignment) представление через пакет `golang.org/x/tools/go/ssa`.

## Задание
Создайте программу для извлечения SSA представления:

```go
// Пример структуры для анализа
func analyzeFunction(source string, funcName string) (*ssa.Function, error) {
    // 1. Парсинг исходного кода
    // 2. Создание SSA представления  
    // 3. Поиск функции по имени
    // 4. Возврат SSA функции
}
```

**Требования:**
- Используйте пакет `golang.org/x/tools/go/ssa`
- Научитесь получать список инструкций для каждого базового блока
- Выведите информацию о блоках и их связях

**В папке homework1:**

```
homework1/
├── examples/
│   └── test_functions.go    // Функции для тестирования SSA
├── main.go                  // Демонстрация (использует internal/ssa)
└── README.md
```

## Полезные ресурсы
- [golang.org/x/tools/go/ssa](https://pkg.go.dev/golang.org/x/tools/go/ssa)
- [SSA форма в компиляторах](https://en.wikipedia.org/wiki/Static_single_assignment_form)
