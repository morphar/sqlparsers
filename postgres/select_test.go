package parser

import (
	"testing"
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
		stmt, err := Parse(SelectSQL)
		if err != nil {
			t.Errorf("\nCaught an error from parser: %q", err)
		}

		// Check that we can re-generate the original selects
		generatedSQL := stmt.String()
		if SelectSQL != generatedSQL {
			t.Errorf(errStr, i, generatedSQL, i, SelectSQL)
		}

	}
}
