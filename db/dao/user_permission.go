package dao

type UserPermissionDAO struct {
	User  int64 `db:"user"`
	Path  int64 `db:"path"`
	Read  bool  `db:"read"`
	Write bool  `db:"write"`
}
