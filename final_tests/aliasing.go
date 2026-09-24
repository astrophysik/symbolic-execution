package final_tests

type Foo struct {
	a int
}

func Aliasing(foo1 *Foo, foo2 *Foo) int {
	foo2.a = 5
	foo1.a = 2
	if foo2.a == 2 {
		return 4
	}
	return 5
}

func ArrayAliasing(arr1 []int, arr2 []int) int {
	arr2[1] = 5
	arr1[1] = 2
	if arr2[1] == 2 {
		return 4
	}
	return 5
}
