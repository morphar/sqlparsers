package generator

var SelectSQLs []string
var SelectQueries []string

func init() {
	SelectSQLs = append(SelectSQLs, "SELECT id, \"name\" "+
		"FROM users WHERE \"name\" IN ('joe', 'juice')")

	SelectSQLs = append(SelectSQLs, "SELECT posts.*, users.\"name\" "+
		"FROM posts LEFT JOIN users ON users.id = posts.users_id")

	SelectSQLs = append(SelectSQLs, "SELECT ideas.id, ideas.title, "+
		"CASE votes.users_id WHEN '' THEN false ELSE true END AS i_like "+
		"FROM ideas "+
		"LEFT JOIN votes ON (votes.ideas_id = ideas.id) AND (votes.users_id = '538dc3820f942a4767ffcd75')")

	SelectSQLs = append(SelectSQLs, "SELECT CASE WHEN votes.users_id IS NULL THEN false ELSE true END AS i_like "+
		"FROM ideas LEFT JOIN votes ON (votes.ideas_id = ideas.id) AND (votes.users_id = '538dc3820f942a4767ffcd75')")

	SelectSQLs = append(SelectSQLs, "SELECT author.email, author.fname, author.id, author.lname, "+
		"posts.content, posts.id, posts.title, "+
		"status.id, status.\"name\" "+
		"FROM posts "+
		"JOIN users AS author ON author.id = posts.users_id "+
		"LEFT JOIN status ON status.id = posts.status_id "+
		"WHERE (author.fname LIKE $1) AND (author.id IN ($2, $3)) "+
		"ORDER BY author.fname DESC, author.lname DESC "+
		"LIMIT 5 OFFSET 10")

	SelectQueries = append(SelectQueries, `&parser.Select{
	Select: &parser.SelectClause{
		Distinct: false,
		Exprs: parser.SelectExprs{
			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("id"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("name"),
				},
				As: parser.Name(""),
			},
		},
		From: &parser.From{
			Tables: parser.TableExprs{
				&parser.AliasedTableExpr{
					Expr: &parser.NormalizableTableName{
						TableNameReference: parser.UnresolvedName{
							parser.Name("users"),
						},
					},
					Hints:      nil,
					Ordinality: false,
					As: parser.AliasClause{
						Alias: parser.Name(""),
						Cols:  nil,
					},
				},
			},
			AsOf: parser.AsOfClause{
				Expr: nil,
			},
		},
		Where: &parser.Where{
			Type: "WHERE",
			Expr: &parser.ComparisonExpr{
				Operator:    6,
				SubOperator: 0,
				Left: parser.UnresolvedName{
					parser.Name("name"),
				},
				Right: &parser.Tuple{
					Exprs: parser.Exprs{
						parser.NewDString("joe"),

						parser.NewDString("juice"),
					},
				},
			},
		},
		GroupBy: nil,
		Having:  nil,
		Window:  nil,
	},
	OrderBy: nil,
	Limit:   nil,
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Select: &parser.SelectClause{
		Distinct: false,
		Exprs: parser.SelectExprs{
			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("posts"),

					parser.UnqualifiedStar{},
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("users"),

					parser.Name("name"),
				},
				As: parser.Name(""),
			},
		},
		From: &parser.From{
			Tables: parser.TableExprs{
				&parser.JoinTableExpr{
					Join: "LEFT JOIN",
					Left: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("posts"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Right: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("users"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Cond: &parser.OnJoinCond{
						Expr: &parser.ComparisonExpr{
							Operator:    0,
							SubOperator: 0,
							Left: parser.UnresolvedName{
								parser.Name("users"),

								parser.Name("id"),
							},
							Right: parser.UnresolvedName{
								parser.Name("posts"),

								parser.Name("users_id"),
							},
						},
					},
				},
			},
			AsOf: parser.AsOfClause{
				Expr: nil,
			},
		},
		Where:   nil,
		GroupBy: nil,
		Having:  nil,
		Window:  nil,
	},
	OrderBy: nil,
	Limit:   nil,
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Select: &parser.SelectClause{
		Distinct: false,
		Exprs: parser.SelectExprs{
			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("ideas"),

					parser.Name("id"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("ideas"),

					parser.Name("title"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: &parser.CaseExpr{
					Expr: parser.UnresolvedName{
						parser.Name("votes"),

						parser.Name("users_id"),
					},
					Whens: []*parser.When{
						&parser.When{
							Cond: parser.NewDString(""),
							Val:  parser.DBoolFalse,
						},
					},
					Else: parser.DBoolTrue,
				},
				As: parser.Name("i_like"),
			},
		},
		From: &parser.From{
			Tables: parser.TableExprs{
				&parser.JoinTableExpr{
					Join: "LEFT JOIN",
					Left: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("ideas"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Right: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("votes"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Cond: &parser.OnJoinCond{
						Expr: &parser.AndExpr{
							Left: &parser.ParenExpr{
								Expr: &parser.ComparisonExpr{
									Operator:    0,
									SubOperator: 0,
									Left: parser.UnresolvedName{
										parser.Name("votes"),

										parser.Name("ideas_id"),
									},
									Right: parser.UnresolvedName{
										parser.Name("ideas"),

										parser.Name("id"),
									},
								},
							},
							Right: &parser.ParenExpr{
								Expr: &parser.ComparisonExpr{
									Operator:    0,
									SubOperator: 0,
									Left: parser.UnresolvedName{
										parser.Name("votes"),

										parser.Name("users_id"),
									},
									Right: parser.NewDString("538dc3820f942a4767ffcd75"),
								},
							},
						},
					},
				},
			},
			AsOf: parser.AsOfClause{
				Expr: nil,
			},
		},
		Where:   nil,
		GroupBy: nil,
		Having:  nil,
		Window:  nil,
	},
	OrderBy: nil,
	Limit:   nil,
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Select: &parser.SelectClause{
		Distinct: false,
		Exprs: parser.SelectExprs{
			parser.SelectExpr{
				Expr: &parser.CaseExpr{
					Expr: nil,
					Whens: []*parser.When{
						&parser.When{
							Cond: &parser.ComparisonExpr{
								Operator:    20,
								SubOperator: 0,
								Left: parser.UnresolvedName{
									parser.Name("votes"),

									parser.Name("users_id"),
								},
								Right: parser.DNull,
							},
							Val: parser.DBoolFalse,
						},
					},
					Else: parser.DBoolTrue,
				},
				As: parser.Name("i_like"),
			},
		},
		From: &parser.From{
			Tables: parser.TableExprs{
				&parser.JoinTableExpr{
					Join: "LEFT JOIN",
					Left: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("ideas"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Right: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("votes"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Cond: &parser.OnJoinCond{
						Expr: &parser.AndExpr{
							Left: &parser.ParenExpr{
								Expr: &parser.ComparisonExpr{
									Operator:    0,
									SubOperator: 0,
									Left: parser.UnresolvedName{
										parser.Name("votes"),

										parser.Name("ideas_id"),
									},
									Right: parser.UnresolvedName{
										parser.Name("ideas"),

										parser.Name("id"),
									},
								},
							},
							Right: &parser.ParenExpr{
								Expr: &parser.ComparisonExpr{
									Operator:    0,
									SubOperator: 0,
									Left: parser.UnresolvedName{
										parser.Name("votes"),

										parser.Name("users_id"),
									},
									Right: parser.NewDString("538dc3820f942a4767ffcd75"),
								},
							},
						},
					},
				},
			},
			AsOf: parser.AsOfClause{
				Expr: nil,
			},
		},
		Where:   nil,
		GroupBy: nil,
		Having:  nil,
		Window:  nil,
	},
	OrderBy: nil,
	Limit:   nil,
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Select: &parser.SelectClause{
		Distinct: false,
		Exprs: parser.SelectExprs{
			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("author"),

					parser.Name("email"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("author"),

					parser.Name("fname"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("author"),

					parser.Name("id"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("author"),

					parser.Name("lname"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("posts"),

					parser.Name("content"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("posts"),

					parser.Name("id"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("posts"),

					parser.Name("title"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("status"),

					parser.Name("id"),
				},
				As: parser.Name(""),
			},

			parser.SelectExpr{
				Expr: parser.UnresolvedName{
					parser.Name("status"),

					parser.Name("name"),
				},
				As: parser.Name(""),
			},
		},
		From: &parser.From{
			Tables: parser.TableExprs{
				&parser.JoinTableExpr{
					Join: "LEFT JOIN",
					Left: &parser.JoinTableExpr{
						Join: "JOIN",
						Left: &parser.AliasedTableExpr{
							Expr: &parser.NormalizableTableName{
								TableNameReference: parser.UnresolvedName{
									parser.Name("posts"),
								},
							},
							Hints:      nil,
							Ordinality: false,
							As: parser.AliasClause{
								Alias: parser.Name(""),
								Cols:  nil,
							},
						},
						Right: &parser.AliasedTableExpr{
							Expr: &parser.NormalizableTableName{
								TableNameReference: parser.UnresolvedName{
									parser.Name("users"),
								},
							},
							Hints:      nil,
							Ordinality: false,
							As: parser.AliasClause{
								Alias: parser.Name("author"),
								Cols:  nil,
							},
						},
						Cond: &parser.OnJoinCond{
							Expr: &parser.ComparisonExpr{
								Operator:    0,
								SubOperator: 0,
								Left: parser.UnresolvedName{
									parser.Name("author"),

									parser.Name("id"),
								},
								Right: parser.UnresolvedName{
									parser.Name("posts"),

									parser.Name("users_id"),
								},
							},
						},
					},
					Right: &parser.AliasedTableExpr{
						Expr: &parser.NormalizableTableName{
							TableNameReference: parser.UnresolvedName{
								parser.Name("status"),
							},
						},
						Hints:      nil,
						Ordinality: false,
						As: parser.AliasClause{
							Alias: parser.Name(""),
							Cols:  nil,
						},
					},
					Cond: &parser.OnJoinCond{
						Expr: &parser.ComparisonExpr{
							Operator:    0,
							SubOperator: 0,
							Left: parser.UnresolvedName{
								parser.Name("status"),

								parser.Name("id"),
							},
							Right: parser.UnresolvedName{
								parser.Name("posts"),

								parser.Name("status_id"),
							},
						},
					},
				},
			},
			AsOf: parser.AsOfClause{
				Expr: nil,
			},
		},
		Where: &parser.Where{
			Type: "WHERE",
			Expr: &parser.AndExpr{
				Left: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator:    8,
						SubOperator: 0,
						Left: parser.UnresolvedName{
							parser.Name("author"),

							parser.Name("fname"),
						},
						Right: &parser.Placeholder{
							Name: "1",
						},
					},
				},
				Right: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator:    6,
						SubOperator: 0,
						Left: parser.UnresolvedName{
							parser.Name("author"),

							parser.Name("id"),
						},
						Right: &parser.Tuple{
							Exprs: parser.Exprs{
								&parser.Placeholder{
									Name: "2",
								},

								&parser.Placeholder{
									Name: "3",
								},
							},
						},
					},
				},
			},
		},
		GroupBy: nil,
		Having:  nil,
		Window:  nil,
	},
	OrderBy: parser.OrderBy{
		&parser.Order{
			Expr: parser.UnresolvedName{
				parser.Name("author"),

				parser.Name("fname"),
			},
			Direction: 2,
		},

		&parser.Order{
			Expr: parser.UnresolvedName{
				parser.Name("author"),

				parser.Name("lname"),
			},
			Direction: 2,
		},
	},
	Limit: &parser.Limit{
		Offset: &parser.NumVal{
			Value:      constant.int64Val,
			OrigString: "10",
		},
		Count: &parser.NumVal{
			Value:      constant.int64Val,
			OrigString: "5",
		},
	},
}`)

}
