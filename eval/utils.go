package eval

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"grimlang/ast"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

func CreateModuleFromFile(path string) (*ast.Module, string) {
	data, err := os.ReadFile(path)
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
