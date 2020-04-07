package registry

import (
	"database/sql"
	mysql "github.com/brianshepanek/turnbull/turnbull/output/interface/repository/post/mysql"
	repository "github.com/brianshepanek/turnbull/turnbull/output/usecase/repository"
)

type postMysqlRepositoryRegistry struct {
	db    *sql.DB
	table string
}

func (r *registry) RegisterMysqlPostRepository(db *sql.DB, table string) {
	r.postMysqlRepositoryRegistry.db = db
	r.postMysqlRepositoryRegistry.table = table
}
func (r *registry) newMysqlPostRepository() repository.PostRepository {
	return mysql.New(r.postMysqlRepositoryRegistry.db, r.postMysqlRepositoryRegistry.table)
}
