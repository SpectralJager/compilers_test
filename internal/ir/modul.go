package ir

import "fmt"

type Modul struct {
	Name      string
	Types     map[string]struct{}
	Globals   map[string]string
	Init      []Instruction
	Functions map[string]Function
}

func NewModul(name string) *Modul {
	return &Modul{Name: name}
}

func (m *Modul) MustLoadTypes(types ...string) {
	if err := m.LoadTypes(types...); err != nil {
		panic(err)
	}
}

func (m *Modul) LoadTypes(types ...string) error {
	for _, t := range types {
		if _, ok := m.Types[t]; ok {
			return fmt.Errorf("type %s already loaded", t)
		}
		m.Types[t] = struct{}{}
	}
	return nil
}

func (m *Modul) MustLoadGlobal(name, tp string) {
	if err := m.LoadGlobal(name, tp); err != nil {
		panic(err)
	}
}

func (m *Modul) LoadGlobal(name, tp string) error {
	if _, ok := m.Globals[name]; ok {
		return fmt.Errorf("global %s already loaded", name)
	}
	m.Globals[name] = tp
	return nil
}

func (m *Modul) LoadInstructions(ins ...Instruction) {
	m.Init = append(m.Init, ins...)
}

func (m *Modul) MustAddFunctions(fns ...Function) {
	if err := m.LoadFunctions(fns...); err != nil {
		panic(err)
	}
}

func (m *Modul) LoadFunctions(fns ...Function) error {
	for _, f := range fns {
		if _, ok := m.Functions[f.Name]; ok {
			return fmt.Errorf("function %s already loaded", f.Name)
		}
		m.Functions[f.Name] = f
	}
	return nil

}
