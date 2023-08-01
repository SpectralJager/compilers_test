package runtime

import "fmt"

// int functions

func iadd(a, b *Integer) *Integer {
	return &Integer{a.Value + b.Value}
}
func isub(a, b *Integer) *Integer {
	return &Integer{a.Value - b.Value}
}
func imul(a, b *Integer) *Integer {
	return &Integer{a.Value * b.Value}
}
func idiv(a, b *Integer) *Integer {
	return &Integer{a.Value / b.Value}
}
func ilt(a, b *Integer) *Boolean {
	if a.Value < b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func igt(a, b *Integer) *Boolean {
	if a.Value > b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func igeq(a, b *Integer) *Boolean {
	if a.Value >= b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func ileq(a, b *Integer) *Boolean {
	if a.Value <= b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func ieq(a, b *Integer) *Boolean {
	if a.Value == b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func itof(a *Integer) *Float {
	return &Float{float64(a.Value)}
}
func itos(a *Integer) *String {
	return &String{fmt.Sprint(a.Value)}
}
func itob(a *Integer) *Boolean {
	if a.Value == 0 {
		return &Boolean{false}
	}
	return &Boolean{true}
}
