package sqlc

import "database/sql"

type Store interface {
	Querier
}

type SQLStore struct {
	Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	store := &SQLStore{db: db}
	return store
}
