package backend

import (
	"bytes"
	"fmt"
	"grimlang/internal/frontend"
	"reflect"
)

func GenerateCode(w *bytes.Buffer, node any) error {
	switch node := node.(type) {
	case frontend.Programm:
		w.WriteString("\npackage main\n")
		for _, val := range node.Body {
			err := GenerateCode(w, val)
			if err != nil {
				return err
			}
		}
	case *frontend.FunctionCom:
		w.WriteString(fmt.Sprintf("\n/* %s */\n", node.DocString))
		w.WriteString(fmt.Sprintf("func %s (", node.Symbol.Name))
		for _, val := range node.Args {
			w.WriteString(fmt.Sprintf("%s %s", val.Name, val.PrimitiveType))
		}
		w.WriteString(fmt.Sprintf(") %s {\n", node.Symbol.PrimitiveType))
		for _, val := range node.Body {
			err := GenerateCode(w, val)
			if err != nil {
				return err
			}
		}
		w.WriteString("}\n")
	case *frontend.LocalVaribleCom:

	default:
		t := reflect.TypeOf(node)
		return fmt.Errorf("unsupported type: %s", t)
	}
	return nil
}
