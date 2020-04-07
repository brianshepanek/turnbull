package registry

import (
	"database/sql"
	mysql "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/comment/mysql"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type commentMysqlRepositoryRegistry struct {
	db    *sql.DB
	table string
}

func (r *registry) RegisterMysqlCommentRepository(db *sql.DB, table string) {
	r.commentMysqlRepositoryRegistry.db = db
	r.commentMysqlRepositoryRegistry.table = table
}
func (r *registry) newMysqlCommentRepository() repository.CommentRepository {
	return mysql.New(r.commentMysqlRepositoryRegistry.db, r.commentMysqlRepositoryRegistry.table)
}
