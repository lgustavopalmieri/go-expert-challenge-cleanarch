package orderdb

import "database/sql"

type OrderRepositoyDb struct {
	Db *sql.DB
}

func NewOrderRepositoryDb(db *sql.DB) *OrderRepositoyDb {
	return &OrderRepositoyDb{Db: db}
}
