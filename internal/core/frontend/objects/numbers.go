package objects

type ObjectNumber interface {
	Object
	ToInt() int
	ToFloat() float64
}

type ObjectInt struct {
	Value int
}

func (num *ObjectInt) ObjectValue() interface{} {
	return num.Value
}
func (num *ObjectInt) ToInt() int {
	return num.Value
}
func (num *ObjectInt) ToFloat() float64 {
	return float64(num.Value)
}

type ObjectFloat struct {
	Value float64
}

func (num *ObjectFloat) ObjectValue() interface{} {
	return num.Value
}
func (num *ObjectFloat) ToInt() int {
	return int(num.Value)
}
func (num *ObjectFloat) ToFloat() float64 {
	return num.Value
}
