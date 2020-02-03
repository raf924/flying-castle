package dao

type GroupPermissionDAO struct {
	Group int64 `db:"group"`
	Path  int64 `db:"path"`
	Read  bool  `db:"read"`
	Write bool  `db:"write"`
}
