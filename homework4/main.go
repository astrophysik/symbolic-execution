package main

import (
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
)

func main() {
	var mem = memory.NewSymbolicMemory()
	var array = mem.AllocateArray(symbolic.IntType)

	mem.AssignToArray(array, 5, symbolic.NewIntConstant(10))

	var fromArray = mem.GetFromArray(array, 5)
	println(fromArray)

	var anotherFromArray = mem.GetFromArray(array, 10)
	println(anotherFromArray)
}