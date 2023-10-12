package ph1

import (
	"context"
	"fmt"
	"grimlang/internal/analysis"
	"grimlang/internal/ir"
	tp "grimlang/internal/type"
	"strings"
	"sync"
)

type (
	globals struct{}
)

func AnalyseModule(m *ir.ModuleIR) {
	glbs := make(analysis.SymbolTable, 0)
	for _, v := range analysis.Builtins {
		glbs.Add(v)
	}

	for _, v := range m.Functions {
		err := glbs.Add(
			&analysis.FunctionSymbol{
				Ident:       v.Name.Ident,
				Args:        v.Meta.Args,
				ReturnTypes: v.Meta.Returns,
			})
		if err != nil {
			panic(err)
		}
	}
	for _, inst := range m.Init {
		switch inst.Op {
		case ir.OP_VAR_NEW:
			sm := inst.Args[0].(*ir.SymbolIR).String()
			tp := inst.Args[1].(tp.Type)
			err := glbs.Add(
				&analysis.VariableSymbol{
					Ident: sm,
					Type:  tp,
					Scope: "global",
				},
			)
			if err != nil {
				panic(err)
			}
		}
	}

	ctx := context.WithValue(context.Background(), globals{}, &glbs)

	var wg sync.WaitGroup
	for _, f := range m.Functions {
		fn := f
		wg.Add(1)
		go func() {
			AnaliseFunction(ctx, fn)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(glbs)
}

func AnaliseFunction(ctx context.Context, f *ir.FunctionIR) {
	lcls := make(analysis.SymbolTable, 0)
	scopes := []string{"local"}

	for i, inst := range f.Code {
		switch inst.Op {
		case ir.OP_VAR_NEW:
			sm := inst.Args[0].(*ir.SymbolIR).String()
			tp := inst.Args[1].(tp.Type)
			err := lcls.Add(
				&analysis.VariableSymbol{
					Ident: sm,
					Type:  tp,
					Scope: scopes[len(scopes)-1],
				},
			)
			if err != nil {
				panic(err)
			}
		case ir.OP_LABEL:
			if strings.Contains(inst.Args[0].String(), "begin") {
				lbl := inst.Args[0].String()
				parts := strings.Split(lbl, "_")
				tok := parts[len(parts)-1]
				scopes = append(scopes, tok)
			} else if strings.Contains(inst.Args[0].String(), "end") {
				lbl := inst.Args[0].String()
				parts := strings.Split(lbl, "_")
				tok := parts[len(parts)-1]
				if scopes[len(scopes)-1] == tok {
					scopes = scopes[:len(scopes)-1]
				}
				code := make([]*ir.InstrIR, 0)
				for _, idnt := range lcls.GetVars(tok) {
					code = append(code, ir.VarFree(&ir.SymbolIR{Ident: idnt.(*analysis.VariableSymbol).Ident}))
				}
				f.Code = append(f.Code[:i+1], append(code, f.Code[i+1:]...)...)
			}
		}
	}
	if len(scopes) > 1 {
		panic(fmt.Errorf("unclosed blocks with tokens %v", scopes[1:]))
	}
	fmt.Println(lcls)
}
