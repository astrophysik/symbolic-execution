// Package z3wrapper предоставляет удобную обёртку для работы с Z3 solver
// в контексте символьного исполнения и анализа Go программ.
package z3wrapper

import (
	"fmt"
	"strconv"

	"github.com/ebukreev/go-z3/z3"
)

// Solver представляет обёртку над Z3 solver
type Solver struct {
	ctx    *z3.Context
	solver *z3.Solver
}

// NewSolver создаёт новый экземпляр Z3 solver
func NewSolver() *Solver {
	config := z3.NewContextConfig()
	ctx := z3.NewContext(config)
	solver := z3.NewSolver(ctx)

	return &Solver{
		ctx:    ctx,
		solver: solver,
	}
}

// Close освобождает ресурсы solver'а
func (s *Solver) Close() {
	// В этой версии Z3 нет метода Close для solver и context
	// Ресурсы освобождаются автоматически Go GC
}

// Context возвращает контекст Z3
func (s *Solver) Context() *z3.Context {
	return s.ctx
}

// Assert добавляет ограничение в solver
func (s *Solver) Assert(constraint z3.Bool) {
	s.solver.Assert(constraint)
}

// Check проверяет выполнимость текущих ограничений
func (s *Solver) Check() (bool, error) {
	sat, err := s.solver.Check()
	return sat, err
}

// Model возвращает модель, если ограничения выполнимы
func (s *Solver) Model() *z3.Model {
	return s.solver.Model()
}

// Push сохраняет текущее состояние solver'а
func (s *Solver) Push() {
	s.solver.Push()
}

// Pop восстанавливает предыдущее состояние solver'а
func (s *Solver) Pop() {
	s.solver.Pop()
}

// CreateIntVar создаёт целочисленную переменную
func (s *Solver) CreateIntVar(name string) z3.Int {
	return s.ctx.IntConst(name)
}

// CreateBoolVar создаёт булеву переменную
func (s *Solver) CreateBoolVar(name string) z3.Bool {
	return s.ctx.BoolConst(name)
}

// CreateIntLit создаёт целочисленную константу
func (s *Solver) CreateIntLit(value int64) z3.Int {
	return s.ctx.FromInt(value, s.ctx.IntSort()).(z3.Int)
}

// IsSatisfiable проверяет, выполнимы ли текущие ограничения
func (s *Solver) IsSatisfiable() (bool, error) {
	return s.solver.Check()
}

// GetIntValue получает значение целочисленной переменной из модели
func (s *Solver) GetIntValue(model *z3.Model, variable z3.Int) (int64, error) {
	value := model.Eval(variable, false)
	if value == nil {
		return 0, fmt.Errorf("variable not found in model")
	}

	// Конвертируем в строку и затем в int64
	str := value.String()
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse integer value: %v", err)
	}

	return result, nil
}

// GetBoolValue получает значение булевой переменной из модели
func (s *Solver) GetBoolValue(model *z3.Model, variable z3.Bool) (bool, error) {
	value := model.Eval(variable, false)
	if value == nil {
		return false, fmt.Errorf("variable not found in model")
	}

	// Проверяем строковое представление
	str := value.String()
	switch str {
	case "true":
		return true, nil
	case "false":
		return false, nil
	default:
		return false, fmt.Errorf("unexpected boolean value: %s", str)
	}
}
