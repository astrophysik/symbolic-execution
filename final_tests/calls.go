package final_tests

type InvokeClass struct {
	Value int
}

func (i *InvokeClass) DivBy(den int) int {
	return i.Value / den
}

func (i *InvokeClass) UpdateValue(newValue int) {
	i.Value = newValue
}

func mult(a, b int) int {
	return a * b
}

func half(a int) int {
	return a / 2
}

func SimpleFormula(fst, snd int) int {
	if fst < 100 {
		return 0
	} else if snd < 100 {
		return 0
	}

	x := fst + 5
	y := half(snd)

	return mult(x, y)
}

func initialize(value int) *InvokeClass {
	objectValue := &InvokeClass{
		Value: value,
	}
	return objectValue
}

func CreateObjectFromValue(value int) *InvokeClass {
	if value == 0 {
		value = 1
	}
	return initialize(value)
}

func changeValue(objectValue *InvokeClass, value int) {
	objectValue.Value = value
}

func ChangeObjectValueByMethod(objectValue *InvokeClass) *InvokeClass {
	objectValue.Value = 1
	changeValue(objectValue, 4)
	return objectValue
}

func getFive() int {
	return 5
}

func getTwo() int {
	return 2
}

func ParticularValue(invokeObject *InvokeClass) *InvokeClass {
	if invokeObject.Value < 0 {
		return nil
	}
	x := getFive() * getTwo()
	y := getFive() / getTwo()

	invokeObject.Value = x + y
	return invokeObject
}

func getNull() *InvokeClass {
	return nil
}

func GetNullOrValue(invokeObject *InvokeClass) *InvokeClass {
	if invokeObject.Value < 100 {
		return getNull()
	}
	invokeObject.Value = getFive()
	return invokeObject
}
