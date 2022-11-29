package eval

import (
	"fmt"
	"reflect"
)

func Add(values []interface{}) (interface{}, error) {
	t := reflect.TypeOf(values[0])
	switch t.Kind() {
	case reflect.Int:
		return _add[int](values)
	case reflect.Float64:
		return _add[float64](values)
	case reflect.String:
		return _add[string](values)
	default:
		return nil, fmt.Errorf("unsupported type for symbol add")
	}
}
func _add[T int | float64 | string](values []interface{}) (T, error) {
	var res T
	for _, val := range values {
		res += val.(T)
	}
	return res, nil
}

func Sub(values []interface{}) (interface{}, error) {
	t := reflect.TypeOf(values[0])
	switch t.Kind() {
	case reflect.Int:
		return _sub[int](values)
	case reflect.Float64:
		return _sub[float64](values)
	default:
		return nil, fmt.Errorf("unsupported type for symbol add")
	}
}
func _sub[T int | float64](values []interface{}) (T, error) {
	var res T
	for i, val := range values {
		if i == 0 {
			res = val.(T)
			continue
		}
		res -= val.(T)
	}
	return res, nil
}

func Mul(values []interface{}) (interface{}, error) {
	t := reflect.TypeOf(values[0])
	switch t.Kind() {
	case reflect.Int:
		return _mul[int](values)
	case reflect.Float64:
		return _mul[float64](values)
	default:
		return nil, fmt.Errorf("unsupported type for symbol add")
	}
}
func _mul[T int | float64](values []interface{}) (T, error) {
	var res T
	for i, val := range values {
		if i == 0 {
			res = val.(T)
			continue
		}
		res *= val.(T)
	}
	return res, nil
}

func Div(values []interface{}) (interface{}, error) {
	t := reflect.TypeOf(values[0])
	switch t.Kind() {
	case reflect.Int:
		return _div[int](values)
	case reflect.Float64:
		return _div[float64](values)
	default:
		return nil, fmt.Errorf("unsupported type for symbol add")
	}
}
func _div[T int | float64](values []interface{}) (T, error) {
	var res T
	for i, val := range values {
		if i == 0 {
			res = val.(T)
			continue
		}
		res /= val.(T)
	}
	return res, nil
}
