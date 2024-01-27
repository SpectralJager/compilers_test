package eval

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"grimlang/ast"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/participle/v2"
)

func CreateModuleFromFile(path string) (*ast.Module, string) {
	file := filepath.Base(path)
	dir := filepath.Dir(path)
	err := os.Chdir(dir)
	// fmt.Println(os.Getwd())
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var errBuf strings.Builder
	module, err := ast.Parser.ParseBytes("", data, participle.Trace(&errBuf))
	if err != nil {
		fmt.Println(errBuf.String())
		panic(err)
	}
	hash := md5.Sum(data)
	return module, hex.EncodeToString(hash[:])
}
