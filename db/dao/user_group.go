package dao

type UserGroupDAO struct {
	UserId  int64 `db:"user_id"`
	GroupId int64 `db:"group_id"`
}
