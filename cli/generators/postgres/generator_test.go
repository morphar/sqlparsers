package generator

import (
	"testing"

	parser "github.com/morphar/sqlparsers/postgres"
)

var (
	errStr string = `
-------------------------------- From Generator (%d) --------------------------------

%s

----------------------------------- Expected (%d) -----------------------------------

%s

-------------------------------------------------------------------------------------
-------------------------------------------------------------------------------------

`
)

func TestCheckParse(t *testing.T) {
	for i, SelectSQL := range SelectSQLs {
		stmtList, err := parser.Parse(SelectSQL)
		if err != nil || len(stmtList) == 0 {
			t.Errorf("\nCaught an error from parser: %q", err)
		}
		stmt := stmtList[0]

		// Check that we can parse out to an expected format
		parsedQuery := Parse(stmt)
		if parsedQuery != SelectQueries[i] {
			t.Errorf(errStr, i, parsedQuery, i, SelectQueries[i])
		}

		// Check that we can re-generate the original selects
		generatedSQL := stmt.String()
		if SelectSQL != generatedSQL {
			t.Errorf(errStr, i, generatedSQL, i, SelectSQL)
		}

	}
}

//
// Doesn't really matter, as this is a helper tool ;)
//

func BenchmarkParseSmallSelectQuery(b *testing.B) {
	stmt, err := parser.Parse(SelectSQLs[0])
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(stmt[0])
	}
}

func BenchmarkParseSmallSelectSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stmt, err := parser.Parse(SelectSQLs[0])
		if err != nil {
			panic(err)
		}
		Parse(stmt[0])
	}
}

func BenchmarkParseLargeSelectQuery(b *testing.B) {
	stmt, err := parser.Parse(SelectSQLs[2])
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(stmt[0])
	}
}

func BenchmarkParseLargeSelectSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stmt, err := parser.Parse(SelectSQLs[2])
		if err != nil {
			panic(err)
		}
		Parse(stmt[0])
	}
}
