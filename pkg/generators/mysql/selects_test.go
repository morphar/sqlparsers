package generator

var SelectSQLs []string
var SelectQueries []string

func init() {
	SelectSQLs = append(SelectSQLs, "select id, 'name' "+
		"from users where 'name' in ('joe', 'juice')")

	SelectSQLs = append(SelectSQLs, "select posts.*, users.name "+
		"from posts left join users on users.id = posts.users_id")

	SelectSQLs = append(SelectSQLs, "select ideas.id, ideas.title, "+
		"case votes.users_id when '' then false else true end as i_like "+
		"from ideas "+
		"left join votes on (votes.ideas_id = ideas.id) and (votes.users_id = '538dc3820f942a4767ffcd75')")

	SelectSQLs = append(SelectSQLs, "select case when votes.users_id is null then false else true end as i_like "+
		"from ideas left join votes on (votes.ideas_id = ideas.id) and (votes.users_id = '538dc3820f942a4767ffcd75')")

	SelectSQLs = append(SelectSQLs, "select author.email, author.fname, author.id, author.lname, "+
		"posts.content, posts.id, posts.title, "+
		"status.id, status.name "+
		"from posts "+
		"join users as author on author.id = posts.users_id "+
		"left join status on status.id = posts.status_id "+
		"where (author.fname like :v1) and (author.id in (:v2, :v3)) "+
		"order by author.fname desc, author.lname desc "+
		"limit 10, 5")

	SelectQueries = append(SelectQueries, `&parser.Select{
	Cache:    "",
	Comments: nil,
	Distinct: "",
	Hints:    "",
	SelectExprs: parser.SelectExprs{
		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("id"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent(""),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.SQLVal{
				Type: 0,
				Val:  []byte("name"),
			},
			As: parser.NewColIdent(""),
		},
	},
	From: parser.TableExprs{
		&parser.AliasedTableExpr{
			Expr: parser.TableName{
				Name:      parser.NewTableIdent("users"),
				Qualifier: parser.NewTableIdent(""),
			},
			As:    parser.NewTableIdent(""),
			Hints: nil,
		},
	},
	Where: &parser.Where{
		Type: "where",
		Expr: &parser.ComparisonExpr{
			Operator: "in",
			Left: &parser.SQLVal{
				Type: 0,
				Val:  []byte("name"),
			},
			Right: parser.ValTuple{
				&parser.SQLVal{
					Type: 0,
					Val:  []byte("joe"),
				},

				&parser.SQLVal{
					Type: 0,
					Val:  []byte("juice"),
				},
			},
			Escape: nil,
		},
	},
	GroupBy: nil,
	Having:  nil,
	OrderBy: nil,
	Limit:   nil,
	Lock:    "",
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Cache:    "",
	Comments: nil,
	Distinct: "",
	Hints:    "",
	SelectExprs: parser.SelectExprs{
		&parser.StarExpr{
			TableName: parser.TableName{
				Name:      parser.NewTableIdent("posts"),
				Qualifier: parser.NewTableIdent(""),
			},
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("name"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("users"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},
	},
	From: parser.TableExprs{
		&parser.JoinTableExpr{
			LeftExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("posts"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			Join: "left join",
			RightExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("users"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			On: &parser.ComparisonExpr{
				Operator: "=",
				Left: &parser.ColName{
					Metadata: nil,
					Name:     parser.NewColIdent("id"),
					Qualifier: parser.TableName{
						Name:      parser.NewTableIdent("users"),
						Qualifier: parser.NewTableIdent(""),
					},
				},
				Right: &parser.ColName{
					Metadata: nil,
					Name:     parser.NewColIdent("users_id"),
					Qualifier: parser.TableName{
						Name:      parser.NewTableIdent("posts"),
						Qualifier: parser.NewTableIdent(""),
					},
				},
				Escape: nil,
			},
		},
	},
	Where:   nil,
	GroupBy: nil,
	Having:  nil,
	OrderBy: nil,
	Limit:   nil,
	Lock:    "",
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Cache:    "",
	Comments: nil,
	Distinct: "",
	Hints:    "",
	SelectExprs: parser.SelectExprs{
		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("id"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("ideas"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("title"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("ideas"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.CaseExpr{
				Expr: &parser.ColName{
					Metadata: nil,
					Name:     parser.NewColIdent("users_id"),
					Qualifier: parser.TableName{
						Name:      parser.NewTableIdent("votes"),
						Qualifier: parser.NewTableIdent(""),
					},
				},
				Whens: []*parser.When{
					&parser.When{
						Cond: &parser.SQLVal{
							Type: 0,
							Val:  []byte(""),
						},
						Val: parser.BoolVal(false),
					},
				},
				Else: parser.BoolVal(true),
			},
			As: parser.NewColIdent("i_like"),
		},
	},
	From: parser.TableExprs{
		&parser.JoinTableExpr{
			LeftExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("ideas"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			Join: "left join",
			RightExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("votes"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			On: &parser.AndExpr{
				Left: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator: "=",
						Left: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("ideas_id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("votes"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Right: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("ideas"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Escape: nil,
					},
				},
				Right: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator: "=",
						Left: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("users_id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("votes"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Right: &parser.SQLVal{
							Type: 0,
							Val:  []byte("538dc3820f942a4767ffcd75"),
						},
						Escape: nil,
					},
				},
			},
		},
	},
	Where:   nil,
	GroupBy: nil,
	Having:  nil,
	OrderBy: nil,
	Limit:   nil,
	Lock:    "",
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Cache:    "",
	Comments: nil,
	Distinct: "",
	Hints:    "",
	SelectExprs: parser.SelectExprs{
		&parser.AliasedExpr{
			Expr: &parser.CaseExpr{
				Expr: nil,
				Whens: []*parser.When{
					&parser.When{
						Cond: &parser.IsExpr{
							Operator: "is null",
							Expr: &parser.ColName{
								Metadata: nil,
								Name:     parser.NewColIdent("users_id"),
								Qualifier: parser.TableName{
									Name:      parser.NewTableIdent("votes"),
									Qualifier: parser.NewTableIdent(""),
								},
							},
						},
						Val: parser.BoolVal(false),
					},
				},
				Else: parser.BoolVal(true),
			},
			As: parser.NewColIdent("i_like"),
		},
	},
	From: parser.TableExprs{
		&parser.JoinTableExpr{
			LeftExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("ideas"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			Join: "left join",
			RightExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("votes"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			On: &parser.AndExpr{
				Left: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator: "=",
						Left: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("ideas_id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("votes"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Right: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("ideas"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Escape: nil,
					},
				},
				Right: &parser.ParenExpr{
					Expr: &parser.ComparisonExpr{
						Operator: "=",
						Left: &parser.ColName{
							Metadata: nil,
							Name:     parser.NewColIdent("users_id"),
							Qualifier: parser.TableName{
								Name:      parser.NewTableIdent("votes"),
								Qualifier: parser.NewTableIdent(""),
							},
						},
						Right: &parser.SQLVal{
							Type: 0,
							Val:  []byte("538dc3820f942a4767ffcd75"),
						},
						Escape: nil,
					},
				},
			},
		},
	},
	Where:   nil,
	GroupBy: nil,
	Having:  nil,
	OrderBy: nil,
	Limit:   nil,
	Lock:    "",
}`)

	SelectQueries = append(SelectQueries, `&parser.Select{
	Cache:    "",
	Comments: nil,
	Distinct: "",
	Hints:    "",
	SelectExprs: parser.SelectExprs{
		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("email"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("fname"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("id"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("lname"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("content"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("posts"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("id"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("posts"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("title"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("posts"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("id"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("status"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},

		&parser.AliasedExpr{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("name"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("status"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			As: parser.NewColIdent(""),
		},
	},
	From: parser.TableExprs{
		&parser.JoinTableExpr{
			LeftExpr: &parser.JoinTableExpr{
				LeftExpr: &parser.AliasedTableExpr{
					Expr: parser.TableName{
						Name:      parser.NewTableIdent("posts"),
						Qualifier: parser.NewTableIdent(""),
					},
					As:    parser.NewTableIdent(""),
					Hints: nil,
				},
				Join: "join",
				RightExpr: &parser.AliasedTableExpr{
					Expr: parser.TableName{
						Name:      parser.NewTableIdent("users"),
						Qualifier: parser.NewTableIdent(""),
					},
					As:    parser.NewTableIdent("author"),
					Hints: nil,
				},
				On: &parser.ComparisonExpr{
					Operator: "=",
					Left: &parser.ColName{
						Metadata: nil,
						Name:     parser.NewColIdent("id"),
						Qualifier: parser.TableName{
							Name:      parser.NewTableIdent("author"),
							Qualifier: parser.NewTableIdent(""),
						},
					},
					Right: &parser.ColName{
						Metadata: nil,
						Name:     parser.NewColIdent("users_id"),
						Qualifier: parser.TableName{
							Name:      parser.NewTableIdent("posts"),
							Qualifier: parser.NewTableIdent(""),
						},
					},
					Escape: nil,
				},
			},
			Join: "left join",
			RightExpr: &parser.AliasedTableExpr{
				Expr: parser.TableName{
					Name:      parser.NewTableIdent("status"),
					Qualifier: parser.NewTableIdent(""),
				},
				As:    parser.NewTableIdent(""),
				Hints: nil,
			},
			On: &parser.ComparisonExpr{
				Operator: "=",
				Left: &parser.ColName{
					Metadata: nil,
					Name:     parser.NewColIdent("id"),
					Qualifier: parser.TableName{
						Name:      parser.NewTableIdent("status"),
						Qualifier: parser.NewTableIdent(""),
					},
				},
				Right: &parser.ColName{
					Metadata: nil,
					Name:     parser.NewColIdent("status_id"),
					Qualifier: parser.TableName{
						Name:      parser.NewTableIdent("posts"),
						Qualifier: parser.NewTableIdent(""),
					},
				},
				Escape: nil,
			},
		},
	},
	Where: &parser.Where{
		Type: "where",
		Expr: &parser.AndExpr{
			Left: &parser.ParenExpr{
				Expr: &parser.ComparisonExpr{
					Operator: "like",
					Left: &parser.ColName{
						Metadata: nil,
						Name:     parser.NewColIdent("fname"),
						Qualifier: parser.TableName{
							Name:      parser.NewTableIdent("author"),
							Qualifier: parser.NewTableIdent(""),
						},
					},
					Right: &parser.SQLVal{
						Type: 5,
						Val:  []byte(":v1"),
					},
					Escape: nil,
				},
			},
			Right: &parser.ParenExpr{
				Expr: &parser.ComparisonExpr{
					Operator: "in",
					Left: &parser.ColName{
						Metadata: nil,
						Name:     parser.NewColIdent("id"),
						Qualifier: parser.TableName{
							Name:      parser.NewTableIdent("author"),
							Qualifier: parser.NewTableIdent(""),
						},
					},
					Right: parser.ValTuple{
						&parser.SQLVal{
							Type: 5,
							Val:  []byte(":v2"),
						},

						&parser.SQLVal{
							Type: 5,
							Val:  []byte(":v3"),
						},
					},
					Escape: nil,
				},
			},
		},
	},
	GroupBy: nil,
	Having:  nil,
	OrderBy: parser.OrderBy{
		&parser.Order{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("fname"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			Direction: "desc",
		},

		&parser.Order{
			Expr: &parser.ColName{
				Metadata: nil,
				Name:     parser.NewColIdent("lname"),
				Qualifier: parser.TableName{
					Name:      parser.NewTableIdent("author"),
					Qualifier: parser.NewTableIdent(""),
				},
			},
			Direction: "desc",
		},
	},
	Limit: &parser.Limit{
		Offset: &parser.SQLVal{
			Type: 1,
			Val:  []byte("10"),
		},
		Rowcount: &parser.SQLVal{
			Type: 1,
			Val:  []byte("5"),
		},
	},
	Lock: "",
}`)

}
