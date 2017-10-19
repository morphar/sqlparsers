package generator

var InsertSQLs []string
var InsertQueries []string

func init() {
	InsertSQLs = append(InsertSQLs, "INSERT INTO users(id, \"name\") VALUES (1, 'joe')")

	InsertSQLs = append(InsertSQLs, "INSERT INTO users(\"name\") VALUES ('joe') RETURNING id")

	InsertQueries = append(InsertQueries, `&parser.Insert{
	Table: &parser.NormalizableTableName{
		TableNameReference: parser.UnresolvedName{
			parser.Name("users"),
		},
	},
	Columns: parser.UnresolvedNames{
		parser.UnresolvedName{
			parser.Name("id"),
		},

		parser.UnresolvedName{
			parser.Name("name"),
		},
	},
	Rows: &parser.Select{
		Select: &parser.ValuesClause{
			Tuples: []*parser.Tuple{
				&parser.Tuple{
					Exprs: parser.Exprs{
						&parser.NumVal{
							Value:      constant.int64Val,
							OrigString: "1",
						},

						parser.NewDString("joe"),
					},
				},
			},
		},
		OrderBy: nil,
		Limit:   nil,
	},
	OnConflict: nil,
	Returning:  &parser.NoReturningClause{},
}`)

	InsertQueries = append(InsertQueries, `&parser.Insert{
	Table: &parser.NormalizableTableName{
		TableNameReference: parser.UnresolvedName{
			parser.Name("users"),
		},
	},
	Columns: parser.UnresolvedNames{
		parser.UnresolvedName{
			parser.Name("name"),
		},
	},
	Rows: &parser.Select{
		Select: &parser.ValuesClause{
			Tuples: []*parser.Tuple{
				&parser.Tuple{
					Exprs: parser.Exprs{
						parser.NewDString("joe"),
					},
				},
			},
		},
		OrderBy: nil,
		Limit:   nil,
	},
	OnConflict: nil,
	Returning: &parser.ReturningExprs{
		parser.SelectExpr{
			Expr: parser.UnresolvedName{
				parser.Name("id"),
			},
			As: parser.Name(""),
		},
	},
}`)

}
