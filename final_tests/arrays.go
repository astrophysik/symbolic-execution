package final_tests

func DefaultBooleanValues() []bool {
	array := make([]bool, 3)
	if !array[1] {
		return array
	}
	return array
}

func ByteArray(a []byte, x byte) byte {
	if len(a) != 2 {
		return 255
	}
	a[0] = 5
	a[1] = x
	if a[0]+a[1] > 20 {
		return 1
	}
	return 0
}

func CharSizeAndIndex(a []rune, x rune) byte {
	if a == nil || len(a) <= int(x) || x < 1 {
		return 255
	}
	b := make([]rune, x)
	b[0] = 5
	a[x] = x
	if b[0]+a[x] > 7 {
		return 1
	}
	return 0
}

func BooleanArray(arr []bool) int {
	if len(arr) == 0 {
		return 1
	}
	if arr[0] {
		return 2
	}
	arr[0] = true
	return 3
}

func CreateArray(x, y, length int) []*ObjectWithPrimitivesClass {
	if length < 3 {
		return nil
	}

	array := make([]*ObjectWithPrimitivesClass, length)

	for i := 0; i < len(array); i++ {
		array[i] = NewObjectWithPrimitivesClass()
		array[i].x = x + i
		array[i].y = y + i
	}

	return array
}

func IsIdentityMatrix(matrix [][]int) bool {
	if len(matrix) < 3 {
		return false
	}
	for i := 0; i < len(matrix); i++ {
		if len(matrix[i]) != len(matrix) {
			return false
		}
		for j := 0; j < len(matrix[i]); j++ {
			if i == j && matrix[i][j] != 1 {
				return false
			}
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func ReallyMultiDimensionalArray(array [][][]int) [][][]int {
	if array[1][2][3] != 12345 {
		array[1][2][3] = 12345
	} else {
		array[1][2][3] -= 12345 * 2
	}
	return array
}

func FillMultiArrayWithArray(value []int) [][]int {
	if len(value) < 2 {
		return make([][]int, 0)
	}
	for i := range value {
		value[i] += i
	}
	length := 3
	array := make([][]int, length)
	for i := 0; i < length; i++ {
		array[i] = value
	}
	return array
}
