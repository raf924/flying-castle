package dao

type UserDAO struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
