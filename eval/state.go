package eval

import (
	"fmt"
	"grimlang/runtime"
	"strings"
)

type EvalState interface {
	String() string
	GetBuiltinEnv() runtime.Enviroment
	InsertGlobalEnv(runtime.Enviroment)
	SearchGlobalEnv(string) runtime.Enviroment
	IsReturn() bool
	SetReturn(runtime.Litteral)
	GetReturn() runtime.Litteral
	GetReturnType() runtime.Type
	GetSymbolFlag() bool
	SetSymbolFlag()
	ClrSymbolFlag()
}

type state struct {
	builtin runtime.Enviroment
	envs    map[string]runtime.Enviroment
	ret     runtime.Litteral

	symbFlag bool
}

func (s *state) GetBuiltinEnv() runtime.Enviroment {
	return s.builtin
}

func (s *state) InsertGlobalEnv(env runtime.Enviroment) {
	if _, ok := s.envs[env.Name()]; ok {
		panic("can't insert enviroment: enviroment with such name already exists")
	}
	s.envs[env.Name()] = env
}

func (s *state) SearchGlobalEnv(name string) runtime.Enviroment {
	return s.envs[name]
}

func (s *state) IsReturn() bool {
	return s.ret != nil
}

func (s *state) SetReturn(value runtime.Litteral) {
	s.ret = value
}

func (s *state) GetReturn() runtime.Litteral {
	tmp := s.ret
	s.ret = nil
	return tmp
}

func (s *state) GetReturnType() runtime.Type {
	if s.ret != nil {
		return s.ret.Type()
	}
	return runtime.NewVoidType()
}

func (s *state) String() string {
	envs := []string{}
	for _, env := range s.envs {
		envs = append(envs, env.String())
	}
	return fmt.Sprintf("-----< Enviroments\n%s%s", s.builtin.String(), strings.Join(envs, ""))
}

func (s *state) GetSymbolFlag() bool {
	return s.symbFlag
}
func (s *state) SetSymbolFlag() {
	s.symbFlag = true
}
func (s *state) ClrSymbolFlag() {
	s.symbFlag = false
}

func NewEvalState(builtin runtime.Enviroment) *state {
	return &state{
		envs:    map[string]runtime.Enviroment{},
		builtin: builtin,
	}
}
