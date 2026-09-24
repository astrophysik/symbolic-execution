package internal

import "math/rand"

type PathSelector interface {
	CalculatePriority(interpreter Interpreter) int
}

type DfsPathSelector struct {
	counter int // int.min_value
}

func (dfs *DfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	dfs.counter++
	return dfs.counter
}

type BfsPathSelector struct {
	counter int // int.max_value
}

func (bfs *BfsPathSelector) CalculatePriority(interpreter Interpreter) int {
	bfs.counter--
	return bfs.counter
}

type RandomPathSelector struct{}

func (random *RandomPathSelector) CalculatePriority(interpreter Interpreter) int {
	return rand.Int()
}
