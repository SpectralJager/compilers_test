package semantic

type Symbol struct {
}

type Scope interface {
	scope()
	AddSymbol(*Symbol) error
	RemoveSymbol(string) error
	GetSymbol(string) (*Symbol, error)
}

type ScopeStack struct {
	Scopes []*Scope
}

func (s *ScopeStack) Push(sc *Scope) {
	s.Scopes = append(s.Scopes, sc)
}

func (s *ScopeStack) Pop() *Scope {
	if len(s.Scopes) == 0 {
		return nil
	}
	sc := s.Scopes[len(s.Scopes)-1]
	s.Scopes = s.Scopes[:len(s.Scopes)-1]
	return sc
}

func (s *ScopeStack) Top() *Scope {
	return s.Scopes[len(s.Scopes)-1]
}

func (s *ScopeStack) Def(sb Symbol) Scope {
	scope := s.Top()
	scope.AddSymbol(&sb)
	return scope
}
