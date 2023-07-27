package vm

func Add[T int](val1, val2 T) T {
	return val1 + val2
}

func Sub[T int](val1, val2 T) T {
	return val1 - val2
}

func Mul[T int](val1, val2 T) T {
	return val1 * val2
}

func Div[T int](val1, val2 T) T {
	return val1 / val2
}

func Neg[T int](val1 T) T {
	return -val1
}
