package dao

import "github.com/jmoiron/sqlx"

type UserPermissionDAO struct {
	User  int64 `db:"user"`
	Path  int64 `db:"path"`
	Read  bool  `db:"read"`
	Write bool  `db:"write"`
}

type UserPermissionRepository struct {
	tx *sqlx.Tx
}

func NewUserPermissionRepository(tx *sqlx.Tx) UserPermissionRepository {
	return UserPermissionRepository{tx: tx}
}

func (userPermRepo *UserPermissionRepository) CanUserAccessPath(userId int64, pathId int64) bool {
	var hasPerm = struct {
		HasAccess bool `db:"has_access"`
	}{}
	err := userPermRepo.tx.Get(&hasPerm,
		"SELECT COUNT(*) > 0 as has_access from user_permission up where up.user = ? and up.path = ? and (up.read = 1 or up.write = 1)",
		userId, pathId)
	if err != nil {
		panic(err)
	}
	return hasPerm.HasAccess
}
