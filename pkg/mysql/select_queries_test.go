package parser

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

}
