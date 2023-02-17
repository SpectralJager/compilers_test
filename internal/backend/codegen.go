package backend

import (
	"bytes"
	"grimlang/internal/frontend"
)

func GenerateProgram(w *bytes.Buffer, programm *frontend.Programm) {
	programm.Package = "main"
}
