package repository

import "database/sql"

type mysqlPostRepository struct {
	mysqlPostRepositoryStruct
}

func New(db *sql.DB, table string) *mysqlPostRepository {
	return &mysqlPostRepository{mysqlPostRepositoryStruct{
		db:    db,
		table: table,
	}}
}
