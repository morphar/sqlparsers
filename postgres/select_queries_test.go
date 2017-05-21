package parser

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

}
