package runtime

import (
	"fmt"
	"strconv"
)

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

// float functions

func fadd(a, b *Float) *Float {
	return &Float{a.Value + b.Value}
}
func fsub(a, b *Float) *Float {
	return &Float{a.Value - b.Value}
}
func fmul(a, b *Float) *Float {
	return &Float{a.Value * b.Value}
}
func fdiv(a, b *Float) *Float {
	return &Float{a.Value / b.Value}
}
func flt(a, b *Float) *Boolean {
	if a.Value < b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func fgt(a, b *Float) *Boolean {
	if a.Value > b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func fgeq(a, b *Float) *Boolean {
	if a.Value >= b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func fleq(a, b *Float) *Boolean {
	if a.Value <= b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func feq(a, b *Float) *Boolean {
	if a.Value == b.Value {
		return &Boolean{true}
	}
	return &Boolean{false}
}
func ftoi(a *Float) *Integer {
	return &Integer{int(a.Value)}
}
func ftos(a *Float) *String {
	return &String{fmt.Sprint(a.Value)}
}
func ftob(a *Float) *Boolean {
	if a.Value == 0.0 {
		return &Boolean{false}
	}
	return &Boolean{true}
}

// string functions

func sconcat(s1, s2 *String) *String {
	return &String{Value: s1.Value + s2.Value}
}

func slen(s1 *String) *Integer {
	return &Integer{Value: len(s1.Value)}
}

func stoi(s1 *String) *Integer {
	val, err := strconv.ParseInt(s1.Value, 10, 64)
	if err != nil {
		panic(err)
	}
	return &Integer{Value: int(val)}
}

func stof(s1 *String) *Float {
	val, err := strconv.ParseFloat(s1.Value, 64)
	if err != nil {
		panic(err)
	}
	return &Float{Value: val}
}

func stob(s1 *String) *Boolean {
	if s1.Value == "" {
		return &Boolean{Value: false}
	}
	return &Boolean{Value: true}
}
