package antlr

import (
	"io"

	"github.com/morphar/sqlparsers/internal/antlr/ast"
)

// Parse generates a walkable AST from an Antlr grammar file.
func Parse(filename string, r io.Reader) (dst *ast.AntlrFile, err error) {
	dst = &ast.AntlrFile{}
	err = ast.Parser.Parse(filename, r, dst)
	return
}
