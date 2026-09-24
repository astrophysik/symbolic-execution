package internal

import (
	"golang.org/x/tools/go/ssa"
	"symbolic-execution-course/internal/translator"
)

type Analyser struct {
	Package      *ssa.Package
	StatesQueue  PriorityQueue
	PathSelector PathSelector
	Results      []Interpreter
	Z3Translator *translator.Z3Translator
}

func Analyse(source string, functionName string) []Interpreter {
	// TODO implement me
	panic("implement me")
}
