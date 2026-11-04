// Package ssa предоставляет функции для построения SSA представления
package ssa

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// Builder отвечает за построение SSA из исходного кода Go
type Builder struct {
	fset *token.FileSet
}

// NewBuilder создаёт новый экземпляр Builder
func NewBuilder() *Builder {
	return &Builder{
		fset: token.NewFileSet(),
	}
}

// ParseAndBuildSSA парсит исходный код Go и создаёт SSA представление
// Возвращает SSA программу и функцию по имени
func (b *Builder) ParseAndBuildSSA(source string, funcName string) (*ssa.Function, error) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "", source, 0)
	if err != nil {
		return nil, err
	}
	files := []*ast.File{file}

	pkg := types.NewPackage("program", "")

	program, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fileSet, pkg, files, ssa.SanityCheckFunctions)
	
	if err != nil {
		return nil, err
	}

	return program.Func(funcName), nil
}
