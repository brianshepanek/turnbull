package registry

import http "github.com/brianshepanek/turnbull/turnbull/output/interface/controller/app/http"

type registry struct {
	postMongoRepositoryRegistry
	postMysqlRepositoryRegistry
	postRedisRepositoryRegistry
	postDefaultPresenterRegistry
	commentMongoRepositoryRegistry
	commentMysqlRepositoryRegistry
	commentRedisRepositoryRegistry
	commentDefaultPresenterRegistry
}

func New() *registry {
	return &registry{}
}
func (r *registry) NewHttpAppController() http.HttpAppController {
	r.NewHttpPostController()
	r.NewHttpCommentController()
	return r
}
