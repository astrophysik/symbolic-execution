package final_tests

func Complement(x int) bool {
	return ^x == 1
}

func Xor(x, y int) bool {
	return (x ^ y) == 0
}

func Or(val int) bool {
	return (val | 7) == 15
}

func And(value int) bool {
	return (value & (value - 1)) == 0
}

func BooleanNot(boolA, boolB bool) int {
	d := boolA && boolB
	e := (!boolA) || boolB
	if d && e {
		return 100
	}
	return 200
}

func BooleanXorCompare(aBool, bBool bool) int {
	if aBool != bBool {
		return 1
	}
	return 0
}

func ShlWithBigLongShift(shift int64) int {
	if shift < 40 {
		return 1
	}
	if (0x77777777 << shift) == 0x77777770 {
		return 2
	}
	return 3
}
