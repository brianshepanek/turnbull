package repository

import "database/sql"

type mysqlCommentRepository struct {
	mysqlCommentRepositoryStruct
}

func New(db *sql.DB, table string) *mysqlCommentRepository {
	return &mysqlCommentRepository{mysqlCommentRepositoryStruct{
		db:    db,
		table: table,
	}}
}
