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
		"arg":  `{{.name}} {{.dtype}}`,
		"var":  `var {{.name}} {{.dtype}} = {{.value}}`,
		"set":  `{{.name}} = {{.value}}`,
		"atom": `{{.value}}`,
		"call": `{{.name}}(ctx,{{range .args}} {{.}},{{end}})`,
		"if": `switch {
	case {{.expr}}:
		{{range .then}}{{.}}
		{{end}}
	{{if .else}}default:
		{{range .else}}{{.}}
		{{end}}{{end}}
	}`,
		"return": `return {{.value}}`,
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
		var locals, args []string
		for _, arg := range node.Arguments {
			args = append(args, Generate(arg))
		}
		for _, lc := range node.Body {
			locals = append(locals, Generate(lc))
		}
		tmpl, err := template.New("").Parse(tmpls["core"]["fn"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name":    node.Ident,
			"args":    args,
			"returns": nil,
			"locals":  locals,
		})
	case SymbolDeclAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["arg"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"name":  node.Ident,
			"dtype": "int",
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
	case *ReturnAST:
		tmpl, err := template.New("").Parse(tmpls["core"]["return"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"value": Generate(node.Expression),
		})
	case *IfAST:
		var th, el []string
		for _, lc := range node.ThenBody {
			th = append(th, Generate(lc))
		}
		for _, lc := range node.ElseBody {
			el = append(el, Generate(lc))
		}
		tmpl, err := template.New("").Parse(tmpls["core"]["if"])
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(&buf, map[string]any{
			"expr": Generate(node.Expr),
			"then": th,
			"else": el,
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
