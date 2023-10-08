package ir

import (
	"fmt"
	"strings"
)

type Function struct {
	Name        string
	Args        map[string]string
	ReturnTypes map[string]struct{}
	Code        []Instruction
}

func NewFunction(name string) Function {
	return Function{
		Name: name,
	}
}

func (f *Function) MustLoadArg(name, tp string) {
	if err := f.LoadArg(name, tp); err != nil {
		panic(err)
	}
}

func (f *Function) LoadArg(name, tp string) error {
	if _, ok := f.Args[name]; ok {
		return fmt.Errorf("global %s already loaded", name)
	}
	f.Args[name] = tp
	return nil
}

func (f *Function) MustLoadRetTypes(types ...string) {
	if err := f.LoadRetTypes(types...); err != nil {
		panic(err)
	}
}

func (f *Function) LoadRetTypes(types ...string) error {
	for _, t := range types {
		if _, ok := f.ReturnTypes[t]; ok {
			return fmt.Errorf("type %s already loaded", t)
		}
		f.ReturnTypes[t] = struct{}{}
	}
	return nil
}

func (f *Function) LoadInstructions(ins ...Instruction) {
	f.Code = append(f.Code, ins...)
}

func (f *Function) String() string {
	var buf strings.Builder

	return buf.String()
}
