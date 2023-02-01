package meta

import "grimlang/internal/core/frontend/tokens"

type Metadata interface {
	Type() string
}

// --------------------------
type MetaVarible struct {
	DataType tokens.Token
	Scope    string
}

func (meta *MetaVarible) Type() string {
	return "metavarible"
}

// --------------------------
type MetaFunction struct {
	ReturnType tokens.Token
}

func (meta *MetaFunction) Type() string {
	return "metafunction"
}

// --------------------------
type MetaBlock struct {
	ScopeName string
}

func (meta *MetaBlock) Type() string {
	return "metafunction"
}
