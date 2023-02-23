package frontend

import "fmt"

type Instruction interface {
	inst() string
}

type INST_LC struct {
	Value string
}

func (in *INST_LC) inst() string {
	return fmt.Sprintf("load %s", in.Value)
}

type INST_SCS struct {
	Symbol string
}

func (in *INST_SCS) inst() string {
	return fmt.Sprintf("save %s", in.Symbol)
}

type INST_LS struct {
	Symbol string
}

func (in *INST_LS) inst() string {
	return fmt.Sprintf("load %s", in.Symbol)
}

type INST_CALL struct {
	Symbol string
}

func (in *INST_CALL) inst() string {
	return fmt.Sprintf("call %s", in.Symbol)
}
