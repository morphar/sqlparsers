package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
	"github.com/morphar/sqlparsers/pkg/tsql/tsql"
	"github.com/stretchr/testify/require"
)

// var Parser = participle.MustBuild(tsql.Lexer, participle.CaseInsensitive())

var alphaNumeric = regexp.MustCompile(`[a-zA-Z0-9_\-]`)

var caseInsensitiveTokens []string

func init() {
	var todos []struct {
		key string
		i   int
	}
	for key, g := range tsql.Rules {
		for i, r := range g {
			if alphaNumeric.MatchString(r.Pattern) || inStr(r.Pattern, "PROCEDURE", "TRANSACTION", "EXECUTE") {
				todos = append(todos, struct {
					key string
					i   int
				}{key, i})
				caseInsensitiveTokens = append(caseInsensitiveTokens, r.Name)
			}
		}
	}
	for _, t := range todos {
		tsql.Rules[t.key][t.i].Pattern = "(?i)" + tsql.Rules[t.key][t.i].Pattern
	}
}

func TestExe(t *testing.T) {
	matches, err := filepath.Glob("testdata/*.sql")
	require.NoError(t, err)

	trace, err := os.Create("parse.trace")
	require.NoError(t, err)
	defer trace.Close()

	var Lexer = lexer.MustStateful(tsql.Rules, lexer.MatchLongest())

	var Parser = participle.MustBuild(
		&tsql.Script{},
		participle.Lexer(Lexer),
		participle.UseLookahead(200),
		participle.Trace(trace),
		participle.CaseInsensitive(caseInsensitiveTokens...),
	)

	for _, m := range matches {
		t.Run(m, func(t *testing.T) {
			trace.Write([]byte(m))
			r, err := os.Open(m)
			require.NoError(t, err)

			res := &tsql.Script{}

			buf := bytes.NewBuffer(nil)
			if _, err := io.Copy(buf, r); err != nil {
				t.Fatal(err)
			}
			rdr := bytes.NewReader(buf.Bytes())

			tokens, err := Parser.Lex(m, rdr)
			twts := makeTwts(tokens, Lexer)
			_ = twts
			if err != nil {
				t.Logf(repr.String(twts))
				t.Fatal(err)
			}

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%v", r)
					t.Logf(repr.String(twts))
					repr.Print(res)
				}
			}()

			rdr.Seek(0, 0)
			err = Parser.Parse(m, rdr, res)
			if err != nil {
				t.Logf(repr.String(twts))
				repr.Print(res)
			}
			require.NoError(t, err)

			err = r.Close()
			require.NoError(t, err)
		})
	}
}

type tokenWithType struct {
	tok  lexer.Token
	name string
}

func (twt tokenWithType) GoString() string {
	return fmt.Sprintf("\n"+`Token@%s{"%s", "%s"}`, twt.tok.Pos.String(), twt.name, twt.tok.Value)
}

type tokenWithTypes []tokenWithType

func makeTwts(toks []lexer.Token, l lexer.Definition) (ret tokenWithTypes) {
	tokNames := lexer.SymbolsByRune(l)
	ret = make([]tokenWithType, len(toks))
	for i, t := range toks {
		ret[i] = tokenWithType{t, tokNames[t.Type]}
	}
	return
}

func inStr(n string, h ...string) bool {
	for _, s := range h {
		if n == s {
			return true
		}
	}
	return false
}
