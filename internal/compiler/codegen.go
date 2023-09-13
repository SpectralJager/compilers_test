package compiler

import (
	"log"
	"strings"
	"text/template"
)

var tmpls = map[string]map[string]string{
	"core": {
		"mainBody": `package main
import (
	"context"
	"log"
	"os"
)
{{range .globals}}
{{.}}
{{end}}

func main() {
	_main(ctx)
	os.Exit(0)
}
`,
		"fn": `func _{{.name}}(ctx context.Context,{{range .args}} {{.}},{{end}}) ({{range .returns}}{{.}},{{end}}) {
	{{range .locals}}{{.}}
	{{end}}
}`,
		"var":  `var {{.name}} {{.dtype}} = {{.value}}`,
		"set":  `{{.name}} = {{.value}}`,
		"atom": `{{.value}}`,
		"call": `{{.name}}(ctx,{{range .args}} {{.}},{{end}})`,
	},
	"buildin": {
		"iadd": `func iadd(ctx context.Context, args ...int) int{
    res := args[0]
    for _, v := range args[1:] {
        res += v
    }
    return res
}`,
		"isub": `func isub(ctx context.Context, args ...int) int{
    res := args[0]
    for _, v := range args[1:] {
        res -= v
    }
    return res
}`,
		"imul": `func imul(ctx context.Context, args ...int) int{
    res := args[0]
    for _, v := range args {
        res *= v
    }
    return res
}`,
		"idiv": `func idiv(ctx context.Context, args ...int32) int{
    res := args[0]
    for _, v := range args {
        res /= v
    }
    return res
}`,
	},
}

func Generate(node AST) string {
	var buf strings.Builder
	switch node := node.(type) {
	case *ProgramAST:
		var globals []string
		for _, gl := range node.Body {
			globals = append(globals, Generate(gl))
		}
		tmpl, err := template.New("").Parse(tmpls["core"]["mainBody"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"globals": globals,
		})
	case *FunctionAST:
		var locals []string
		for _, lc := range node.Body {
			locals = append(locals, Generate(lc))
		}
		tmpl, err := template.New("").Parse(tmpls["core"]["fn"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name":    node.Ident,
			"args":    nil,
			"returns": nil,
			"locals":  locals,
		})
	case *VaribleAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["var"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name":  node.Symbol.Ident,
			"dtype": "int",
			"value": Generate(node.Value),
		})
	case *SetAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["set"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name":  node.Ident,
			"value": Generate(node.Value),
		})
	case *SymbolExpressionAST:
		var args []string
		for _, arg := range node.Arguments {
			args = append(args, Generate(arg))
		}
		tmpl, err := template.New("").Parse(tmpls["core"]["call"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name": node.Symbol,
			"args": args,
		})
	case *IntegerAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["atom"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"value": node.Integer,
		})
	case *SymbolAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["atom"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"value": node.Symbol,
		})
	default:
		log.Fatalf("unexpected node type %T", node)
	}
	return buf.String()
}
