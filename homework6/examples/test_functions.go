package main

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func testRecursive(x int) int {
	if x < 5 {
		return factorial(x)
	}
	return -1
}

func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return isOdd(n - 1)
}

func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func testMutualRecursion(x int) bool {
	if x >= 0 {
		return isEven(x)
	}
	return false
}

func testArrayOperations(n int) int {
	arr := [5]int{1, 2, 3, 4, 5}
	sum := 0

	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			sum += arr[i] * n
		} else {
			sum += arr[i]
		}
	}
	return sum
}

func testSliceDynamic(n int) int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i * i
	}

	result := 0
	for _, val := range slice {
		if val%2 == 0 {
			result += val
		} else {
			result -= val
		}
	}
	return result
}

type Point struct {
	X, Y int
}

type Rectangle struct {
	TopLeft, BottomRight Point
}

func (r Rectangle) Area() int {
	width := r.BottomRight.X - r.TopLeft.X
	height := r.TopLeft.Y - r.BottomRight.Y
	if width < 0 || height < 0 {
		return -1
	}
	return width * height
}

func testStructOperations(x1, y1, x2, y2 int) int {
	rect := Rectangle{
		TopLeft:     Point{X: x1, Y: y1},
		BottomRight: Point{X: x2, Y: y2},
	}
	return rect.Area()
}

func testPointers(x int) int {
	a := x
	b := &a
	c := &b

	if x > 0 {
		**c = **c * 2
	} else {
		*b = *b - 1
	}

	return a + *b
}

func modifyViaPointer(ptr *int, n int) {
	if n > 0 {
		*ptr = *ptr + n
	} else {
		*ptr = *ptr - 1
	}
}

func testPointerParameters(x, y int) int {
	value := x
	modifyViaPointer(&value, y)
	return value
}

func testComplexLoop(n int) int {
	sum := 0
	i := 0

outer:
	for i < n {
		j := 0
		for j < n {
			if i*j > 100 {
				break outer
			}
			if (i+j)%2 == 0 {
				sum += i + j
				j++
				continue
			}
			sum -= i - j
			j++
		}
		i++
	}
	return sum
}

func multipleReturns(x, y int) (int, bool) {
	if x > y {
		return x - y, true
	}
	return y - x, false
}

func testMultipleReturnValues(a, b int) int {
	diff, isPositive := multipleReturns(a, b)

	if isPositive {
		return diff * 2
	}
	return -diff
}

type Shape interface {
	Area() int
	Perimeter() int
}

type Circle struct {
	Radius int
}

func (c Circle) Area() int {
	return 3 * c.Radius * c.Radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.Radius
}

type Square struct {
	Side int
}

func (s Square) Area() int {
	return s.Side * s.Side
}

func (s Square) Perimeter() int {
	return 4 * s.Side
}

func testInterface(shape Shape, multiplier int) int {
	area := shape.Area()
	perimeter := shape.Perimeter()

	if area > perimeter {
		return area * multiplier
	}
	return perimeter * multiplier
}

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, nil // имитация ошибки
	}
	return a / b, nil
}

func testErrorHandling(x, y int) int {
	result, err := safeDivide(x, y)
	if err != nil {
		return -1
	}

	if result > 10 {
		return result * 2
	}
	return result
}

func testMixedConditions(arr []int, ptr *int, n int) bool {
	if len(arr) > 0 && ptr != nil && n > 0 {
		if arr[0] == *ptr || n%2 == 0 {
			return true
		}
	}

	if ptr == nil && len(arr) == n {
		return false
	}

	return len(arr) > n
}

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

func sumTree(node *TreeNode) int {
	if node == nil {
		return 0
	}

	sum := node.Value
	for _, child := range node.Children {
		sum += sumTree(child)
	}
	return sum
}

func testTreeStructure(values []int) int {
	if len(values) == 0 {
		return 0
	}

	root := &TreeNode{Value: values[0]}
	if len(values) > 1 {
		root.Children = append(root.Children, &TreeNode{Value: values[1]})
	}
	if len(values) > 2 {
		root.Children = append(root.Children, &TreeNode{Value: values[2]})
	}

	return sumTree(root)
}

func comprehensiveTest(x int, y int, useRecursion bool) int {
	a := x
	ptr := &a

	var result int
	if useRecursion {
		result = testRecursive(x)
	}

	rect := Rectangle{
		TopLeft:     Point{X: 0, Y: y},
		BottomRight: Point{X: x, Y: 0},
	}
	area := rect.Area()

	if result > area && *ptr > 0 {
		return result + area
	} else if result < 0 {
		return result - area
	}

	return area
}
