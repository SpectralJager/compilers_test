package symboltable

import "fmt"

type SymbolTable struct {
	Table map[string]interface{}
}

func (s *SymbolTable) Get(symbol string) (interface{}, error) {
	val, ok := s.Table[symbol]
	if !ok {
		return val, fmt.Errorf("cant find symbol %s in symbol table", symbol)
	}
	return val, nil
}
func (s *SymbolTable) Save(symbol string, value interface{}) error {
	s.Table[symbol] = value
	return nil
}

func (s *SymbolTable) String() string {
	res := "SymbolTable \n"
	for k, v := range s.Table {
		res += fmt.Sprintf("|%12s|%20v|\n", k, v)
	}
	res += "\n"
	return res
}

var symboltable = SymbolTable{
	Table: make(map[string]interface{}),
}

func GetSymbolTableEntity() *SymbolTable {
	return &symboltable
}
