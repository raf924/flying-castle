package dao

import "github.com/jmoiron/sqlx"

type GroupDAO struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}

type GroupRepository struct {
	tx *sqlx.Tx
}

func NewGroupRepository(tx *sqlx.Tx) GroupRepository {
	return GroupRepository{tx: tx}
}

func (groupRep *GroupRepository) GetById(id int64) GroupDAO {
	var groupDAO = GroupDAO{}
	var err = groupRep.tx.Get(&groupDAO, "SELECT * FROM \"group\" WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return groupDAO
}
