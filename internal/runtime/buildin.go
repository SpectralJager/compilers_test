package runtime

var buildins = map[string]any{
	"iadd": iadd,
}

func iadd(a, b *Integer) *Integer {
	return &Integer{a.Value + b.Value}
}
