package eval

import (
	"fmt"
	"grimlang/runtime"
	"strings"
)

type EvalState interface {
	String() string

	GetBuiltinEnv() runtime.Environment
	InsertGlobalEnv(runtime.Environment)
	SearchGlobalEnv(string) runtime.Environment

	IsReturn() bool
	SetReturn(runtime.Litteral)
	GetReturn() runtime.Litteral
	GetReturnType() runtime.Type

	GetSymbolFlag() bool
	SetSymbolFlag()
	ClrSymbolFlag()

	IsSwitchEnv() bool
	SetSwitchEnv(runtime.Environment)
	GetSwitchEnv() runtime.Environment
}

type state struct {
	builtin   runtime.Environment
	envs      map[string]runtime.Environment
	ret       runtime.Litteral
	switchEnv runtime.Environment

	symbFlag bool
}

func (s *state) GetBuiltinEnv() runtime.Environment {
	return s.builtin
}

func (s *state) InsertGlobalEnv(env runtime.Environment) {
	if _, ok := s.envs[env.Name()]; ok {
		panic("can't insert Environment: Environment with such name already exists")
	}
	s.envs[env.Name()] = env
}

func (s *state) SearchGlobalEnv(name string) runtime.Environment {
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
	return fmt.Sprintf("-----< Environments\n%s%s", s.builtin.String(), strings.Join(envs, ""))
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

func (s *state) IsSwitchEnv() bool {
	return s.switchEnv != nil
}

func (s *state) SetSwitchEnv(env runtime.Environment) {
	s.switchEnv = env
}

func (s *state) GetSwitchEnv() runtime.Environment {
	tmp := s.switchEnv
	s.switchEnv = nil
	return tmp
}

func NewEvalState(builtin runtime.Environment) *state {
	return &state{
		envs:    map[string]runtime.Environment{},
		builtin: builtin,
	}
}
