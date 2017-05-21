// Command qbcli takes 2 arguments: an SQL string to parse into Go code in the
// format that the sql/parser uses and a boolean indicating if the parsed query
// should also be printed as SQL.
// DISCLAIMER: This is a work in progress, currently only tested with a very
// narrow set of queries.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	mysqlGenerator "github.com/morphar/sqlparsers/generators/mysql"
	postgresGenerator "github.com/morphar/sqlparsers/generators/postgres"
	mysqlParser "github.com/morphar/sqlparsers/mysql"
	postgresParser "github.com/morphar/sqlparsers/postgres"
)

var (
	syntax   string = "postgres" // mysql or postgres
	sql      string = ""
	printSQL bool   = false
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	flag.StringVar(&syntax, "syntax", syntax, "postgres or mysql syntax. Default is postgres")
	flag.BoolVar(&printSQL, "print", printSQL, "Print the generated SQL after it has been parsed?")
	flag.StringVar(&sql, "sql", sql, "Required. The SQL to parse.")
}

func main() {
	flag.Parse()

	if sql == "" && len(os.Args) != 2 {
		fmt.Println("\nPlease provide an SQL to parse")
		fmt.Println(`Example: sqlparsers -print -sql "SELECT * FROM users"`)
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
		fmt.Println("")
		return
	}

	// TODO: Is this right?
	if sql == "" && len(os.Args) == 2 {
		sql = os.Args[1]
	}

	if syntax == "postgres" {
		stmt, err := postgresParser.Parse(sql)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(postgresGenerator.Parse(stmt[0]))

		if printSQL || true == true {
			fmt.Println("------------------------------- GENERATED QUERY -------------------------------------")
			fmt.Println(stmt.String())
			return
		}

	} else if syntax == "mysql" {
		stmt, err := mysqlParser.Parse(sql)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(mysqlGenerator.Parse(stmt))

		if printSQL || true == true {
			pq := mysqlParser.GenerateParsedQuery(stmt)
			fmt.Println("------------------------------- GENERATED QUERY -------------------------------------")
			fmt.Println(pq.Query)
			return
		}
	}

}
