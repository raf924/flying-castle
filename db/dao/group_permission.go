package dao

import "github.com/jmoiron/sqlx"

type GroupPermissionDAO struct {
	Group int64 `db:"group"`
	Path  int64 `db:"path"`
	Read  bool  `db:"read"`
	Write bool  `db:"write"`
}

type GroupPermissionRepository struct {
	tx *sqlx.Tx
}

func NewGroupPermissionRepository(tx *sqlx.Tx) GroupPermissionRepository {
	return GroupPermissionRepository{tx: tx}
}

func (groupPermRepo *GroupPermissionRepository) CanGroupAccessPath(groupId int64, pathId int64) bool {
	var hasPerm = struct {
		HasAccess bool `db:"has_access"`
	}{}
	err := groupPermRepo.tx.Get(&hasPerm,
		"SELECT COUNT(*) > 0 as has_access from group_permission gp where gp.\"group\" = ? and gp.path = ? and (gp.read = 1 or gp.write = 1)",
		groupId, pathId)
	if err != nil {
		panic(err)
	}
	return hasPerm.HasAccess
}
