package runtime

func iadd(a, b *Integer) *Integer {
	return &Integer{a.Value + b.Value}
}
func isub(a, b *Integer) *Integer {
	return &Integer{a.Value - b.Value}
}
func ilt(a, b *Integer) *Boolean {
	if a.Value < b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
